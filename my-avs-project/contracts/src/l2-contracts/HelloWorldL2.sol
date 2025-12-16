// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ITaskMailbox, ITaskMailboxTypes} from "@eigenlayer-contracts/src/contracts/interfaces/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

/// @title HelloWorldL2 - Sum Verification AVS usando TaskMailbox
/// @notice Questo contratto permette di richiedere la verifica di una somma tramite AVS
contract HelloWorldL2 {
    // ============ State Variables ============
    
    /// @notice Riferimento al contratto TaskMailbox di EigenLayer
    ITaskMailbox public immutable taskMailbox;
    
    /// @notice Indirizzo dell'AVS
    address public immutable avsAddress;
    
    /// @notice ID dell'operator set executor
    uint32 public immutable executorOperatorSetId;
    
    /// @notice Mapping da taskHash ai dati della richiesta
    mapping(bytes32 => SumVerificationRequest) public requests;
    
    /// @notice Contatore task per generare ID unici
    uint256 public taskCounter;
    
    // ============ Structs ============
    
    struct SumVerificationRequest {
        uint256 a;
        uint256 b;
        uint256 claimedResult;
        address requester;
        uint256 timestamp;
        bool verified;
        bool isCorrect;
    }
    
    // ============ Events ============
    
    event SumVerificationRequested(
        bytes32 indexed taskHash,
        uint256 indexed taskId,
        address indexed requester,
        uint256 a,
        uint256 b,
        uint256 claimedResult
    );
    
    event SumVerificationCompleted(
        bytes32 indexed taskHash,
        uint256 indexed taskId,
        bool isCorrect,
        uint256 actualResult
    );
    
    // ============ Constructor ============
    
    constructor(
        address _taskMailbox,
        address _avsAddress,
        uint32 _executorOperatorSetId
    ) {
        require(_taskMailbox != address(0), "Invalid TaskMailbox address");
        require(_avsAddress != address(0), "Invalid AVS address");
        
        taskMailbox = ITaskMailbox(_taskMailbox);
        avsAddress = _avsAddress;
        executorOperatorSetId = _executorOperatorSetId;
    }
    
    // ============ External Functions ============
    
    /// @notice Richiede la verifica di una somma tramite AVS
    /// @param a Primo numero
    /// @param b Secondo numero
    /// @param claimedResult Risultato dichiarato (a + b)
    /// @return taskHash Hash del task creato nel TaskMailbox
    function requestSumVerification(
        uint256 a,
        uint256 b,
        uint256 claimedResult
    ) external returns (bytes32 taskHash) {
        uint256 taskId = taskCounter++;
        
        // Codifica il payload: (taskId, a, b, claimedResult)
        bytes memory payload = abi.encode(taskId, a, b, claimedResult);
        
        // Prepara i parametri per il TaskMailbox
        OperatorSet memory executorOperatorSet = OperatorSet({
            avs: avsAddress,
            id: executorOperatorSetId
        });
        
        ITaskMailboxTypes.TaskParams memory taskParams = ITaskMailboxTypes.TaskParams({
            refundCollector: address(0), // Nessun refund collector per semplicità
            executorOperatorSet: executorOperatorSet,
            payload: payload
        });
        
        // Crea il task nel TaskMailbox
        taskHash = taskMailbox.createTask(taskParams);
        
        // Salva i dati della richiesta
        requests[taskHash] = SumVerificationRequest({
            a: a,
            b: b,
            claimedResult: claimedResult,
            requester: msg.sender,
            timestamp: block.timestamp,
            verified: false,
            isCorrect: false
        });
        
        emit SumVerificationRequested(
            taskHash,
            taskId,
            msg.sender,
            a,
            b,
            claimedResult
        );
        
        return taskHash;
    }
    
    /// @notice Completa la verifica di una somma (chiamato dall'AVS tramite TaskMailbox callback)
    /// @param taskHash Hash del task
    /// @param isCorrect Se la somma è corretta
    function completeSumVerification(
        bytes32 taskHash,
        bool isCorrect
    ) external {
        SumVerificationRequest storage request = requests[taskHash];
        require(request.timestamp != 0, "Request not found");
        require(!request.verified, "Already verified");
        
        // TODO: In produzione, aggiungere controllo che msg.sender sia autorizzato
        // (es. solo il TaskMailbox o l'AVS possono chiamare questa funzione)
        
        request.verified = true;
        request.isCorrect = isCorrect;
        
        uint256 actualResult = request.a + request.b;
        
        emit SumVerificationCompleted(
            taskHash,
            0, // taskId non disponibile qui
            isCorrect,
            actualResult
        );
    }
    
    // ============ View Functions ============
    
    /// @notice Verifica on-chain se una somma è corretta (funzione pure)
    /// @param a Primo numero
    /// @param b Secondo numero  
    /// @param result Risultato da verificare
    /// @return true se a + b == result
    function verifySumOnChain(
        uint256 a,
        uint256 b,
        uint256 result
    ) public pure returns (bool) {
        return (a + b) == result;
    }
    
    /// @notice Ottiene i dettagli di una richiesta
    /// @param taskHash Hash del task
    /// @return La richiesta di verifica
    function getRequest(bytes32 taskHash) external view returns (SumVerificationRequest memory) {
        return requests[taskHash];
    }
    
    /// @notice Verifica se un task è stato verificato
    /// @param taskHash Hash del task
    /// @return true se il task è stato verificato
    function isTaskVerified(bytes32 taskHash) external view returns (bool) {
        return requests[taskHash].verified;
    }
}
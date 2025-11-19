// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/console.sol";

/**
 * @title HelloWorldL1
 * @notice Contratto L1 per AVS di verifica somme
 * @dev Versione semplificata senza dipendenze EigenLayer complesse
 */
contract HelloWorldL1 {
    // Struttura per rappresentare un task di verifica somma
    struct SumVerificationTask {
        uint256 a;
        uint256 b;
        uint256 claimedResult;
        address requester;
        bool verified;
        bool isCorrect;
        uint256 timestamp;
    }
    
    // Storage
    mapping(uint256 => SumVerificationTask) public tasks;
    uint256 public taskCounter;
    
    // Eventi
    event SumVerificationRequested(
        uint256 indexed taskId,
        address indexed requester,
        uint256 a,
        uint256 b,
        uint256 claimedResult
    );
    
    event SumVerificationCompleted(
        uint256 indexed taskId,
        bool isCorrect,
        uint256 actualResult
    );
    
    constructor() {
        console.log("HelloWorldL1 deployed - Sum Verification AVS");
    }
    
    /**
     * @notice Richiedi verifica di una somma
     * @param a Primo addendo
     * @param b Secondo addendo
     * @param claimedResult Risultato dichiarato
     * @return taskId ID del task creato
     */
    function requestSumVerification(
        uint256 a,
        uint256 b,
        uint256 claimedResult
    ) external returns (uint256 taskId) {
        taskId = taskCounter++;
        
        tasks[taskId] = SumVerificationTask({
            a: a,
            b: b,
            claimedResult: claimedResult,
            requester: msg.sender,
            verified: false,
            isCorrect: false,
            timestamp: block.timestamp
        });
        
        emit SumVerificationRequested(taskId, msg.sender, a, b, claimedResult);
        
        console.log("Sum verification requested:");
        console.log("  Task ID:", taskId);
        console.log("  A:", a);
        console.log("  B:", b);
        console.log("  Claimed Result:", claimedResult);
        
        return taskId;
    }
    
    /**
     * @notice Completa la verifica (chiamato dagli operatori AVS)
     * @param taskId ID del task da verificare
     * @param isCorrect Se la somma è corretta
     */
    function completeSumVerification(
        uint256 taskId,
        bool isCorrect
    ) external {
        require(!tasks[taskId].verified, "Task already verified");
        
        tasks[taskId].verified = true;
        tasks[taskId].isCorrect = isCorrect;
        
        uint256 actualResult = tasks[taskId].a + tasks[taskId].b;
        
        emit SumVerificationCompleted(taskId, isCorrect, actualResult);
        
        console.log("Sum verification completed:");
        console.log("  Task ID:", taskId);
        console.log("  Is Correct:", isCorrect);
        console.log("  Actual Result:", actualResult);
    }
    
    /**
     * @notice Verifica on-chain di una somma (per testing)
     * @param a Primo addendo
     * @param b Secondo addendo
     * @param result Risultato da verificare
     * @return true se la somma è corretta
     */
    function verifySumOnChain(
        uint256 a,
        uint256 b,
        uint256 result
    ) external pure returns (bool) {
        return (a + b == result);
    }
    
    /**
     * @notice Ottieni i dettagli di un task
     * @param taskId ID del task
     * @return task Dettagli del task
     */
    function getTask(uint256 taskId) external view returns (SumVerificationTask memory) {
        return tasks[taskId];
    }
    
    /**
     * @notice Verifica se un task è stato completato
     * @param taskId ID del task
     * @return true se verificato
     */
    function isTaskVerified(uint256 taskId) external view returns (bool) {
        return tasks[taskId].verified;
    }
}
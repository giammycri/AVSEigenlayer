// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IAVSRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSRegistrar.sol";

/**
 * @title TaskAVSRegistrar
 * @notice Manages operator registration and allowlist for task-based AVS
 * @dev Implements the ITaskAVSRegistrarBase interface expected by Hourglass
 */
contract TaskAVSRegistrar is IAVSRegistrar, Initializable, OwnableUpgradeable {
    
    // ============ Structs ============
    
    /**
     * @notice Configuration for the Task-based AVS
     * @param aggregatorOperatorSetId The operator set ID responsible for aggregating results
     * @param executorOperatorSetIds Array of operator set IDs responsible for executing tasks
     */
    struct AvsConfig {
        uint32 aggregatorOperatorSetId;
        uint32[] executorOperatorSetIds;
    }
    
    // ============ State Variables ============
    
    // EigenLayer contract references
    IAllocationManager public immutable allocationManager;
    IKeyRegistrar public immutable keyRegistrar;
    IPermissionController public immutable permissionController;
    
    // AVS configuration
    address public avsAddress;
    
    // Operator Set configuration
    uint32 public aggregatorOperatorSetId;
    uint32[] public executorOperatorSetIds;
    
    // Operator tracking
    mapping(address => bool) private _registeredOperators;
    mapping(address => bool) private _allowlistedOperators;
    address[] public registeredOperatorsList;
    
    // ⭐ AGGIUNTO: Socket management
    mapping(address => string) private _operatorSockets;
    
    // ============ Events ============
    
    event OperatorRegistered(address indexed operator);
    event OperatorDeregistered(address indexed operator);
    event OperatorAllowlisted(address indexed operator, bool allowed);
    event AVSInitialized(address indexed avs, address indexed owner);
    event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds);
    event OperatorSocketSet(address indexed operator, string socket);
    
    // ============ Constructor ============
    
    /**
     * @dev Constructor - sets immutable EigenLayer contract references
     */
    constructor(
        IAllocationManager _allocationManager,
        IKeyRegistrar _keyRegistrar,
        IPermissionController _permissionController
    ) {
        require(address(_allocationManager) != address(0), "Invalid AllocationManager");
        require(address(_keyRegistrar) != address(0), "Invalid KeyRegistrar");
        require(address(_permissionController) != address(0), "Invalid PermissionController");
        
        allocationManager = _allocationManager;
        keyRegistrar = _keyRegistrar;
        permissionController = _permissionController;
        
        _disableInitializers();
    }

    // ============ Initialization ============

    /**
     * @notice Initialize the AVS
     */
    function initialize(
        address _avs,
        address _owner
    ) 
        external 
        initializer 
    {
        require(_avs != address(0), "Invalid AVS address");
        require(_owner != address(0), "Invalid owner address");
        
        __Ownable_init();
        _transferOwnership(_owner);
        avsAddress = _avs;
        
        // Set default operator set IDs based on Hourglass standard configuration
        // Aggregator uses operator set 0, Executor uses operator set 1
        aggregatorOperatorSetId = 0;
        executorOperatorSetIds.push(1);
        
        emit AVSInitialized(_avs, _owner);
        emit AvsConfigSet(aggregatorOperatorSetId, executorOperatorSetIds);
    }

    // ============ AVS Configuration Functions ============

    /**
     * @notice Set AVS operator set configuration
     * @param config Configuration struct containing aggregator and executor operator set IDs
     */
    function setAvsConfig(AvsConfig memory config) 
        external 
        onlyOwner 
    {
        require(config.executorOperatorSetIds.length > 0, "Executor operator set IDs cannot be empty");
        
        aggregatorOperatorSetId = config.aggregatorOperatorSetId;
        
        // Clear existing array and copy new values
        delete executorOperatorSetIds;
        for (uint i = 0; i < config.executorOperatorSetIds.length; i++) {
            executorOperatorSetIds.push(config.executorOperatorSetIds[i]);
        }
        
        emit AvsConfigSet(config.aggregatorOperatorSetId, config.executorOperatorSetIds);
    }

    /**
     * @notice Get AVS configuration
     * @return config Configuration struct containing aggregator and executor operator set IDs
     */
    function getAvsConfig() 
        external 
        view 
        returns (AvsConfig memory config) 
    {
        config.aggregatorOperatorSetId = aggregatorOperatorSetId;
        config.executorOperatorSetIds = executorOperatorSetIds;
        return config;
    }

    // ============ AVS Support ============

    /**
     * @notice Check if this registrar supports a specific AVS
     * @param avs The AVS address to check
     * @return bool True if this registrar supports the AVS
     */
    function supportsAVS(address avs) external view returns (bool) {
        return avs == avsAddress;
    }

    // ============ Allowlist Management ============

    /**
     * @notice Add operator to allowlist
     * @param operator The operator address to allowlist
     */
    function addOperatorToAllowlist(address operator) external onlyOwner {
        _setOperatorAllowlist(operator, true);
    }

    /**
     * @notice Remove operator from allowlist
     * @param operator The operator address to remove
     */
    function removeOperatorFromAllowlist(address operator) external onlyOwner {
        _setOperatorAllowlist(operator, false);
    }

    /**
     * @notice Internal function to set operator allowlist status
     */
    function _setOperatorAllowlist(address operator, bool allowed) internal {
        require(operator != address(0), "Invalid operator");
        _allowlistedOperators[operator] = allowed;
        emit OperatorAllowlisted(operator, allowed);
    }

    // ============ Socket Management ============

    /**
     * ⭐ AGGIUNTO: Get operator socket (RICHIESTO DA HOURGLASS)
     * @param operator The operator address
     * @return socket The operator's network socket
     */
    function getOperatorSocket(address operator) 
        external 
        view 
        returns (string memory) 
    {
        return _operatorSockets[operator];
    }

    /**
     * ⭐ AGGIUNTO: Set operator socket
     * @param operator The operator address
     * @param socket The operator's network socket
     */
    function setOperatorSocket(address operator, string memory socket) 
        external 
        onlyOwner 
    {
        require(operator != address(0), "Invalid operator");
        require(bytes(socket).length > 0, "Invalid socket");
        _operatorSockets[operator] = socket;
        emit OperatorSocketSet(operator, socket);
    }

    /**
     * ⭐ AGGIUNTO: Internal function to set operator socket
     * @param operator The operator address
     * @param socket The operator's network socket
     */
    function _setOperatorSocket(address operator, string memory socket) internal {
        _operatorSockets[operator] = socket;
        emit OperatorSocketSet(operator, socket);
    }

    // ============ Operator Registration ============

    /**
     * @notice Register an operator - implements IAVSRegistrar interface
     * ⭐ MODIFICATO: Ora decodifica il socket dal parametro operatorSignature
     */
    function registerOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        bytes memory operatorSignature
    ) 
        external 
    {
        require(avs == avsAddress, "Invalid AVS");
        require(operator != address(0), "Invalid operator");
        require(!_registeredOperators[operator], "Already registered");
        require(_allowlistedOperators[operator], "Operator not allowlisted");
        
        // ⭐ AGGIUNTO: Decodifica il socket dal payload
        // Il parametro operatorSignature contiene il socket encodato come string
        if (operatorSignature.length > 0) {
            string memory socket = abi.decode(operatorSignature, (string));
            _setOperatorSocket(operator, socket);
        }
        
        _registeredOperators[operator] = true;
        registeredOperatorsList.push(operator);
        
        emit OperatorRegistered(operator);
    }

    /**
     * @notice Deregister an operator - implements IAVSRegistrar interface
     */
    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) 
        external 
    {
        require(avs == avsAddress, "Invalid AVS");
        require(_registeredOperators[operator], "Not registered");
        
        _registeredOperators[operator] = false;
        
        emit OperatorDeregistered(operator);
    }

    // ============ View Functions ============

    /**
     * @notice Check if operator is registered
     */
    function isOperatorRegistered(address operator) 
        external 
        view 
        returns (bool) 
    {
        return _registeredOperators[operator];
    }
    
    /**
     * @notice Check if operator is allowlisted
     */
    function isOperatorAllowlisted(address operator) 
        external 
        view 
        returns (bool) 
    {
        return _allowlistedOperators[operator];
    }
    
    /**
     * @notice Get all registered operators
     */
    function getRegisteredOperators() 
        external 
        view 
        returns (address[] memory) 
    {
        uint256 count = 0;
        for (uint256 i = 0; i < registeredOperatorsList.length; i++) {
            if (_registeredOperators[registeredOperatorsList[i]]) {
                count++;
            }
        }
        
        address[] memory activeOperators = new address[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < registeredOperatorsList.length; i++) {
            if (_registeredOperators[registeredOperatorsList[i]]) {
                activeOperators[index] = registeredOperatorsList[i];
                index++;
            }
        }
        
        return activeOperators;
    }
}
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
 */
contract TaskAVSRegistrar is IAVSRegistrar, Initializable, OwnableUpgradeable {
    
    // EigenLayer contract references
    IAllocationManager public immutable allocationManager;
    IKeyRegistrar public immutable keyRegistrar;
    IPermissionController public immutable permissionController;
    
    // AVS configuration
    address public avsAddress;
    
    // Operator tracking
    mapping(address => bool) private _registeredOperators;
    mapping(address => bool) private _allowlistedOperators;
    address[] public registeredOperatorsList;
    
    // Events
    event OperatorRegistered(address indexed operator);
    event OperatorDeregistered(address indexed operator);
    event OperatorAllowlisted(address indexed operator, bool allowed);
    event AVSInitialized(address indexed avs, address indexed owner);
    
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
        
        emit AVSInitialized(_avs, _owner);
    }

    /**
     * @notice Check if this registrar supports a specific AVS
     * @param avs The AVS address to check
     * @return bool True if this registrar supports the AVS
     */
    function supportsAVS(address avs) external view returns (bool) {
        return avs == avsAddress;
    }

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

    /**
     * @notice Register an operator - implements IAVSRegistrar interface
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
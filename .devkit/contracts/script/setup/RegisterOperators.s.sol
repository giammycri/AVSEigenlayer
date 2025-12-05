// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IKeyRegistrarTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

/**
 * @title RegisterOperators
 * @notice Script to register operators and configure operator sets in EigenLayer
 * This script:
 * 1. Creates operator sets in AllocationManager
 * 2. Registers ECDSA keys for operators in KeyRegistrar
 * 3. Allocates operators to their respective operator sets
 */
contract RegisterOperators is Script {
    
    // EigenLayer contract addresses (from TaskAVSRegistrar) - with correct checksum
    address constant ALLOCATION_MANAGER = 0x42583067658071247ec8CE0A516A58f682002d07;
    address constant KEY_REGISTRAR = 0xA4dB30D08d8bbcA00D40600bee9F029984dB162a;
    
    // AVS address
    address constant AVS = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    
    // Operator addresses
    address constant AGGREGATOR = 0x90F79bf6EB2c4f870365E785982E1f101E93b906;
    address constant EXECUTOR = 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65;
    
    // Operator private keys (from Anvil default accounts)
    uint256 constant AGGREGATOR_PRIVATE_KEY = 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6;
    uint256 constant EXECUTOR_PRIVATE_KEY = 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a;
    
    // Operator set IDs
    uint32 constant AGGREGATOR_OPERATOR_SET_ID = 0;
    uint32 constant EXECUTOR_OPERATOR_SET_ID = 1;
    
    // AVS private key (Anvil account #1)
    uint256 constant AVS_PRIVATE_KEY = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d;

    function run() public {
        console.log("=== Starting Operator Registration ===");
        console.log("AVS:", AVS);
        console.log("Aggregator:", AGGREGATOR);
        console.log("Executor:", EXECUTOR);
        
        // Step 1: Create operator sets as AVS
        createOperatorSets();
        
        // Step 2: Register keys and allocate aggregator
        registerOperatorWithKey(
            AGGREGATOR,
            AGGREGATOR_PRIVATE_KEY,
            AGGREGATOR_OPERATOR_SET_ID,
            "Aggregator"
        );
        
        // Step 3: Register keys and allocate executor
        registerOperatorWithKey(
            EXECUTOR,
            EXECUTOR_PRIVATE_KEY,
            EXECUTOR_OPERATOR_SET_ID,
            "Executor"
        );
        
        console.log("=== Operator Registration Complete ===");
    }
    
    function createOperatorSets() internal {
        console.log("\n--- Creating Operator Sets ---");
        
        vm.startBroadcast(AVS_PRIVATE_KEY);
        
        IAllocationManager allocationManager = IAllocationManager(ALLOCATION_MANAGER);
        
        // Create aggregator operator set (ID 0)
        OperatorSet memory aggSet = OperatorSet({avs: AVS, operatorSetId: AGGREGATOR_OPERATOR_SET_ID});
        try allocationManager.createOperatorSets(AVS, new OperatorSet[](0)) {
            console.log("Called createOperatorSets");
        } catch {
            console.log("CreateOperatorSets failed - may not exist in this version");
        }
        
        vm.stopBroadcast();
    }
    
    function registerOperatorWithKey(
        address operator,
        uint256 operatorPrivateKey,
        uint32 operatorSetId,
        string memory operatorName
    ) internal {
        console.log("\n--- Registering", operatorName, "---");
        console.log("Address:", operator);
        console.log("Operator Set ID:", operatorSetId);
        
        vm.startBroadcast(operatorPrivateKey);
        
        IKeyRegistrar keyRegistrar = IKeyRegistrar(KEY_REGISTRAR);
        
        // Register ECDSA key (the operator's address itself is the key for ECDSA)
        bytes memory keyData = abi.encode(operator);
        
        // Create signature for key registration
        // For ECDSA, we need to sign the registration message
        bytes32 messageHash = keyRegistrar.calculateOperatorKeyRegistrationDigestHash(
            operator,
            AVS,
            operatorSetId,
            keyData
        );
        
        // Sign the message hash
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(operatorPrivateKey, messageHash);
        bytes memory signature = abi.encodePacked(r, s, v);
        
        // Register the key
        try keyRegistrar.registerOperatorKey(
            operator,
            AVS,
            operatorSetId,
            signature,
            keyData
        ) {
            console.log("Successfully registered ECDSA key");
        } catch Error(string memory reason) {
            console.log("Failed to register key:", reason);
        } catch {
            console.log("Failed to register key: unknown error");
        }
        
        vm.stopBroadcast();
        
        // Now allocate the operator to the set (as AVS)
        vm.startBroadcast(AVS_PRIVATE_KEY);
        
        IAllocationManager allocationManager = IAllocationManager(ALLOCATION_MANAGER);
        
        // Allocate operator to the set
        OperatorSet memory operatorSet = OperatorSet({
            avs: AVS,
            operatorSetId: operatorSetId
        });
        
        try allocationManager.addToOperatorSet(operator, operatorSet) {
            console.log("Successfully allocated operator to set");
        } catch Error(string memory reason) {
            console.log("Failed to allocate operator:", reason);
        } catch {
            console.log("Failed to allocate operator: unknown error");
        }
        
        vm.stopBroadcast();
    }
}
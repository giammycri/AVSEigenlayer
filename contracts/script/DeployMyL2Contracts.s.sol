// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {HelloWorldL2} from "../src/l2-contracts/HelloWorldL2.sol";
import {AVSTaskHook} from "../src/l2-contracts/AVSTaskHook.sol";

contract DeployMyL2Contracts is Script {
    function run(string memory, string memory) external {
        // Hardcoded values per devnet - leggono dalle variabili d'ambiente o context
        // TaskMailbox L2 address (sempre lo stesso per devnet)
        address taskMailboxAddress = 0xB99CC53e8db7018f557606C2a5B066527bF96b26;
        
        // AVS address (sempre lo stesso per devnet)
        address avsAddress = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        
        // Executor operator set ID (sempre 1 per devnet)
        uint32 executorOperatorSetId = 1;
        
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        
        console.log("=== Deploying L2 contracts ===");
        console.log("TaskMailbox:", taskMailboxAddress);
        console.log("AVS:", avsAddress);
        console.log("Executor OperatorSet ID:", executorOperatorSetId);
        
        vm.startBroadcast(deployerPrivateKey);

        HelloWorldL2 helloWorldL2 = new HelloWorldL2(
            taskMailboxAddress,
            avsAddress,
            executorOperatorSetId
        );
        console.log("HelloWorldL2 deployed at:", address(helloWorldL2));

        AVSTaskHook avsTaskHook = new AVSTaskHook();
        console.log("AVSTaskHook deployed at:", address(avsTaskHook));

        vm.stopBroadcast();
        
        console.log("=== L2 Deployment Complete ===");
    }
}
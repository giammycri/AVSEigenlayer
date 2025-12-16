// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {HelloWorldL2} from "../src/l2-contracts/HelloWorldL2.sol";
import {AVSTaskHook} from "../src/l2-contracts/AVSTaskHook.sol";

contract DeployMyL2Contracts is Script {
    function run(string memory, string memory) external {
        // Use hardcoded "devnet" as environment name
        string memory environment = "devnet";
        
        // Hardcoded values per devnet - leggono dalle variabili d'ambiente o context
        // TaskMailbox L2 address (sempre lo stesso per devnet)
        address taskMailboxAddress = 0xB99CC53e8db7018f557606C2a5B066527bF96b26;
        
        // AVS address (sempre lo stesso per devnet)
        address avsAddress = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        
        // Executor operator set ID (sempre 1 per devnet)
        uint32 executorOperatorSetId = 1;
        
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("=== Deploying L2 contracts ===");
        console.log("Deployer address:", deployer);
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
        
        // Write deployment info to output file
        _writeOutputToJson(environment, address(helloWorldL2), address(avsTaskHook));
    }
    
    function _writeOutputToJson(
        string memory environment,
        address helloWorldL2,
        address avsTaskHook
    ) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        addresses = vm.serializeAddress(addresses, "HelloWorldL2", helloWorldL2);
        addresses = vm.serializeAddress(addresses, "avsTaskHook", avsTaskHook);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat(
            "script/",
            environment,
            "/output/deploy_custom_contracts_l2_output.json"
        );
        vm.writeJson(finalJson, outputFile);
        
        console.log("Output written to:", outputFile);
    }
}

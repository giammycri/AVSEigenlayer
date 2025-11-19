// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "@project/l1-contracts/HelloWorldL1.sol";

contract DeployMyL1Contracts is Script {
    function run(string memory, string memory) public {
        console.log("=== Deploying Sum Verification L1 Contracts ===");
        
        // Use Anvil's default deployer private key for devnet (from devnet.yaml)
        // This is safe for local testing only
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Deploying with address:", deployer);
        
        vm.startBroadcast(deployerPrivateKey);
        
        HelloWorldL1 sumVerifier = new HelloWorldL1();
        
        console.log("HelloWorldL1 deployed at:", address(sumVerifier));
        
        vm.stopBroadcast();
        
        console.log("=== Sum Verification L1 Deployment Complete ===");
    }
}
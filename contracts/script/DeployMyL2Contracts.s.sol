// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import {HelloWorldL2} from "../src/l2-contracts/HelloWorldL2.sol";
import {AVSTaskHook} from "../src/l2-contracts/AVSTaskHook.sol";

contract DeployMyL2Contracts is Script {
    function run(string memory, string memory) external {
        // Use Anvil's default deployer private key for devnet (from devnet.yaml)
        // This is safe for local testing only
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Deploying L2 contracts with deployer:", deployer);
        
        vm.startBroadcast(deployerPrivateKey);

        HelloWorldL2 helloWorldL2 = new HelloWorldL2();
        console.log("HelloWorldL2 deployed at:", address(helloWorldL2));

        AVSTaskHook avsTaskHook = new AVSTaskHook();
        console.log("AVSTaskHook deployed at:", address(avsTaskHook));

        vm.stopBroadcast();
    }
}
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";

interface ISumVerification {
    function requestSumVerification(uint256 a, uint256 b, uint256 c) external returns (uint256, bytes32);
}

contract CreateSumTask is Script {
    function run(uint256 a, uint256 b, uint256 claimedResult) external {
        address sumVerification = 0x093Dc47bBd22C27D4e6996bdD34cb7eE7FfbA657;
        
        console.log("Creating sum verification task...");
        console.log("A:", a);
        console.log("B:", b);
        console.log("Claimed Result:", claimedResult);
        console.log("Contract:", sumVerification);
        
        vm.startBroadcast();
        
        ISumVerification sum = ISumVerification(sumVerification);
        
        try sum.requestSumVerification(a, b, claimedResult) returns (uint256 taskId, bytes32 taskHash) {
            console.log("Task created successfully!");
            console.log("Task ID:", taskId);
            console.log("Task Hash:");
            console.logBytes32(taskHash);
        } catch Error(string memory reason) {
            console.log("Error:", reason);
            revert(reason);
        } catch (bytes memory lowLevelData) {
            console.log("Low-level error:");
            console.logBytes(lowLevelData);
            revert("Task creation failed");
        }
        
        vm.stopBroadcast();
    }
}
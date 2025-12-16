// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "@project/l1-contracts/HelloWorldL1.sol";  // ← CAMBIATO da ../src/ a @project/

contract HelloWorldL1Test is Test {
    HelloWorldL1 public helloWorld;
    
    address public user1 = address(0x1);
    address public operator = address(0x2);
    
    function setUp() public {
        helloWorld = new HelloWorldL1();
    }
    
    function testDeployment() public view {
        // Verifica che il contratto sia stato deployato
        assertTrue(address(helloWorld) != address(0));
    }
    
    function testRequestSumVerification() public {
        uint256 a = 5;
        uint256 b = 10;
        uint256 claimedResult = 15;
        
        vm.prank(user1);
        uint256 taskId = helloWorld.requestSumVerification(a, b, claimedResult);
        
        // Verifica che il task sia stato creato
        assertEq(taskId, 0);
        
        // Verifica i dati del task
        HelloWorldL1.SumVerificationTask memory task = helloWorld.getTask(taskId);
        assertEq(task.a, a);
        assertEq(task.b, b);
        assertEq(task.claimedResult, claimedResult);
        assertEq(task.requester, user1);
        assertFalse(task.verified);
    }
    
    function testCompleteSumVerification() public {
        // Crea un task
        vm.prank(user1);
        uint256 taskId = helloWorld.requestSumVerification(5, 10, 15);
        
        // Completa la verifica
        vm.prank(operator);
        helloWorld.completeSumVerification(taskId, true);
        
        // Verifica che il task sia stato completato
        HelloWorldL1.SumVerificationTask memory task = helloWorld.getTask(taskId);
        assertTrue(task.verified);
        assertTrue(task.isCorrect);
    }
    
    function testVerifySumOnChain() public view {
        assertTrue(helloWorld.verifySumOnChain(5, 10, 15));
        assertFalse(helloWorld.verifySumOnChain(5, 10, 14));
    }
    
    function testCannotVerifyTwice() public {
        vm.prank(user1);
        uint256 taskId = helloWorld.requestSumVerification(5, 10, 15);
        
        vm.prank(operator);
        helloWorld.completeSumVerification(taskId, true);
        
        // Prova a verificare di nuovo
        vm.prank(operator);
        vm.expectRevert("Task already verified");
        helloWorld.completeSumVerification(taskId, true);
    }
    
    function testMultipleTasks() public {
        vm.startPrank(user1);
        uint256 task1 = helloWorld.requestSumVerification(1, 2, 3);
        uint256 task2 = helloWorld.requestSumVerification(10, 20, 30);
        uint256 task3 = helloWorld.requestSumVerification(100, 200, 300);
        vm.stopPrank();
        
        assertEq(task1, 0);
        assertEq(task2, 1);
        assertEq(task3, 2);
        
        assertEq(helloWorld.taskCounter(), 3);
    }
}
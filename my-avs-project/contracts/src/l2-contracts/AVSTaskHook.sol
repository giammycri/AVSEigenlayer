// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

struct OperatorSet {
    address avs;
    uint32 id;
}

struct TaskParams {
    address refundCollector;
    OperatorSet executorOperatorSet;
    bytes payload;
}

/// @title AVSTaskHook
/// @notice Hook contract for AVS tasks on L2
contract AVSTaskHook {
    mapping(bytes32 => bytes) public taskResults;
    mapping(bytes32 => bool) public taskCompleted;
    
    event TaskCallbackReceived(bytes32 indexed taskHash, bytes result);

    function validatePreTaskCreation(
        address creator,
        TaskParams calldata taskParams
    ) external pure {
        // Accept all tasks
    }

    function calculateTaskFee(
        TaskParams calldata taskParams
    ) external pure returns (uint96 fee) {
        return 0;
    }

    function handlePostTaskCreation(
        bytes32 taskHash
    ) external {
        // Empty
    }

    function validatePreTaskResultSubmission(
        address sender,
        bytes32 taskHash,
        bytes calldata executorCert,
        bytes calldata result
    ) external pure {
        // Accept all
    }

    function validatePostTaskExecution(
        bytes32 taskHash,
        bytes calldata result
    ) external pure {
        // Accept all
    }

    /// @notice QUESTO È IL METODO CHE VIENE CHIAMATO DAL TASKMAILBOX
    function handlePostTaskResultSubmission(
        address sender,
        bytes32 taskHash
    ) external {
        // Accetta tutto - il contratto può essere vuoto
    }

    function handleCallback(
        bytes32 taskHash,
        bytes calldata result
    ) external {
        taskResults[taskHash] = result;
        taskCompleted[taskHash] = true;
        emit TaskCallbackReceived(taskHash, result);
    }
    
    function getTaskResult(bytes32 taskHash) external view returns (bytes memory) {
        require(taskCompleted[taskHash], "Task not completed");
        return taskResults[taskHash];
    }
    
    function isTaskResultCorrect(bytes32 taskHash) external view returns (bool) {
        require(taskCompleted[taskHash], "Task not completed");
        bytes memory result = taskResults[taskHash];
        if (result.length >= 32) {
            return uint8(result[31]) == 1;
        }
        return false;
    }
}
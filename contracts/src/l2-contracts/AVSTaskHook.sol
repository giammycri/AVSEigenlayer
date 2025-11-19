// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/// @title AVSTaskHook
/// @notice Hook contract for AVS tasks on L2. Extend with your logic as needed.
contract AVSTaskHook {
    // Event emitted when a task is hooked
    event TaskHooked(address indexed sender, bytes data);

    /// @notice Example hook function
    /// @param data Arbitrary data for the task
    function hookTask(bytes calldata data) external {
        emit TaskHooked(msg.sender, data);
        // Qui puoi aggiungere la logica di gestione del task
    }
}
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test} from "forge-std/Test.sol";
import {console2} from "forge-std/Test.sol";
import {TaskAVSRegistrar} from "@project/l1-contracts/TaskAVSRegistrar.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";

contract TaskAVSRegistrarTest is Test {
    TaskAVSRegistrar public registrar;
    
    address public mockAllocationManager = address(0x100);
    address public mockKeyRegistrar = address(0x200);
    address public mockPermissionController = address(0x300);
    
    address public avs = address(0x400);
    address public owner = address(0x500);
    address public operator = address(0x600);

    function setUp() public {
        // Deploy TaskAVSRegistrar
        registrar = new TaskAVSRegistrar(
            IAllocationManager(mockAllocationManager),
            IKeyRegistrar(mockKeyRegistrar),
            IPermissionController(mockPermissionController)
        );
    }

    function testConstructor() public view {
        // Test that contract was deployed
        assertTrue(address(registrar) != address(0), "Registrar should be deployed");
    }

    function testCanCheckOperatorRegistration() public view {
        // Test that we can call isOperatorRegistered
        bool isRegistered = registrar.isOperatorRegistered(operator);
        assertFalse(isRegistered, "Operator should not be registered initially");
    }
}
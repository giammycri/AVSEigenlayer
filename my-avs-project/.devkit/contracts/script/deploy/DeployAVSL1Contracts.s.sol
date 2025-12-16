// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IAVSRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSRegistrar.sol";

import {TaskAVSRegistrar} from "@project/l1-contracts/TaskAVSRegistrar.sol";

contract DeployAVSL1Contracts is Script {
    function run(
        string memory environment,
        address avs,
        address allocationManager,
        address keyRegistrar,
        address permissionController,
        uint32 aggregatorOperatorSetId,
        uint32 executorOperatorSetId,
        address[] memory aggregatorWhitelistedOperators,
        address[] memory executorWhitelistedOperators
    ) public {
        // Load the private key from the environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");

        // Deploy the TaskAVSRegistrar middleware contract with proxy pattern
        vm.startBroadcast(deployerPrivateKey);
        console.log("Deployer address:", vm.addr(deployerPrivateKey));

        // Deploy ProxyAdmin
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        console.log("ProxyAdmin deployed to:", address(proxyAdmin));

        // Deploy implementation
        TaskAVSRegistrar taskAVSRegistrarImpl = new TaskAVSRegistrar(
            IAllocationManager(allocationManager),
            IKeyRegistrar(keyRegistrar),
            IPermissionController(permissionController)
        );
        console.log("TaskAVSRegistrar implementation deployed to:", address(taskAVSRegistrarImpl));

        // Deploy proxy with initialization
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(taskAVSRegistrarImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(
                TaskAVSRegistrar.initialize.selector, 
                avs, 
                vm.addr(deployerPrivateKey)
            )
        );
        console.log("TaskAVSRegistrar proxy deployed to:", address(proxy));

        // Add aggregator operators to allowlist
        console.log("Adding aggregator operators to allowlist...");
        for (uint256 i = 0; i < aggregatorWhitelistedOperators.length; i++) {
            TaskAVSRegistrar(address(proxy)).addOperatorToAllowlist(
                aggregatorWhitelistedOperators[i]
            );
            console.log("Added aggregator operator to allowlist:", aggregatorWhitelistedOperators[i]);
        }

        // Add executor operators to allowlist
        console.log("Adding executor operators to allowlist...");
        for (uint256 i = 0; i < executorWhitelistedOperators.length; i++) {
            TaskAVSRegistrar(address(proxy)).addOperatorToAllowlist(
                executorWhitelistedOperators[i]
            );
            console.log("Added executor operator to allowlist:", executorWhitelistedOperators[i]);
        }

        console.log("WARNING: TaskAVSRegistrar ownership NOT transferred - still owned by deployer");
        console.log("Transfer ownership manually after AVS registration if needed");

        // Transfer ProxyAdmin ownership to avs (or a multisig in production)
        proxyAdmin.transferOwnership(avs);
        console.log("Transferred ProxyAdmin ownership to:", avs);

        vm.stopBroadcast();

        // Register AVS with EigenLayer AllocationManager
        // This must be called FROM the AVS address
        console.log("Registering AVS with EigenLayer...");
        vm.startBroadcast(vm.envUint("AVS_PRIVATE_KEY")); // Use AVS private key
        
        IAllocationManager(allocationManager).setAVSRegistrar(avs, IAVSRegistrar(address(proxy)));
        console.log("Successfully registered AVS with EigenLayer");
        
        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, address(proxy), address(taskAVSRegistrarImpl), address(proxyAdmin));
    }

    function _writeOutputToJson(
        string memory environment,
        address taskAVSRegistrarProxy,
        address taskAVSRegistrarImpl,
        address proxyAdmin
    ) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        vm.serializeAddress(addresses, "taskAVSRegistrar", taskAVSRegistrarProxy);
        vm.serializeAddress(addresses, "taskAVSRegistrarImpl", taskAVSRegistrarImpl);
        addresses = vm.serializeAddress(addresses, "l1ProxyAdmin", proxyAdmin);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/", environment, "/output/deploy_avs_l1_output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
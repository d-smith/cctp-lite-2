// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/Token.sol";

contract DeployScript is Script {
    uint256 private deployerPrivateKey;
    address private ownerAddress;

    function setUp() public {
        deployerPrivateKey = vm.envUint("DEPLOYER_KEY");
        ownerAddress = vm.envAddress("OWNER_ADDRESS");
    }

    function run() public {
        vm.startBroadcast(deployerPrivateKey);

        FiddyCent fiddyCent = new FiddyCent(ownerAddress);

        vm.stopBroadcast();
    }
}
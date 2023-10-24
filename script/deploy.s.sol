// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/Token.sol";
import "../src/Transporter.sol";

contract DeployScript is Script {
    uint256 private deployerPrivateKey;
    address private ownerAddress;

    uint32 private localDomain;
    uint32 private remoteDomain;
    address private transporterAddress;
    address private fiddyTokenAddress;
    address private remoteAttestor;

    function setUp() public {
        deployerPrivateKey = vm.envUint("DEPLOYER_KEY");
        ownerAddress = vm.envAddress("OWNER_ADDRESS");

        localDomain = uint32(
            vm.envUint("LOCAL_DOMAIN")
        );

        remoteDomain = uint32(
            vm.envUint("REMOTE_DOMAIN")
        );

        remoteAttestor = vm.envAddress("REMOTE_ATTESTOR_ADDRESS");
    }

    function run() public {
        vm.startBroadcast(deployerPrivateKey);

        FiddyCent fiddyCent = new FiddyCent(ownerAddress);
        fiddyTokenAddress = address(fiddyCent);

        Transporter transporter = new Transporter(localDomain, remoteDomain, remoteAttestor, fiddyTokenAddress);
        transporterAddress = address(transporter);
        fiddyCent.addCCTPMinter(address(transporterAddress));

        vm.stopBroadcast();
    }
}
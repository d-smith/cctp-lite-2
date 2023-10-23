// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {console} from "forge-std/console.sol";
import {stdStorage, StdStorage, Test} from "forge-std/Test.sol";

import {Utils} from "./utils/Utils.sol";
import {FiddyCent} from "../src/Token.sol";

contract BaseSetup is FiddyCent, Test {
    Utils internal utils;
    address payable[] internal users;

    
    address internal alice;
    address internal bob;

    bytes32  contractOwner = keccak256(abi.encodePacked("owner address"));
    address ownerAddress = address(uint160(uint256(contractOwner)));

    constructor() FiddyCent(ownerAddress) {}

    function setUp() public virtual {
        utils = new Utils();
        users = utils.createUsers(2);

        alice = users[0];
        vm.label(alice, "Alice");
        bob = users[1];
        vm.label(bob, "Bob");
    }
}

contract WhenTransferringTokens is BaseSetup {
    uint256 internal maxTransferAmount = 12e18 ether;

    function setUp() public virtual override {
        BaseSetup.setUp();
        console.log("When transferring tokens");
    }

    function transferToken(
        address from,
        address to,
        uint256 transferAmount
    ) public returns (bool) {
        vm.prank(from);
        return this.transfer(to, transferAmount);
    }
}

contract WhenAliceHasSufficientFunds is WhenTransferringTokens {
    using stdStorage for StdStorage;
    uint256 internal mintAmount = maxTransferAmount;

    function setUp() public override {
        WhenTransferringTokens.setUp();
        console.log("When Alice has sufficient funds");
        _mint(alice, mintAmount);
    }

    function itTransfersAmountCorrectly(
        address from,
        address to,
        uint256 transferAmount
    ) public {
        uint256 fromBalanceBefore = balanceOf(from);
        bool success = transferToken(from, to, transferAmount);

        assertTrue(success);
        assertEqDecimal(
            balanceOf(from),
            fromBalanceBefore - transferAmount,
            decimals()
        );
        assertEqDecimal(balanceOf(to), transferAmount, decimals());
    }

    function testTransferAllTokens() public {
        itTransfersAmountCorrectly(alice, bob, maxTransferAmount);
    }

    function testTransferHalfTokens() public {
        itTransfersAmountCorrectly(alice, bob, maxTransferAmount / 2);
    }

    function testTransferOneToken() public {
        itTransfersAmountCorrectly(alice, bob, 1);
    }

    function testTransferWithFuzzing(uint64 transferAmount) public {
        vm.assume(transferAmount != 0);
        itTransfersAmountCorrectly(
            alice,
            bob,
            transferAmount % maxTransferAmount
        );
    }

    function testTransferWithMockedCall() public {
        vm.prank(alice);
        vm.mockCall(
            address(this),
            abi.encodeWithSelector(
                this.transfer.selector,
                bob,
                maxTransferAmount
            ),
            abi.encode(false)
        );
        bool success = this.transfer(bob, maxTransferAmount);
        assertTrue(!success);
        vm.clearMockedCalls();
    }

    // example how to use https://github.com/foundry-rs/forge-std stdStorage
    function testFindMapping() public {
        uint256 slot = stdstore
            .target(address(this))
            .sig(this.balanceOf.selector)
            .with_key(alice)
            .find();
        bytes32 data = vm.load(address(this), bytes32(slot));
        assertEqDecimal(uint256(data), mintAmount, decimals());
    }
}

contract WhenAliceHasInsufficientFunds is WhenTransferringTokens {
    uint256 internal mintAmount = maxTransferAmount - 1e18;

    function setUp() public override {
        WhenTransferringTokens.setUp();
        console.log("When Alice has insufficient funds");
        _mint(alice, mintAmount);
    }


    function testCannotTransferMoreThanAvailable() public {
        vm.expectRevert();
        transferToken(alice, bob, maxTransferAmount);
    }


    function testCannotTransferToZero() public {
        vm.expectRevert();
        transferToken(alice, address(0), mintAmount);
    }
    
}
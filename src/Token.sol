// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin-contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract FiddyCent is ERC20Burnable, Ownable {

    address private cctpMinter;
    constructor(address initialOwner)
        ERC20("FiddyCent", "FDDC")
        Ownable(initialOwner)
    {
        _mint(initialOwner, 10000000 ether); // A cool 10 million initial supply
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function addCCTPMinter(address minter) public onlyOwner {
        cctpMinter = minter;
    }

    function delegateMint(address to, uint256 amount) public {
        if(msg.sender != cctpMinter) revert("wrong sender");
        _mint(to,amount);
    }
}

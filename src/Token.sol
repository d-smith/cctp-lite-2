// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract FiddyCent is ERC20, Ownable {
    constructor(address initialOwner)
        ERC20("FiddyCent", "FDDC")
        Ownable(initialOwner)
    {
        _mint(initialOwner, 10000000 ether); // A cool 10 million initial supply
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }
}

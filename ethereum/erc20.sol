// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

abstract contract SimpleERC20 {
    function totalSupply() public virtual returns (uint256);
    function balanceOf(address who) public virtual returns (uint256);
    function transfer(address to, uint256 value) public virtual returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
}

contract DS1 is SimpleERC20 {
    string public constant name = "DS1-TOKEN";
    string public constant symbol = "DS1";
    uint8 public constant decimals = 18;

    mapping(address => uint256) balances;
    uint256 totalSupply_;

    constructor() {
        totalSupply_ = 1000 * (10**18);
        balances[msg.sender] = totalSupply_;
    }

    function totalSupply() public override view returns (uint256) {
        return totalSupply_;
    }

    function balanceOf(address who) public override view returns (uint256) {
        return balances[who];
    }

    function transfer(address to, uint256 value) public override returns (bool) {
        require(balances[msg.sender] >= value);
        balances[msg.sender] -= value;
        balances[to] += value;
        emit Transfer(msg.sender, to, value);
        return true;
    }
}

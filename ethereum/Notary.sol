pragma solidity ^0.8.4;

contract Notary {

    mapping (address => mapping (bytes32 => uint)) stamps;

    function store(bytes32 hash) public {
        stamps[msg.sender][hash] = block.timestamp;
    }

    function verify(address recipient, bytes32 hash) public constant returns (uint) {
        return stamps[recipient][hash];
    }
}

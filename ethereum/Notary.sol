pragma solidity ^0.8.21;

contract Notary {

    mapping (address => mapping (bytes32 => uint256)) stamps;

    function store(bytes32 hash) public {
        stamps[msg.sender][hash] = block.timestamp;
    }

    function verify(address recipient, bytes32 hash) public view returns (uint) {
        return stamps[recipient][hash];
    }
}

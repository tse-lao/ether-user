pragma solidity ^0.8.6;


contract Store {
    event ItemSet(bytes32 key, bytes32 value);
    
    string public version;
    address public owner;
    mapping (bytes32 => bytes32) public items;
    
    constructor(string memory _version) {
        owner = msg.sender;
        version = _version;
    }
    
    
    
    function setItem(bytes32 key, bytes32 value) external{
        items[key] = value;
        emit ItemSet(key,value);
    }
    
}
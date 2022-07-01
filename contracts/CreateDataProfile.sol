// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.6;

contract Contracts{
    address[] public pool;
    address tokenAddress;
    string public name;
    string private link; 
    address public owner;
    
    
    constructor(address _tokenAddress, string memory _name, string memory _link){
        owner = msg.sender;
        tokenAddress = _tokenAddress;
        name = _name;
        link = _link;
        //we would like it to have a IPFS link. 
    }
    
    function updateInfo(string memory _name, string memory _link) public {
        require(msg.sender == owner, "Only owner can update info");
        name = _name;
        link = _link;
    }

    //this needs to be changed for now. 
    function CreateNewContract(string memory _url, string memory _metadata, uint _minFee) public returns (address){
        //we need to create something so we can invoke a new contract. 
         ShareContract a = new ShareContract(_url, _metadata, _minFee, msg.sender);
         pool.push(address(a));
         return address(a);
    }
    
    function contractSettings(string memory _link) private{
        link = _link;
        
    }
    
    function getSettings() public view returns (string memory) {
        require(msg.sender == owner, "You need to be owner to view this");
        return link;
    } 
    
    function countContracts() public view returns(uint){
        return pool.length;
    }
    
}
contract ShareContract {
    //public variables.
    address public owner;
    address public createdBy;
    string public url;
    string public metadata;
    uint public fee;

    uint256 public totalUsers;
    uint256 public totalWaiting;
    uint256[] waitingList;

    struct Access {
        address user;
        uint256 downloads;
        string link;
        bool access;
        bool waiting;
    }
    Access[] private accessList;
    mapping(address => uint256) addressIndex;

    constructor(string memory _url, string memory _metadata,  uint _minFee, address _owner) payable {
        owner = _owner;
        createdBy = msg.sender;
        url = _url;
        fee = _minFee;
        metadata = _metadata;
        
    }

    
     modifier isOwner() {
        require(msg.sender == owner, "Permission denied! Only owner");
        
        _;
    }

    function grantAccess(address user) public isOwner {
        //check if there is already a profile in place.
        uint256 index = addressIndex[user];
        if (index > 0) {
            accessList[index].waiting = false;
            accessList[index].access = true;
        } else {
            createProfile(user);
        }
    }

    function revokeAccess(address user) public isOwner {
        uint256 index = addressIndex[user];

        Access memory userProfile = accessList[index];
        userProfile.access = false;
        userProfile.waiting = false;

        //updates the accessList.
        accessList[index] = userProfile;
    }

    function askAccess(address user) public {
        uint256 index = addressIndex[user];
        //check if user already has access.
        if (index > 0) {
            //user is already on the list, now check if has access.
            require(
                !accessList[index].access,
                "You already have access to the file."
            );
            addToWaitingList(index);
        }

        createProfile(user);
    }

    function createProfile(address user) private {
        //adding user to the profile.

        require(addressIndex[user] < 1, "users already has an accessprofile");

        require(user != msg.sender, "owner already has access");
        Access memory newProfile = Access(user,0, "here willk the link be place", false, true);

        accessList.push(newProfile);
        addressIndex[user] = accessList.length - 1;

        if (msg.sender != owner) {
            addToWaitingList(accessList.length - 1);
        }
    }

    function addToWaitingList(uint256 user) private {
        require(
            accessList[user].waiting,
            "user already waiting for an acceptance"
        );

        accessList[user].waiting = true;
        totalWaiting +=1;
    }

    function AllWaiting() public view returns (Access[] memory) {
        //get waiting.
        Access[] memory lAccess = new Access[](totalWaiting);
        for (uint256 i = 0; i < accessList.length; i++) {
            if(accessList[i].waiting){
                Access storage lAcc = accessList[i];
                lAccess[i] = lAcc;
            }
           
        }
        return lAccess;
    }

    function AllOnAccess() public view returns (Access[] memory) {
        Access[] memory lAccess = new Access[](totalWaiting);
        for (uint256 i = 0; i < accessList.length; i++) {
            Access storage lAcc = accessList[i];
            lAccess[i] = lAcc;
        }
        return lAccess;
    }
}

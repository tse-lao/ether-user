pragma solidity ^0.8.6;

contract ShareFile {
    address public owner;
    string public url;

    uint256 public totalUsers;
    uint256 public totalWaiting;
    uint256[] waitingList;

    struct Access {
        address user;
        uint256 downloads;
        bool access;
        bool waiting;
    }

    Access[] public accessList;

    mapping(address => uint256) addressIndex;

    constructor(string memory _url) payable {
        owner = msg.sender;
        url = _url;
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
        Access memory newProfile = Access(user, 0, false, true);

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

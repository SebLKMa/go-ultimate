pragma solidity ^0.8.26;

/*
solc --version
solc, the solidity compiler commandline interface
Version: 0.8.26+commit.8a97fa7a.Linux.g++
recommended that the Solidity version should be locked and specified without the ^
*/

contract Basic {

    // Version is the current version of Store.
    // Read getter is automatically exposed.
    // Reads do not cost Gas.
    string public Version;

    // Items is a mapping which will store the key/value data.
    // Do not iterate over maps, costs Gas.
    // You can only search value by key.
    mapping (string => uint256) public Items;

    // ItemSet is an event which will output any updates to the key/value
    // data to the transaction receipt's logs.
    // Ethereum data storage is not free, but static events stored in receipts are free.
    event ItemSet(string key, uint256 value);

    // The constructor is automatically executed when the contract is deployed.
    constructor() {
        Version = "1.1";
    }

    // SetItem is an external-only function which accepts a key/value pair
    // and updates the contract's internal storage accordingly.
    function SetItem(string memory key, uint256 value) external {
        Items[key] = value;
        emit ItemSet(key, value);
    }

}
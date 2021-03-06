// SPDX-License-Identifier: MIT
pragma solidity >=0.5.16 <=0.7.4;

contract SimpleStorage {
    event StorageSet(string _message);

    uint public storedData;

    function set(uint x) public {
        storedData = x + 4;

        emit StorageSet("Data stored successfully!");
    }
}

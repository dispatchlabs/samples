pragma solidity ^0.4.0;

contract TestContract {
    function getMultiReturn() public view returns (int,string,bool) {
        return (1, "test string", true);
    }
}
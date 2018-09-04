pragma solidity ^0.4.24;

contract Deployed {
    uint private _a;

    function setA(uint a) public {
        _a = a;
    }

    function getA() public view returns (uint) {
        return _a;
    }
}
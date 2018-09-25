pragma solidity ^0.4.24;

contract Deployed {
<<<<<<< HEAD
    uint private _a;

    function setA(uint a) public {
        _a = a;
    }

    function getA() public view returns (uint) {
=======
    uint256 private _a;

    function setA(uint256 a) public {
        _a = a;
    }

    function getA() public view returns (uint256) {
>>>>>>> origin/dev
        return _a;
    }
}
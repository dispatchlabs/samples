pragma solidity ^0.4.24;

contract ProblemContract_DG_304 {
    function test_log4() public returns(uint) {
        assembly {
            mstore(0, 1)
            mstore(0x20, 2)
            log4(0, 0x20, "x", "y", "z", "a")
        }
    }
}

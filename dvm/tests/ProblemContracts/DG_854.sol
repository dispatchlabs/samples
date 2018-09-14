pragma solidity ^0.4.24;

contract ProblemContract_DG_854 {
    function test_balance() public view  returns(uint256) {
        assembly {
            mstore(0, balance(origin))
            return(0, 0x100)
        }
    }
}

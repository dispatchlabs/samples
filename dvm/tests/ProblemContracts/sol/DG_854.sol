pragma solidity ^0.4.24;

contract ProblemContract_DG_854 {
    function test_balance() public view  returns(uint256) {
        assembly {
            mstore(0, balance(origin)) // origin - transaction sender
            return(0, 0x100) // end execution, return data mem[p…(p+s))
        }
    }
}

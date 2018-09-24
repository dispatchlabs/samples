pragma solidity ^0.4.24;

contract ProblemContract_DG_304 {
    function test_timestamp() public view returns(uint) {
        assembly {
            mstore(0, timestamp)
            return(0, 0x20)
        }
    }
}

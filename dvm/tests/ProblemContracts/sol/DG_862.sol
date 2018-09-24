pragma solidity ^0.4.24;

contract ContractTest {
    function test_selfdestruct() public returns(uint) {
        assembly {
            selfdestruct(origin)
        }
    }
}

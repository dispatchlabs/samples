pragma solidity ^0.4.24;

contract ContractTest_862 {
    function test_selfdestruct() public returns(uint) {
        assembly {
            selfdestruct(origin)
        }
    }

    function test_string() public pure returns (string ret) {
      ret = "testing";
    }
}

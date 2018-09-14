// Tests contract -> contract opcodes. This is the created contract that gets
// called by call, callcode, delegatecall, and staticcall
pragma solidity ^0.4.24;

contract Callee {
  address public _sender;

  function set_sender() public returns(address) {
    _sender = msg.sender;

    return msg.sender;
  }

  function get_sender() public view returns(address) {
    return msg.sender;
  }
}

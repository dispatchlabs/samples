// Tests contract -> contract opcodes
pragma solidity ^0.4.24;


contract ContractTest {
  address public _sender;
  address public _created_contract;

  function _create(bytes code) internal returns (address addr) {
    assembly {
      addr := create(0,add(code,0x20), mload(code))
      if iszero(extcodesize(addr)) {
        mstore(0, 0)
        return(0, 0x20)
      }
    }
  }

  // create(v, p, s)
  // just make a contract
  function test_create() public returns(bool) {
    _created_contract = _create("608060405234801561001057600080fd5b50610202806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063a8d509ff1461005c578063bb02b363146100b3578063cbd494621461010a575b600080fd5b34801561006857600080fd5b50610071610161565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100bf57600080fd5b506100c8610169565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561011657600080fd5b5061011f6101b1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600033905090565b6000336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058208bf2762c00b9623763de955c67e7ad47e2a36e15c21dc0ed5f2b10e86ae612260029");

    return true;
  }

  // call(g, a, v, in, insize, out, outsize)
  // set sender, should not be same as self
  // storge modifications should not be local
  function test_call() public returns(bool) {
    _sender = msg.sender;

    _created_contract.call(bytes4(keccak256("set_sender()")));

    return _sender != msg.sender;
  }

  // callcode(g, a, v, in, insize, out, outsize)
  // print out the caller, should not be the same as self
  // change storage and make sure it's being modified locally
  function test_callcode() public returns(bool) {
    _created_contract.callcode(bytes4(keccak256("set_sender()")));

    return _sender != msg.sender;
  }

  // delegatecall(g, a, in, insize, out, outsize)
  // print out the caller ... make sure it isn't the calling contract
  // change storage and make sure it is being modified locally
  function test_delegatecall() public returns(bool) {
    _created_contract.delegatecall(bytes4(keccak256("set_sender()")));

    return _sender == msg.sender;
  }

  // staticcall(g, a, in, insize, out, outsize)
  // just use this normally
  // function test_staticcall() public returns(bool) {
  //   _sender = address(this);

  //   address self_sender = _created_contract.staticcall(bytes4(keccak256("get_sender()")));

  //   return self_sender == address(this);
  // }

  // staticcall(g, a, in, insize, out, outsize)
  // try to change data and hopefully fail ... do this by calling 
  // one of the callcode methods above
  // function test_staticcall_failure() public returns(bool) {
  //   _created_contract.staticall(bytes4(keccak256("set_sender()")));

  //   // should never make it here
  //   return true;
  // }
}

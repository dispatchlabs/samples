pragma solidity ^0.4.24;

contract ContractTest {
 address public _sender;
 address public _created_contract;

  function _create(bytes code) public returns (address addr) {
    assembly {
      addr := create(0, add(code,0x20), mload(code))
      if iszero(extcodesize(addr)) {
        revert(0, 0)
      }
    }
  }

  function test_create() public returns (address) {
    _created_contract = _create(
      hex"608060405234801561001057600080fd5b50610202806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063a8d509ff1461005c578063bb02b363146100b3578063cbd494621461010a575b600080fd5b34801561006857600080fd5b50610071610161565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100bf57600080fd5b506100c8610169565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561011657600080fd5b5061011f6101b1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600033905090565b6000336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820fc216cc5ca21176a44db8070649b6e74359ed0e5480781d3dfe8b9c5f82b59220029"
    );

    require(address(_created_contract) != 0);

    return _created_contract;
  }

  // call(g, a, v, in, insize, out, outsize)
  // set sender, should not be same as self
  // storge modifications should not be local
  function test_call() public returns(bool) {
    bytes4 sig = bytes4(keccak256("get_sender()"));
    address addr = _created_contract;
    address sender = 0;

    assembly {
      let p := mload(0x40)
      mstore(p, sig)
      
      if iszero(call(5000, addr, 0, p, 0x04, p, 0x20)) {
        revert(0, 0)
      }

      sender := mload(p)
    }

    return sender == address(this);
  }
  
  // callcode(g, a, v, in, insize, out, outsize)
  // print out the caller, should not be the same as self
  // change storage and make sure it's being modified locally
  function test_callcode() public returns(bool) {
    require(_created_contract.callcode(bytes4(keccak256("set_sender()"))));
    return _sender != msg.sender && _sender == address(this);
  }
  
  // delegatecall(g, a, in, insize, out, outsize)
  // print out the caller ... make sure it isn't the calling contract
  // change storage and make sure it is being modified locally
  function test_delegatecall() public view returns(bool) {
    require(_created_contract.delegatecall(bytes4(keccak256("set_sender()"))));

    return _sender == msg.sender;
  }

  // staticcall(g, a, in, insize, out, outsize)
  // just use this normally
  function test_staticcall() public returns(bool) {
    address sender;
    bytes4 sig = bytes4(keccak256("get_sender()"));
    address addr = _created_contract;

    assembly {
      let p := mload(0x40)
      mstore(p, sig)
      
      if iszero(staticcall(5000, addr, p, 0x04, p, 0x20)) {
        revert(0, 0)
      }

      sender := mload(p)
    }

    return sender == address(this);
  }
  
  // staticcall(g, a, in, insize, out, outsize)
  // try to change data and hopefully fail ... do this by calling
  // one of the callcode methods above
  function test_staticcall_failure() public view returns(bool) {
    address sender;
    bytes4 sig = bytes4(keccak256("set_sender()"));
    address addr = _created_contract;

    assembly {
      let p := mload(0x40)
      mstore(p, sig)
      
      if iszero(staticcall(500000, addr, p, 0x04, p, 0x20)) {
        revert(0, 0)
      }

      sender := mload(p)
    }

    // should never make it here
    return true;
  }
}


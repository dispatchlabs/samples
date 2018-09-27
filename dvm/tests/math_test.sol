// Tests these opp codes
// ADD
// ADDMOD
// AND
// BYTE
// CALLCODE
// CALLDATACOPY
// CALLDATALOAD
// CALLDATASIZE
// CALLVALUE
// CODECOPY
// DIV
// DUP1
// DUP11
// DUP2
// DUP3
// DUP4
// DUP5
// DUP6
// EQ
// EXP
// GT
// ISZERO
// JUMP
// JUMPDEST
// JUMPI
// KECCAK256
// LOG1
// LT
// MLOAD
// MOD
// MSTORE
// MUL
// MULMOD
// NOT
// OR
// POP
// PUSH1
// PUSH17
// PUSH2
// PUSH29
// PUSH31
// PUSH32
// PUSH4
// PUSH5
// PUSH6
// RETURN
// REVERT
// SAR
// SDIV
// SGT
// SHL
// SHR
// SIGNEXTEND
// SLT
// SMOD
// STOP
// SUB
// SWAP1
// SWAP2
// SWAP3
// SWAP4
// XOR

pragma solidity ^0.4.0;

contract MathTest {
  function test_add(uint x, uint y) public pure returns (uint ret) {
    assembly {          
      ret := add(x, y)
    }
  }

  function test_sub(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := sub(x, y)
    }
  }

  function test_mul(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := mul(x, y)
    }
  }

  function test_div(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := div(x, y)
    }
  }

  function test_sdiv2() public pure returns (int ret) {
    assembly {
      ret := sdiv(0xf1, 0xfd)
    }
  }

  function test_sdiv(int x, int y) public pure returns (int ret) {
    assembly {
      ret := sdiv(x, y)
    }
  }

  function test_mod(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := mod(x, y)
    }
  }

  function test_smod(int x, int y) public pure returns (int ret) {
    assembly {
      ret := smod(x, y)
    }
  }

  function test_exp(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := exp(x, y)
    }
  }

  function test_not(bool x) public pure returns (bool ret) {
    assembly {
      ret := not(x)
    }
  }

  function test_lt(uint x, uint y) public pure returns (bool ret) {
    assembly {
      ret := lt(x, y)
    }
  }

  function test_gt(uint x, uint y) public pure returns (bool ret) {
    assembly {
      ret := gt(x, y)
    }
  }

  function test_slt(int x, int y) public pure returns (bool ret) {
    assembly {
      ret := slt(x, y)
    }
  }

  function test_sgt(int x, int y) public pure returns (bool ret) {
    assembly {
      ret := sgt(x, y)
    }
  }

  function test_eq(uint x, uint y) public pure returns (bool ret) {
    assembly {
      ret := eq(x, y)
    }
  }

  function test_iszero(uint x) public pure returns (bool ret) {
    assembly {
      ret := iszero(x)
    }
  }

  function test_and(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := and(x, y)
    }
  }

  function test_or(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := or(x, y)
    }
  }

  function test_xor(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := xor(x, y)
    }
  }

  function test_byte(uint256 index) public pure returns (bytes32 ret) {
    assembly {
      ret := byte(index, "123456789foobarfoo")
    }
  }

  function test_shl(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := shl(x, y)
    }
  }

  function test_shr(uint x, uint y) public pure returns (uint ret) {
    assembly {
      ret := shr(x, y)
    }
  }

  function test_sar(int x, int y) public pure returns (int ret) {
    assembly {
      ret := sar(x, y)
    }
  }

  function test_addmod(uint x, uint y, uint m) public pure returns (uint ret) {
    assembly {
      ret := addmod(x, y, m)
    }
  }

  function test_mulmod(uint x, uint y, uint m) public pure returns (uint ret) {
    assembly {
      ret := mulmod(x, y, m)
    }
  }

  function test_signextend(int32 original) public pure returns (bytes32) {
    assembly {
      mstore(0, signextend(0x10, original))
      return(0, 0x64)
    }
  }

  function test_keccak256() public pure returns (bytes32 ret) {
    assembly {
      mstore(0, "this is a test")
      ret := keccak256(0, 0x0e)
    }
  } 

  /* this gets converted to keccak256 automatically however in the EVM and DVM
  * it's calles opSha3 (inconsistent) */
  function test_sha3() public pure returns (bytes32 ret) {
    assembly {
      mstore(0, "this is a test")
      ret := keccak256(0, 0x0e)
    }
  }
}

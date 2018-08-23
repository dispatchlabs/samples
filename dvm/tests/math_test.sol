// Tests these opp codes
// ADD
// ADDMOD
// AND
// BALANCE
// CALLDATACOPY
// CALLDATALOAD
// CALLDATASIZE
// CALLVALUE
// CODECOPY
// DIV
// DUP1
// DUP2
// DUP3
// DUP4
// DUP5
// DUP6
// EQ
// EXP
// GT
// INVALID
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
// PUSH19
// PUSH2
// PUSH29
// PUSH31
// PUSH32
// PUSH4
// PUSH6
// RETURN
// REVERT
// SDIV
// SGT
// SLT
// SMOD
// STOP
// SUB
// SWAP1
// SWAP2
// SWAP3
// SWAP4
// SWAP5
// XOR

pragma solidity ^0.4.0;

contract MathTest {
				function add(uint x, uint y) public pure returns (uint) {
								return x + y;
				}

				function sub(uint x, uint y) public pure returns (uint) {
								return x - y;
				}

				function mul(uint x, uint y) public pure returns (uint) {
								return x * y;
				}

				function div(uint x, uint y) public pure returns (uint) {
								return x / y;
				}

				function sdiv(int x, int y) public pure returns (int) {
								return x / y;
				}

				function mod(uint x, uint y) public pure returns (uint) {
								return x % y;
				}

				function smod(int x, int y) public pure returns (int) {
								return x % y;
				}

				function exp(uint x, uint y) public pure returns (uint) {
								return x ** y;	
				}

				function not(bool x) public pure returns (bool) {
								return !x;
				}

				function lt(uint x, uint y) public pure returns (bool) {
								return x < y;
				}

				function gt(uint x, uint y) public pure returns (bool) {
								return x > y;
				}

				function slt(int x, int y) public pure returns (bool) {
								return x < y;
				}

				function sgt(int x, int y) public pure returns (bool) {
								return x > y;
				}

				function eq(uint x, uint y) public pure returns (bool) {
								return x == y;
				}

				function iszero(uint x) public pure returns (bool) {
								return x == 0;
				}

				function and(uint x, uint y) public pure returns (uint) {
								return x & y;
				}

				function or(uint x, uint y) public pure returns (uint) {
								return x | y;
				}

				function xor(uint x, uint y) public pure returns (uint) {
								return x ^ y;
				}

				/* This doesn't trigger the opcode */
				function get_byte(bytes b) public pure returns (byte) {
								return b[4];
				}

				function shl(uint x, uint y) public pure returns (uint) {
								return x << y;				
				}

				function shr(uint x, uint y) public pure returns (uint) {
								return x >> y;
				}

				function sar(int x, int y) public pure returns (int) {
								return x >> y;							
				}

				function addmod_test(uint x, uint y, uint m) public pure returns (uint) {
								return addmod(x, y, m);
				}

				function mulmod_test(uint x, uint y, uint m) public pure returns (uint) {
								return mulmod(x, y, m);
				}

				/* this doesn't generate the right op code for some reason */
				function signextend(int i, int x) public pure returns (int) {
								x >> i;
				} 

				/* this is automatically executed when the contract is run
				function keccak256(uint p, uint n) public pure returns (uint) {
								_;
				} */

				/* this gets converted to keccak256 automatically */
				function test_sha3(bytes b) public pure returns (bytes32) {
								return sha3(b);
				}
}

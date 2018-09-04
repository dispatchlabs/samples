// Tests these opp codes


pragma solidity ^0.4.0;

library GetCode {
    function at(address _addr) public view returns (bytes o_code) {
        assembly {
            // retrieve the size of the code, this needs assembly
            let size := extcodesize(_addr)
            // allocate output byte array - this could also be done without assembly
            // by using o_code = new bytes(size)
            o_code := mload(0x40)
            // new "memory end" including padding
            mstore(0x40, add(o_code, and(add(add(size, 0x20), 0x1f), not(0x1f))))
            // store length in memory
            mstore(o_code, size)
            // actually retrieve the code, this needs assembly
            extcodecopy(_addr, add(o_code, 0x20), 0, size)
				}
		}
}

contract NonMathTest {
				// F	jump to label / code position
				function jump() public view returns(uint) {
								assembly {
												let n := calldataload(4)
												let a := 1
												let b := a
												loop:
																jump(loopend)
												a add swap1
												n := sub(n, 1)
												jump(loop)
												loopend:
																mstore(0, a)
												return(0, 0x20)
								}
				}

				// F	jump to label if cond is nonzero
				function jumpi() public view returns(uint) {
								assembly {
												let n := calldataload(4)
												let a := 1
												let b := a
												loop:
																jumpi(loopend, eq(n, 0))
												a add swap1
												n := sub(n, 1)
												jump(loop)
												loopend:
																mstore(0, a)
												return(0, 0x20)
								}
				}

				// F	current position in code
				function pc() public view returns(uint) {
								assembly {
												mstore(0, pc)
												return(0, 0x20)
								}
				}

				/* this gets tested in almost any context since it's crucial
				* to maintainig the stack
				// F	remove the element pushed by x
				function pop(x)) public view returns(uint) {
				} */

				/* These are automatically tested by regular contract use
				// F	copy ith stack slot to the top (counting from top)
				function dup1 … dup16 public view returns(uint) {
				} */

				/* these are also automatically tested by regular contract use */
				// F	swap topmost and ith stack slot below it
				/*
				function swap1 … swap16 public view returns(uint) {
				} */

				// F	mem[p..(p+32))
				function mload() public view returns(uint) {
								assembly {
												let x := 20
												mstore(0x20, add(mload(0x80), 3))

												return(0, 0x20)
								}
				}

				/* tested above 
				// F	mem[p..(p+32)) := v
				function mstore(p, v)) public view returns(uint) {
				} */

				// F	mem[p] := v & 0xff (only modifies a single byte)
				// function mstore8() public view returns(uint) {
				// 				// thi is a tiny bit tricky because we need 
				// 				// a number spanning multiple bytes (a short)
				// 				// and to 
				// 				assembly {
				// 								let x := 0xffff
				// 								mstore(0x20, add(mstore8(0x20), 0))

				// 								return(0, 0x20)
				// 				}
				// }

				// F	storage[p]
				// function sload(p)) public view returns(uint) {
				// }

				// // F	storage[p] := v
				// function sstore(p, v)) public view returns(uint) {
				// }

				// F	size of memory, i.e. largest accessed memory index
				function msize() public view returns(uint) {
								assembly {
												mstore(0, msize)
												return(0, 0x20)
								}
				}

				// F	gas still available to execution
				function gas() public view returns(uint) {
								assembly {
												mstore(0, gas)
												return(0, 0x20)
								}
				}

				// F	address of the current contract / execution context
				function test_address() public view returns(uint) {
								assembly {
												mstore(0, address)
												return(0, 0x20)
								}
				}

				// F	wei balance at address a
				function balance() public view returns(uint) {
								assembly {
												mstore(0, balance(address))
												return(0, 0x20)
								}
				}

				// F	call sender (excluding delegatecall) 
				function caller()	public view returns(uint) {
								assembly {
												mstore(0, caller)
												return(0, 0x20)
								}
				}

				// F	wei sent together with the current call
				function callvalue() public view returns(uint) {
								assembly {
												mstore(0, callvalue)
												return(0,0x20)
								}
				}

				// F	call data starting from position p (32 bytes)
				// function calldataload(p)) public view returns(uint) {
				// }

				// // F	size of call data in bytes
				// function calldatasize public view returns(uint) {
				// }

				// // F	copy s bytes from calldata at position f to mem at position t
				// function calldatacopy(t, f, s)) public view returns(uint) {
				// }

				// F	size of the code of the current contract / execution context
				function codesize() public view returns(uint) {
								assembly {
												mstore(0, codesize)
												return(0,0x20)
								}
				}

				// F	copy s bytes from code at position f to mem at position t
				// function codecopy(t, f, s)) public view returns(uint) {
				// }

				// F	size of the code at address a
				function extcodesize() public view returns(uint) {
								assembly {
												mstore(0, extcodesize(address))
												return(0, 0x20)
								}
				}

				// F	like codecopy(t, f, s) but take code at address a
				function extcodecopy() public view returns(bytes) {
								return GetCode.at(address(this));
				}

				// B	size of the last returndata
				function returndatasize() public view returns(uint) {

				}

				// B	copy s bytes from returndata at position f to mem at position t
				// function returndatacopy(t, f, s) public view returns(uint) {
				// }

				// F	create new contract with code mem[p..(p+s)) and send v wei and return the new address
				// function create(v, p, s)) public view returns(uint) {
				// }

				// C	create new contract with code mem[p..(p+s)) at address keccak256(<address> . n . keccak256(mem[p..(p+s))) and send v wei and return the new address
				// function create2(v, n, p, s) public view returns(uint) {
				// }

				// F	call contract at address a with input mem[in..(in+insize)) providing g gas and v wei and output area mem[out..(out+outsize)) returning 0 on error (eg. out of gas) and 1 on success
				// function call(g, a, v, in, insize, out, outsize)) public view returns(uint) {
				// }

				// F	identical to call but only use the code from a and stay in the context of the current contract otherwise
				// function callcode(g, a, v, in, insize, out, outsize)) public view returns(uint) {
				// }

				// H	identical to callcode but also keep caller and callvalue
				// function delegatecall(g, a, in, insize, out, outsize) public view returns(uint) {
				// }

				// B	identical to call(g, a, 0, in, insize, out, outsize) but do not allow state modifications
				// function staticcall(g, a, in, insize, out, outsize) public view returns(uint) {
				// }

				/* this is naturally tested everywhere */
				// F	end execution, return data mem[p..(p+s))
				/*
				function return(p, s)) public view returns(uint) {
				} */

				// B	end execution, revert state changes, return data mem[p..(p+s))
				// function revert(p, s) public view returns(uint) {
				// }

				// F	end execution, destroy current contract and send funds to a
				function test_selfdestruct() public view returns(uint) {
				}

				// F	end execution with invalid instruction
				function invalid() public view returns(uint) {
								assembly {
												invalid
								}
				}

				// F	log without topics and data mem[p..(p+s))
				/* already implicitly tested 
				function test_log0() public view returns(uint) {
				} */

				// F	test_log with topic t1 and data mem[p..(p+s))
				function test_log1() public view returns(uint) {
								assembly {
												let x:= 1
												let y:= 2
												log1(0, 0x40, "x")
								}
				}

				// F	test_log with topics t1, t2 and data mem[p..(p+s))
				function test_log2() public view returns(uint) {
								assembly {
												let x:= 1
												let y:= 2
												log2(0, 0x40, "x", "y")
								}
				}

				// F	test_log with topics t1, t2, t3 and data mem[p..(p+s))
				function test_log3() public view returns(uint) {
								assembly {
												let x:= 1
												let y:= 2
												log3(0, 0x40, "x", "y", "z")
								}
				}

				// F	test_log with topics t1, t2, t3, t4 and data mem[p..(p+s))
				function test_log4() public view returns(uint) {
								assembly {
												let x:= 1
												let y:= 2
												log4(0, 0x40, "x", "y", "z", "a")
								}
				}

				// F	transaction sender
				function origin() public view returns(uint) {
								assembly {
												mstore(0, origin)
												return(0, 0x20)
								}
				}

				// F	gas price of the transaction
				function gasprice() public view returns(uint) {
								assembly {
												mstore(0, gasprice)
												return(0, 0x20)
								}
				}

				// F	hash of block nr b - only for last 256 blocks excluding current
				function test_blockhash() public view returns(uint) {
								assembly {
												mstore(0, blockhash(number))
												return(0, 0x20)
								}
				}

				// F	current mining beneficiary
				function coinbase() public view returns(uint) {
								assembly {
												mstore(0, coinbase)
												return(0, 0x20)
								}
				}

				// F	timestamp of the current block in seconds since the epoch
				function timestamp() public view returns(uint) {
								assembly {
												mstore(0, timestamp)
												return(0, 0x20)
								}
				}

				// F	current block number
				function number() public view returns(uint) {
								assembly {
												mstore(0, number)
												return(0, 0x20)
								}
				}

				// F	difficulty of the current block
				function difficulty() public view returns(uint) {
								assembly {
												mstore(0, difficulty)
												return(0, 0x20)
								}
				}
}

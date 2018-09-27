pragma solidity ^0.4.24;

contract PrecompiledTest {
  function expmod(uint256 base, uint256 e, uint256 m) public returns (uint256 o) {
        // are all of these inside the precompile now?

        assembly {
            // define pointer
            let p := mload(0x40)
            
            // store data assembly-favouring ways
            mstore(p, 0x20)             // Length of Base
            mstore(add(p, 0x20), 0x20)  // Length of Exponent
            mstore(add(p, 0x40), 0x20)  // Length of Modulus
            mstore(add(p, 0x60), base)  // Base
            mstore(add(p, 0x80), e)     // Exponent
            mstore(add(p, 0xa0), m)     // Modulus
            
            // call modexp precompile! -- old school gas handling
            if iszero(call(gas, 0x05, 0, p, 0xc0, p, 0x20)) {
                revert(0, 0)
            }

            // data
            o := mload(p)
        }
    }

    // should return a value
    function domod() public returns (uint256) {
        return expmod(13, 13, 12);
    }
}
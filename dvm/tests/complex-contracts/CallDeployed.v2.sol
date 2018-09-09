pragma solidity ^0.4.0;

contract CallDeployed  {
    function setAProxy(address originalContract, uint256 a) public {
        originalContract.call(bytes4(keccak256("setA(uint256)")), a);
        // originalContract.call(bytes4(sha3("setA(uint256)")), a);
    }

    function getAPrxoy(address originalContract) public view returns (uint) {        
        uint rv = 0;
        
        //originalContract.call(bytes4(keccak256("getA()")));

        return rv;
    }
}
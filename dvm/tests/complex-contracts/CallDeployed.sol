pragma solidity ^0.4.24;

contract Deployed {
    function setA(uint a) public;
    function getA() public view returns (uint);
}

contract CallDeployed  {
    function setAProxy(address originalContract, uint a) public {
        Deployed dc = Deployed(originalContract);
        dc.setA(a);
    }
    
    function getAPrxoy(address originalContract) public view returns (uint) {
        Deployed dc = Deployed(originalContract);
        return dc.getA();
    }
}
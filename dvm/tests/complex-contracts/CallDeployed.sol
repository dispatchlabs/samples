pragma solidity ^0.4.24;

interface Deployed {
    function setA(uint a) external;
    function getA() external view returns (uint);
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
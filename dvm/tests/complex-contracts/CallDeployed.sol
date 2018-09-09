pragma solidity ^0.4.24;

interface Deployed {
    function setA(uint a) external;
    function getA() external view returns (uint256);
}

contract CallDeployed  {
    function setAProxy(address originalContract, uint256 a) public {
        Deployed dc = Deployed(originalContract);
        dc.setA(a);
    }

    function getAProxy(address originalContract) public view returns (uint256) {
        Deployed dc = Deployed(originalContract);
        return dc.getA();
    }
}
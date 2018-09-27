pragma solidity ^0.4.24;

contract ProblemContract_DG_702_2 {
    function sendAmountTo(address receiver, uint256 amount) public payable {
        receiver.transfer(amount);
    }
}
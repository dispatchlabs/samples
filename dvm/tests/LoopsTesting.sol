pragma solidity ^0.4.0;

contract LoopsTesting {
    uint constant OneH = 100;
    uint constant OneT = 1000;
    uint constant OneM = 1000000;
    uint constant OneB = 1000000000;

    function IncHundredTimes() public pure returns (uint256) {
        uint256 val = 0;

        for (uint i = 0; i < OneH; i++) {
            val++;
        }

        return val;
    }

    function IncThousandTimes() public pure returns (uint256) {
        uint256 val = 0;

        for (uint i = 0; i < OneT; i++) {
            val++;
        }

        return val;
    }

    function IncMilTimes() public pure returns (uint256) {
        uint256 val = 0;

        for (uint i = 0; i < OneM; i++) {
            val++;
        }

        return val;
    }

    function IncBilTimes() public pure returns (uint256) {
        uint256 val = 0;

        for (uint i = 0; i < OneB; i++) {
            val++;
        }

        return val;
    }

    function IncBilTimesForFor() public pure returns (uint256) {
        uint256 val = 0;

        for (uint i = 0; i < OneB; i++) {
            for (uint j = 0; j < OneB; j++) {
                val++;
            }
        }

        return val;
    }

    function IncBilTimesForForFor() public pure returns (uint256) {
        uint256 val = 0;

        for (uint i = 0; i < OneB; i++) {
            for (uint j = 0; j < OneB; j++) {
                for (uint k = 0; k < OneB; k++) {
                    val++;
                }
            }
        }
        
        return val;
    }

    function IncInfiniteTimes() public pure returns (uint256) {
        uint256 val = 0;

        for (;;) {
            val++;
            val--;
        }
        
        return val;
    }
}
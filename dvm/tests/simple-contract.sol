pragma solidity ^0.4.0;
contract TestContract {

    struct ComplexStruct {
        uint var1;
        bool var2;
        uint8 var3;
        string var4;
    }

    string public var5;
    ComplexStruct public var6;

    constructor() public {
        var5 = "aaaaaaaaaaaaa";
        var6.var1 = 0;
        var6.var2 = false;
        var6.var3 = 1;
        var6.var4 = "bbbbbbbbbbb";
    }

    function setVar5(string value) public {
        var5 = value;
    }

    function setVar6Var4(string value) public {
        var6.var4 = value;
    }
    
    function incVar6Var1() public {
        var6.var1++;
    }
    
    function getVar5() public view returns (string) {
        return var5;
    }
}
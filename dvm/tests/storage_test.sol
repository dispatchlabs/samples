// Tests contract -> contract opcodes
pragma solidity ^0.4.24;


contract StorageTest {
  mapping (string => bytes) private _data;

  function get(string key) public view returns(bytes) {
    return _data[key];
  }

  function set(string key, bytes value) public returns(bool) {
    _data[key] = value;

    return true;
  }
}

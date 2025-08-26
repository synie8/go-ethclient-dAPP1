// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Counter {
    int256 private _count; 

    constructor(){
        _count=0;
    }

    function increment() public returns (int256){
        return _count++;
    }

    function decrement() public returns (int256){
        return _count--;
    }

    function reset() public returns (int256){
        _count = 0;
        return _count;
    }
}
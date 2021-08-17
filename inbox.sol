// SPDX-License-Identifier: GPL-3.0
//简单一个合约
pragma solidity >=0.4.16 <0.8.8;//这里我用的0.8.7  这里安装solc版本必须⚠️,不然容易出错

contract Inbox {
    string public message;

    constructor(string memory initialMessage){
        message = initialMessage;
    }

    function setMessage(string memory newMessage) public {
        message = newMessage;
    }
}

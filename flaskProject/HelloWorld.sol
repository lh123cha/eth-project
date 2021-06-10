// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0; // 编译指令，指明solidity版本，上面一行是自由软件许可证

// 合约
contract HelloWorld {

    // 函数, prue 关键字告诉编译器该函数不读也不写状态变量
    function say() public pure returns (string memory) {
        return "Hello world!";
    }
}


// SPDX-License-Identifier: GPL-3.0 
pragma solidity >=0.4.16 <0.9.0; // 编译指令，指明solidity版本，上面一行是自由软件许可证

// 合约
contract WaiMai {   
    // 玩家结构体
    struct User {
        address addr;       // address
        uint id;        //id
        string name;    // name
        string tel;     // telephone
        uint money;       // money
        string dept;    // dept
    }

    uint num;   // user数目
    uint[] ids; //id数组
    User[] users;   // 存储所有student的数据
    mapping(address => uint) userId;  // 地址到id的映射
    mapping(address => bool) isPlay;    // 判断某个地址（即账户）是否参与
    //event Insert(uint id);

    constructor() {
        num = 0;
    }

    function insert_user(string memory _name,string memory _tel,string memory _dept) public {
        require(!isPlay[msg.sender]);
        users.push(User(msg.sender, num,_name, _tel, 100, _dept));
        ids.push(num);
        userId[msg.sender] = num;
        num += 1;
        //uint id = _id;
        //emit Insert(id);
    }

    function select_count() public view returns (uint n){
        return num;
    }

    function select_all_id() public view returns (uint[] memory) {
        require(num>0);
        return ids;
    }

    function select_id(uint _id) public view returns (User memory){
        require(num>0);
        //require(isPlay[msg.sender]);
        for (uint i = 0; i < num; ++i) {
            if (users[i].id == _id) {
                return users[i];
            }
        }
        return User(msg.sender, 0,"null", "null", 0, "null");
    }


}

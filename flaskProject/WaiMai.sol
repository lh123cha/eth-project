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

    struct Deal{
        uint id;       //id
        string text;   //text
        address send_user;  //senduser address
        uint money;     //money
        address receive_user;   //receiveuser address
        bool is_finish;  //whether deal is finish
        uint left_hour;  //hour time left
        bool is_timeout;  //whether time is out
        string tip; //tip of the deal

    }

    uint num;   // user数目
    uint[] ids; //id数组
    User[] users;   // 存储所有user的数据
    mapping(address => uint) userId;  // 地址到id的映射
    mapping(address => bool) isPlay;    // 判断某个地址（即账户）是否参与
    //event Insert(uint id);
    uint deal_num; //deal数目
    uint[] deal_ids;
    Deal[] deals;
    
    User current_user = User(address(0x0), 0,"null", "null", 0, "null");
    User deal_user = User(address(0x0), 0,"null", "null", 0, "null");
    Deal current_deal = Deal(0,"null",address(0x0),0,address(0x0),false,0,false,"null");


    constructor() {
        num = 0;
        deal_num = 0;
    
    }

    //test智能合约用来测试是否连上了以太坊，是测试用例
    function test() public view returns (uint n){
        return num;
    }


    //insert user
    function insert_user(string memory _name,string memory _tel,string memory _dept) public {
        require(!isPlay[msg.sender]);
        users.push(User(msg.sender, num,_name, _tel, 100, _dept));
        ids.push(num);
        userId[msg.sender] = num;
        isPlay[msg.sender] = true;
        num += 1;
        //uint id = _id;
        //emit Insert(id);
    }
 
    //insert deal
    function insert_deal(string memory _text,uint _money,uint _time,string memory _tip) public {
        //require(isPlay[msg.sender]);
        for (uint i = 0; i < num; ++i) {
            if (users[i].addr == msg.sender) {
                current_user = users[i];
            }
        }
        deals.push(Deal(deal_num,_text,msg.sender,_money,address(0x0),false,_time,false,_tip));
        deal_ids.push(deal_num);
        deal_num += 1;
    }

    //get all deals
    function select_all() public view returns (Deal[] memory) {
        require(deal_num>0);
        return deals;
    }


    //finish a deal
    function finish_deal(uint _id) public{
        require(deal_num>0);
        require(isPlay[msg.sender]);

        for (uint i = 0; i < deal_num; ++i) {
            if (deals[i].id == _id) {
                current_deal = deals[i];
                require(deals[i].is_finish == false);
                require(deals[i].is_timeout == false);
                deals[i].receive_user = msg.sender;
                deals[i].is_finish = true;
            }
        }

        require(current_deal.send_user!=msg.sender);

        uint deal_money = current_deal.money;


        for (uint i = 0; i < num; ++i) {
            if (users[i].addr == msg.sender) {
                current_user = users[i];
                require(users[i].money>deal_money);
                users[i].money = users[i].money - deal_money;
            }
        }

        for (uint i = 0; i < num; ++i) {
            if (users[i].addr == current_deal.send_user) {
                deal_user = users[i];
                users[i].money = users[i].money + deal_money;
            }
        }
    }


     function update_dealtime(uint _time_minus_hour) public{
        for (uint i = 0; i < deal_num; ++i) {
            if(deals[i].left_hour>1){
                deals[i].left_hour = deals[i].left_hour - _time_minus_hour;
            }else{
                deals[i].is_timeout = true;
                deals[i].is_finish = true;
            }
            // if(deals[i].left_hour<0){
            //     deals[i].is_timeout = true;
            //     deals[i].is_finish = true;
            // }
        }
     }





//unused
    function select_count() public view returns (uint n){
        return num;
    }

    function select_all_id() public view returns (uint[] memory) {
        require(num>0);
        return ids;
    }

    function select_id(uint _id) public view returns (User memory){
        require(num>0);
        require(isPlay[msg.sender]);
        for (uint i = 0; i < num; ++i) {
            if (users[i].id == _id) {
                return users[i];
            }
        }
        return User(msg.sender, 0,"null", "null", 0, "null");
    }


}

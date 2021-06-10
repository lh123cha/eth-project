// SPDX-License-Identifier: GPL-3.0 
pragma solidity >=0.4.16 <0.9.0; // 编译指令，指明solidity版本，上面一行是自由软件许可证

// 合约
contract WaiMai {   
    // 玩家结构体
    struct Student {
        uint id;        // student id
        string name;    // student's name
        string sex;     // student's sex
        uint age;       // student's age
        string dept;    // student's dept
    }
    uint num;   // student数目
    uint[] ids; //id数组
    Student[] students;   // 存储所有student的数据
    address public headmaster;// 只有校长才能加学生
   
    
    event Insert(uint id);

    constructor() {
        num = 0;
        headmaster= msg.sender;
    }

       
    function insert(uint _id,string memory _name,string memory _sex,uint _age,string memory _dept) public {
        require(
            msg.sender == headmaster,
            "Only headmaster can add students."
        );
        students.push(Student(_id, _name, _sex, _age, _dept));   
        num += 1;
        ids.push(_id);
        uint id = _id;
        emit Insert(id);
    }

    function select_count() public view returns (uint n){
        return num;
    }

    function select_all_id() public view returns (uint[] memory) {
        require(num>0);
        return ids;
    }

    
    function select_id(uint _id) public view returns (Student memory){
        require(num>0);
        for (uint i = 0; i < num; ++i) {
            if (students[i].id == _id) {
                return students[i];
            }
        }
        return Student(0, "init", "init", 0, "init");
    }





}

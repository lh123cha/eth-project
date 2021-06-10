// SPDX-License-Identifier: GPL-3.0 
pragma solidity >=0.4.16 <0.9.0;

contract StudentContract{
    struct Student{
        int id;
        string name;
        string sex;
        int age;
        string dept;
    }

    uint number;

    event Insert(
       int id
    );
    address public adminstor;

    Student[] students;

    constructor(){
        adminstor=msg.sender;
        number=0;
    }

    function insert(int id,string memory name,string memory sex,int age,string memory dept) public{
        require(msg.sender==adminstor);
        students.push(Student(id,name,sex,age,dept));
        emit Insert(id);
        number+=1;
    }

    function select_count()public view returns (uint num) {
        return number;
    }

    function select_all_id() public view returns(int[20] memory){
        int[20] memory ids;
        for(uint i=0;i<number;i++){
            ids[i]=students[i].id;
        }
        return ids;
    }

    function select_id(int id) public view returns(Student memory){
        Student memory temp;
        for(uint i=0;i<number;i++){
            if(students[uint(i)].id==id){
                temp=students[uint(i)];
            }
        }
        return temp;
    }


}

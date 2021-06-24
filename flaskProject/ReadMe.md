# 以太坊readme

## 项目启动顺序：

1.先开启ganache
2.flask项目运行
3.开启前端Vue服务
4.访问localhost:8080的Vue前端

## flask路由函数说明（app.py）：

1.@app.route('/valid', methods=['POST'])
def valid()
判断当前用户是否合法

2.@app.route('/information', methods=['POST'])
def information()
添加用户信息，对应智能合约中insert_user

3.@app.route('/add', methods=['POST'])
def add()
添加订单信息，对应智能合约中insert_deal

4.@app.route('/output', methods=['POST'])
def output()
获取所有待完成订单，对应智能合约中select_all

5.@app.route('/finish_deal', methods=['POST'])
def finish_deal()
订单接收者接受订单，对应智能合约中finish_deal

6.@app.route('/cancel_deal', methods=['POST'])
def cancel_deal()
订单接收者取消订单，对应智能合约中cancel_deal

7.@app.route('/finish_deal_again', methods=['POST'])
def finish_deal_again()
订单发起者确认已完成订单，对应智能合约中finish_deal_again

8.@app.route('/myself', methods=['POST'])
def myself()
返回当前用户信息，对应智能合约中myself

9.@app.route('/myorder', methods=['POST'])
def myorder()
同output，但是返回前端时只保留接受发送者为自己且处于待完成状态的
10.@app.route('/myreceiveorder', methods=['POST'])
def myreceiveorder()
同output，但是返回前端时只保留接受者为自己的订单

定时函数APScheduler说明：
配置两个定时函数：更新订单剩余时间和用户奖励

```python
class Config(object):
    JOBS = [
        {
            'id': 'job1',
            'func': '__main__:update',
            'trigger': 'interval',
            'seconds': 20,
  		},
    	{
        'id': 'job2',
        'func': '__main__:bonus',
        'trigger': 'interval',
        'seconds': 60,

    	},
	]
```

def update() 定时调用智能合约中update_dealtime

def bonus() 定时调用智能合约中bonus_money

## 结构体功能说明（WaiMai.sol）：

1.用户结构体
    struct User {
        address addr;       // address
        uint id;        //id
        string name;    // name
        string tel;     // telephone
        uint money;       // money
        string dept;    // dept
    }
保存用户的信息

2.订单结构体
    struct Deal{
        uint id;       //id
        string text;   //text
        address send_user;  //senduser address
        uint money;     //money
        address receive_user;   //receiveuser address
        bool is_finish;  //whether deal is receive by receiver
        uint left_hour;  //hour time left
        bool is_timeout;  //whether time is out
        string tip; //tip of the deal
        bool is_finish_again; //whether deal is confirmed by sender
        string sender_tel;     // telephone(sender)
        string receiver_tel;     // telephone(receiver)
    }
保存订单的信息

## 智能合约函数功能说明（WaiMai.sol）：

1.function insert_user(string memory _name,string memory _tel,string memory _dept) public
插入新的用户
2.function insert_deal(string memory _text,uint _money,uint _time,string memory _tip) public
插入新的订单
3.function select_all() public view returns (Deal[] memory)
查询所有订单并返回数据
4.function finish_deal(uint _id) public
订单接收者接受订单
5.function cancel_deal(uint _id) public
订单接收者取消订单
6.function finish_deal_again(uint _id) public
订单发起者确认已完成订单
7.function select_myself() public view returns (User memory)
返回当前用户信息
8.function update_dealtime(uint _time_minus_hour) public
将所有的订单的剩余时间进行更新
9.function bonus_money(uint _money_bonus) public
给所有用户发放奖励



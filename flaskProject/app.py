import string

from flask import Flask, render_template, request
# from flask_login import LoginManager, login_user, logout_user, current_user, login_required
from web3_deploy import *
# from user import User
from flask_cors import CORS
import json
from flask_apscheduler import APScheduler

app = Flask(__name__)
CORS(app)
# app.secret_key = '1234567'

# login_manager = LoginManager()
# login_manager.login_view = 'login'
# login_manager.login_message_category = 'info'
# login_manager.login_message = '请登录'
# login_manager.init_app(app)

WaiMai_contract = compile_and_deploy("WaiMai")
eth.default_account = ""


# @login_manager.user_loader
# def load_user(user_id):
#     curr_user = User()
#     curr_user.id = user_id
#
#     return curr_user


@app.route('/')
def hello():
    # answer = helloworld_.functions.say().call()
    # return answer
    return render_template('index.html')


@app.route('/valid', methods=['POST'])
def valid():
    if request.method == 'POST':
        # try:
        data = request.get_data()
        json_data = json.loads(data.decode("utf-8"))
        address = json_data.get("Name")
        eth.default_account = address
        print(eth.default_account)
        test_num = WaiMai_contract.functions.test().call()
        print(test_num)
        # curr_user = User()
        # curr_user.id = address
        # # 通过Flask-Login的login_user方法登录用户
        # login_user(curr_user)
        # return render_template('information.html')
        res_list = {}
        data = json.loads(json.dumps(res_list))
        data['status'] = 1
        data['msg'] = "success"
        res = json.dumps(data, ensure_ascii=False)
        return res
    # except:
    #     return "<h1>以太坊账户登录失败</h1>"


@app.route('/information', methods=['POST'])
# @login_required
def information():
    if request.method == 'POST':
        # try:
        # eth.default_account = current_user.get_id()
        data = request.get_data()
        print("data = %s" % data)
        json_data = json.loads(data.decode("utf-8"))
        name = json_data.get("name")
        tel = json_data.get("tel")
        dept = json_data.get("dept")
        information = (name, tel, dept)
        tx_hash = WaiMai_contract.functions.insert_user(*information).transact()
        tx_receipt = eth.waitForTransactionReceipt(tx_hash)
        in_json = '{"statue": 1, "msg": "add success"}'
        return json.loads(in_json)
        # if student is None:
        #    return "<h1>找不到该id对应的value</h1>"
        # student = tuple(student)
        # return "<h1>信息完善成功</h1>"
    # except:
    #     return "<h1>你已完善过信息或者插入信息失败</h1>"


@app.route('/add', methods=['POST'])
# @login_required
def add():
    if request.method == 'POST':
        # try:
        # eth.default_account = current_user.get_id()
        data = request.get_data()
        print("data = %s" % data)
        json_data = json.loads(data.decode("utf-8"))
        name = json_data.get("name")
        money_string = json_data.get("money")
        time = json_data.get("num_time")
        tip = json_data.get("tip")
        money = int(money_string)
        deal = (name, money, int(time), tip)
        print(deal)
        tx_hash = WaiMai_contract.functions.insert_deal(*deal).transact()
        tx_receipt = eth.waitForTransactionReceipt(tx_hash)
        in_json = '{"statue": 1, "msg": "add record success"}'
        return json.loads(in_json)
    # except:
    #     return "<h1>发布信息失败</h1>"


@app.route('/output', methods=['POST'])
# @login_required
def output():
    if request.method == 'POST':
        # eth.default_account = current_user.get_id()
        deals = WaiMai_contract.functions.select_all().call()
        list1 = []
        for deal in deals:
            if (deal[5] == False):
                print("deal is :", deal)
                list2 = []
                for j in range(12):
                    list2.append(deal[j])
                list1.append(list2)
        print("list1 is: ", list1)
        jsonList = []
        for list in list1:
            aItem = {}
            print("list is ", list)
            aItem["id"] = list[0]
            aItem["username"] = list[2]
            aItem["money"] = list[3]
            aItem["mission"] = list[1]
            aItem["tip"] = list[8]
            aItem["time"] = list[6]
            jsonList.append(aItem)
        print(jsonList)
        jsonArr = json.dumps(jsonList, ensure_ascii=False)
        return jsonArr


# except:
#     return "<h1>查看信息失败</h1>"


@app.route('/add_deal')
# @login_required
def add_deal():
    return render_template('add.html')


@app.route('/finish_deal', methods=['POST'])
# @login_required
def finish_deal():
    if request.method == 'POST':
        # try:
        # eth.default_account = current_user.get_id()
        try:
            data = request.get_data()
            print("data = %s" % data)
            json_data = json.loads(data.decode("utf-8"))
            name = json_data.get("Name")
            deal_id = int(name)
            print(eth.default_account)
            print(deal_id)
            tx_hash = WaiMai_contract.functions.finish_deal(deal_id).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            in_json = '{"statue": 1, "msg": "success"}'
            return json.loads(in_json)
        except:
            in_json = '{"statue": 0, "msg": "failed"}'
            return json.loads(in_json)
    # except:
    #   return "<h1>接单失败</h1>"


# 618 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
# 取消订单和上面finish订单对应
@app.route('/cancel_deal', methods=['POST'])
# @login_required
def cancel_deal():
    if request.method == 'POST':
        # try:
        # eth.default_account = current_user.get_id()
        try:
            data = request.get_data()
            print("data = %s" % data)
            json_data = json.loads(data.decode("utf-8"))
            name = json_data.get("Name")
            deal_id = int(name)
            print(eth.default_account)
            print(deal_id)
            tx_hash = WaiMai_contract.functions.cancel_deal(deal_id).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            in_json = '{"statue": 1, "msg": "success"}'
            print("cancle success\n")
            return json.loads(in_json)
        except:
            in_json = '{"statue": 0, "msg": "failed"}'
            return json.loads(in_json)
    # except:
    #   return "<h1>接单失败</h1>"

# 621 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
@app.route('/finish_deal_again', methods=['POST'])
# @login_required
def finish_deal_again():
    if request.method == 'POST':
        # try:
        # eth.default_account = current_user.get_id()
        try:
            data = request.get_data()
            print("data = %s" % data)
            json_data = json.loads(data.decode("utf-8"))
            name = json_data.get("Name")
            deal_id = int(name)
            print(eth.default_account)
            print(deal_id)
            tx_hash = WaiMai_contract.functions.finish_deal_again(deal_id).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            in_json = '{"statue": 1, "msg": "success"}'
            return json.loads(in_json)
        except:
            in_json = '{"statue": 0, "msg": "failed"}'
            return json.loads(in_json)
    # except:
    #   return "<h1>接单失败</h1>"



# 618 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
# 返回当前账户个人信息
@app.route('/myself', methods=['POST'])
# @login_required
def myself():
    if request.method == 'POST':
        # eth.default_account = current_user.get_id()
        my = WaiMai_contract.functions.select_myself().call()
        print(my)
        jsonList = []
        aItem = {}
        aItem["addr"] = my[0]
        aItem["id"] = my[1]
        aItem["name"] = my[2]
        aItem["tel"] = my[3]
        aItem["money"] = my[4]
        aItem["dept"] = my[5]
        print(aItem)
        jsonArr = json.dumps(aItem, ensure_ascii=False)
        return jsonArr




# except:
#     return "<h1>查看信息失败</h1>"


# 618 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
# add myorder lh*******************************************************************************************
@app.route('/myorder', methods=['POST'])
def myorder():
    if request.method == 'POST':
        deals = WaiMai_contract.functions.select_all().call()
        list1 = []
        for deal in deals:
            if (deal[2] == eth.default_account and deal[5] == True):
                print("deal is :", deal)
                list2 = []
                for j in range(12):
                    list2.append(deal[j])
                list1.append(list2)
        jsonList = []
        for list in list1:
            aItem = {}
            print("list is ", list)
            aItem["id"] = list[0]
            aItem["username"] = list[2]
            aItem["money"] = list[3]
            aItem["mission"] = list[1]
            aItem["tip"] = list[8]
            aItem["time"] = list[6]
            jsonList.append(aItem)
        print("myorder is: ", jsonList)
        jsonArr = json.dumps(jsonList, ensure_ascii=False)
        return jsonArr


# 621 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
@app.route('/myreceiveorder', methods=['POST'])
def myreceiveorder():
    if request.method == 'POST':
        deals = WaiMai_contract.functions.select_all().call()
        list1 = []
        for deal in deals:
            if (deal[4] == eth.default_account):
                print("deal is :", deal)
                list2 = []
                for j in range(12):
                    list2.append(deal[j])
                list1.append(list2)
        jsonList = []
        for list in list1:
            aItem = {}
            print("list is ", list)
            aItem["id"] = list[0]
            aItem["username"] = list[2]
            aItem["money"] = list[3]
            aItem["mission"] = list[1]
            aItem["tip"] = list[8]
            aItem["time"] = list[6]
            jsonList.append(aItem)
        print("myorder is: ", jsonList)
        jsonArr = json.dumps(jsonList, ensure_ascii=False)
        return jsonArr


#
# 配置自动任务的类
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

# 618 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
def update():
    # try:
    now_account = eth.default_account
    eth.default_account = eth.accounts[0]
    print("starting update!")
    tx_hash = WaiMai_contract.functions.update_dealtime(1).transact()
    tx_receipt = eth.waitForTransactionReceipt(tx_hash)
    print("update success!")
    eth.default_account = now_account;


# except:
#     print("update failed!")

# 621 edited by qk!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
def bonus():
    # try:
    now_account = eth.default_account
    eth.default_account = eth.accounts[0]
    print("starting bonus!")
    tx_hash = WaiMai_contract.functions.bonus_money(10).transact()
    tx_receipt = eth.waitForTransactionReceipt(tx_hash)
    print("bonus success!")
    eth.default_account = now_account;


# except:
#     print("update failed!")


# unused
@app.route('/query/<int:key>')
# @login_required
def query_process(key):
    key_int = int(key)
    # eth.default_account = current_user.get_id()

    try:
        student = WaiMai_contract.functions.select_id(key_int).call()
        if student is None:
            return "<h1>找不到该id对应的value</h1>"
        student = tuple(student)
        return render_template('answer.html', Value=student)
    except:
        return "<h1>query error</h1>"


# @app.route('/logout')
# # @login_required
# def logout():
#     logout_user()
#     return 'Logged out successfully!'

if __name__ == '__main__':
    app.config.from_object(Config())
    scheduler = APScheduler()
    scheduler.init_app(app)
    scheduler.start()
    app.run()

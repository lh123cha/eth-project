import string
import json
from flask import Flask, render_template, request
from flask_login import LoginManager, login_user, logout_user, current_user, login_required
from web3_deploy import *
from flask_cors import CORS

app = Flask(__name__)

CORS(app)
WaiMai_contract = compile_and_deploy("WaiMai")
eth.default_account = ""

@app.route('/')
def hello():
    # answer = helloworld_.functions.say().call()
    # return answer
    return render_template('index.html')

@app.route('/valid', methods=['POST'])
def valid():
    if request.method == 'POST':
        try:
            data = request.get_data()
            print("data = %s" % data)
            json_data = json.loads(data.decode("utf-8"))
            eth.default_account = json_data.get("Name")
            print(eth.default_account)
            test_num = WaiMai_contract.functions.test().call()
            print(test_num)
            #return render_template('information.html')
            res_list = {}
            data = json.loads(json.dumps(res_list))
            data['status'] = 1
            data['msg'] = "success"
            res = json.dumps(data, ensure_ascii=False)
            return res
        except:
            return "<h1>以太坊账户登录失败</h1>"


@app.route('/information', methods=['POST'])
def information():
    if request.method == 'POST':
        try:
            data = request.get_data()
            print("data = %s" % data)
            json_data = json.loads(data.decode("utf-8"))
            name = json_data.get("Name")
            tel = json_data.get("Tel")
            dept = json_data.get("Dept")
            information = (name,tel,dept)
            tx_hash = WaiMai_contract.functions.insert_user(*information).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            in_json='{"statue": 1, "msg": "add success"}'
            return json.loads(in_json)
        except:
            return "<h1>你已完善过信息或者插入信息失败</h1>"



@app.route('/add', methods=['POST'])
def add():
    if request.method == 'POST':
        # try:
            data = request.get_data()
            print("data = %s" % data)
            json_data = json.loads(data.decode("utf-8"))
            name = json_data.get("name")
            money_string = json_data.get("money")
            time=json_data.get("num_time")
            tip=json_data.get("tip")
            money = int(money_string)
            deal = (name,money,int(time),tip)
            print(deal)
            tx_hash = WaiMai_contract.functions.insert_deal(*deal).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            in_json = '{"statue": 1, "msg": "add record success"}'
            return json.loads(in_json)
        # except:
        #     return "<h1>发布信息失败</h1>"


@app.route('/output',methods=['POST'])
def output():
    # try:
    if request.method == 'POST':
        deals = WaiMai_contract.functions.select_all().call()
        print(deals)
        list1 = []
        for deal in deals:
            print(deal)
            print(len(deal))
            list2 = []
            for j in range():
                list2.append(deal[j])
            list1.append(list2)
            print(list1)
        jsonList=[]
        aItem = {}
        for list in list1:
            aItem["id"]=list[0]
            aItem["username"]=list[1]
            aItem["money"]=list[2]
            aItem["mission"]=list[3]
            aItem["tip"]=list[4]
            aItem["time"]=list[5]
            jsonList.append(aItem)
        jsonArr = json.dumps(jsonList, ensure_ascii=False)
        return jsonArr

        # return render_template('output.html', list3=list1,now_user=eth.default_account)
    # except:
    #     return "<h1>查看信息失败</h1>"



@app.route('/add_deal')
def add_deal():
    return render_template('add.html')


@app.route('/finish_deal', methods=['POST'])
def finish_deal():
    if request.method == 'POST':
        # try:
            name = request.form.get("Name")
            deal_id = int(name)
            tx_hash = WaiMai_contract.functions.finish_deal(deal_id).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            return "<h1>接单成功</h1>"
        # except:
        #   return "<h1>接单失败</h1>"


@app.route('/query/<int:key>')
def query_process(key):
    key_int = int(key)

    try:
        student = WaiMai_contract.functions.select_id(key_int).call()
        if student is None:
            return "<h1>找不到该id对应的value</h1>"
        student = tuple(student)
        return render_template('answer.html', Value=student)
    except:
        return "<h1>query error</h1>"


if __name__ == '__main__':
    app.run()

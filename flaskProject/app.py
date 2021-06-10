import json
import string

from flask import Flask, render_template, request
# from flask_login import LoginManager, login_user, logout_user, current_user, login_required
from web3_deploy import *
from flask_cors import CORS


app = Flask(__name__)
CORS(app)
student_contract = compile_and_deploy("WaiMai")


@app.route('/')
def hello():
    # answer = helloworld_.functions.say().call()
    # return answer
    return render_template('index.html')


@app.route('/query', methods=['POST'])
def query():
    if request.method == 'POST':
        id_name = request.form.get("Name")
        return query_process(id_name)


@app.route('/add',methods=['POST'])
def add():
    try:
        if request.method=='POST':
            data=request.get_data()
            print("data = %s"%data)
            json_data=json.loads(data.decode("utf-8"))
            id = json_data.get("id")
            name=json_data.get("name")
            sex=json_data.get("sex")
            age=json_data.get("age")
            dept=json_data.get("dept")
            mystudent=(int(id),name,sex,int(age),dept)
            print(mystudent)
            tx_hash = student_contract.functions.insert(*mystudent).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            return "<h1>insert success!</h1>"
        # student = (1, "qk", "M", 22, "ECNU")
        # student2 = (2, "lh", "M", 22, "ECNU")
        # student3 = (3, "czr", "M", 22, "ECNU")
        # student4 = (4, "yzk", "M", 22, "ECNU")
        # # 构造函数通过rpc向以太坊发送该函数交易
        # # *为解包操作，即传入的参数为: id, name, sex, age, dept
        # tx_hash = student_contract.functions.insert(*student).transact()
        # tx_receipt = eth.waitForTransactionReceipt(tx_hash)
        # tx_hash2 = student_contract.functions.insert(*student2).transact()
        # tx_receipt2 = eth.waitForTransactionReceipt(tx_hash2)
        # tx_hash3 = student_contract.functions.insert(*student3).transact()
        # tx_receipt3 = eth.waitForTransactionReceipt(tx_hash3)
        # tx_hash4 = student_contract.functions.insert(*student4).transact()
        # tx_receipt = eth.waitForTransactionReceipt(tx_hash4)
        # return "<h1>insert success!</h1>"
    except:
        return "<h1>insert error!</h1>"


def query_process(key):
    # all_id = []
    # try:
    #     all_id = student_contract.functions.select_all_id().call()
    # except:
    #     return "<h1>all query error </h1>"
    key_int = int(key)

    try:
        student = student_contract.functions.select_id(key_int).call()
        if student is None:
            return "<h1>找不到该id对应的value</h1>"
        student = tuple(student)
        return render_template('answer.html', Value=student)
    except:
        return "<h1>query error</h1>"


if __name__ == '__main__':
    app.run()

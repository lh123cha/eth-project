import json
import string

from flask import Flask, render_template, request
# from flask_login import LoginManager, login_user, logout_user, current_user, login_required
from web3_deploy import *
from flask_cors import CORS


app = Flask(__name__)
CORS(app)
waimai_contract = compile_and_deploy("WaiMai")

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
            name = json_data.get("name")
            tel=json_data.get("tel")
            dept=json_data.get("dept")
            mywaimai=(name,tel,dept)
            print(mywaimai)
            eth.default_account = eth.accounts[1]
            tx_hash = waimai_contract.functions.insert_user(*mywaimai).transact()
            tx_receipt = eth.waitForTransactionReceipt(tx_hash)
            return "<h1>insert success!</h1>"
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
        student = waimai_contract.functions.select_id(key_int).call()
        if student is None:
            return "<h1>找不到该id对应的value</h1>"
        student = tuple(student)
        return render_template('answer.html', Value=student)
    except:
        return "<h1>query error</h1>"


if __name__ == '__main__':
    app.run()

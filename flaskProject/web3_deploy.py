from web3 import Web3
import os

# ganache 默认rpc url和端口
url = 'http://localhost:7545'
w3 = Web3(Web3.HTTPProvider(url))
if not w3.isConnected():
    raise Exception("Ethereum connection failed!")
eth = w3.eth
# 设置发送交易的账户
eth.default_account = eth.accounts[0]
print("Ethereum({}) connection successful!".format(url))


def read_file(file):
    """
        读取文件中的全部内容
    """
    with open(file, 'r') as f:
        return f.read()


def deploy(contract_name, *args):
    """
        部署合约
    """
    # 从文件中读取abi和bytecode
    # 默认路径为 contract_name/contract_name.abi 和 contract_name/contract_name.bin
    abi_file = contract_name + r'/' + contract_name + '.abi';
    bytecode_file = contract_name + r'/' + contract_name + '.bin';
    abi = read_file(abi_file)
    bytecode = read_file(bytecode_file)
    # 生成合约，调用其构造函数将其部署至以太坊上
    contract = eth.contract(abi=abi, bytecode=bytecode)
    tx_hash = contract.constructor(*args).transact()
    # 获取交易回执，使用其中的合约地址定位已部署的合约
    tx_receipt = eth.waitForTransactionReceipt(tx_hash)
    deployed_contract = eth.contract(address=tx_receipt.contractAddress, abi=abi)
    print("Contract '{}' was successfully deployed! Contract address: '{}'".
            format(contract_name, tx_receipt.contractAddress))
    return deployed_contract


def compile(contract_name):
    """
        编译合约
    """
    cmd = 'solc --abi --bin --overwrite -o ' + contract_name + ' ' + contract_name + '.sol'
    os.system(cmd)


def compile_and_deploy(contract_name, *args):
    """
        编译合约并部署
    """
    compile(contract_name)
    return deploy(contract_name, *args)
    


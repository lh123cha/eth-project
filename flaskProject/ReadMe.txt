顺序
1.先开启ganache
2.flask项目运行
3.首页输入以太坊账户号（可以从ganache复制）
4.如果是第一次使用需要提交自己信息否则无法发布任务（对应智能合约中的insert user）
5.超链接到发布任务页面发布任务（对应智能合约中的insert deal）
6.超链接到查看任务查看当前任务（对应智能合约中select all），在此页面还可以接收任务（对应智能合约中finish deal）
7.查询某个账户信息，这块暂时用url get 访问，比如localhost:5000/query/0表示查询0号账户信息



注意：当前用flask的template模板作为前端，修改成Vue前后端分离需要修改每个函数的数据交互

已知bug:
由于是登录时修改eth当前账户，两个人同时登陆会使得eth变成后一个，这个问题现在先不管也许flask login可以解决
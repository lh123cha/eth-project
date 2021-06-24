beego项目运行，需要实现安装go语言环境
1.设置代理
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn,direct

2.下载
go get github.com/beego/bee
go get github.com/astaxie/beego

3.进入myproject文件夹，执行
go run main.go

4.打开浏览器在localhost:5000可看到项目界面
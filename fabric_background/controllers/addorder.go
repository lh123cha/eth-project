package controllers

import (

	//"encoding/json"
	//"fmt"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"

	//"fmt"
	"myproject/models"


)

var orderId int =1

type AddorderController struct {
	beego.Controller
}

type Order struct {

	Name string

	Money string

	Tip string

	Num_time json.Number

	MethodName string

	Msg string

}

func (this *AddorderController) Addorder ()  {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", this.Ctx.Request.Header.Get("Origin"))
	//var user User
	//
	//data := this.Ctx.Input.RequestBody
	//
	////json数据封装到user对象中
	//
	//err := json.Unmarshal(data, &user)
	//
	//if err != nil {
	//
	//	fmt.Println("json.Unmarshal is err:", err.Error())
	//
	//}
	//
	//fmt.Println(user)
	//
	//this.Ctx.WriteString(user.name)

	//var ob User
	//var err error
	//if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
	//	this.Data["json"] = "{\"Name\":\"" + ob.name + "\"}"
	//} else {
	//	this.Data["json"] = err.Error()
	//}
	//this.ServeJSON()


	//get方法获取参数
	//id := this.GetString("id")

	//name := this.GetString("Name")
	fmt.Println("hello")

	//fmt.Println(name)
	//this.Ctx.WriteString(name)

	var order Order
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	fmt.Println(data)
	err := json.Unmarshal(data, &order)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(order)
	//this.Ctx.WriteString(user.Name+user.Tel)

	if order.Name != "" && order.Money!= "" && order.Tip!= "" && order.Num_time!= ""{

		var old string=models.Query("transaction")
		fmt.Println(old)
		var tran string
		var time=order.Num_time.String()
		if old=="" {
			tran = strconv.Itoa(orderId) + " " + "lh" + " " + order.Name + " " + order.Money + " " + order.Tip + " " + time
		} else{
			tran = old + "-" + strconv.Itoa(orderId) + " " + "lh" + " " + order.Name + " " + order.Money + " " + order.Tip + " " + time
		}
		fmt.Println(tran)
		models.Insert("transaction",tran)


	}

}
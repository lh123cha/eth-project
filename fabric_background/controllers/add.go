package controllers

import (

	//"encoding/json"
	//"fmt"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"fmt"
	"myproject/models"


)

type AddController struct {
	beego.Controller
}

type User struct {

	Name string

	Tel string

	Dept string

	MethodName string

	Msg string

}

func (this *AddController) Add ()  {
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

	var user User
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	fmt.Println(data)
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(user)
	this.Ctx.WriteString(user.Name+user.Tel)

	if user.Name != "" && user.Tel != ""{
		models.Insert(user.Name,user.Tel)
	}


	


}

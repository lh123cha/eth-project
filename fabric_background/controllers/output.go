package controllers

import (

	//"encoding/json"
	//"fmt"

	"fmt"
	//"container/list"
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	//"os"
	"strings"

	//"fmt"
	"myproject/models"


)

type OutController struct {
	beego.Controller
}


func (this *OutController) Out ()  {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", this.Ctx.Request.Header.Get("Origin"))

	var tran string=models.Query("transaction")
	type entrygroup struct{
		OrderId string `json:"id"`
		Owner string `json:"username"`
		Money string `json:"money"`
		Name string `json:"mission"`
		Tip string `json:"tip"`
		Time string `json:"time"`

	}

	//str:="name1 money1 tip1 time1;name2 money2 tip2 time2;"//甩头ring
	split1:=strings.Split(tran,"-")//甩头ring[]

	//fmt.Printf("%q\n",split1)
	//l := list.New() //创建一个新的list
	//for i := 0; i < 5; i++ {
	//
	//}
	var entrygroups  []entrygroup
	for _,entry:=range split1{
		split2:=strings.Split(entry," ")
		if len(split2)==1{
			continue
		}

		jsongroup:=entrygroup{
			split2[0],
			split2[1],
			split2[3],
			split2[2],
			split2[4],
			split2[5]}
		entrygroups=append(entrygroups,jsongroup)
		//b,err:=json.Marshal(jsongroup)
		//fmt.Printf("%s\n",b)
		//if err!=nil{
		//	fmt.Println(err)
		//}
		//os.Stdout.Write([]byte("1234"))
		//os.Stdout.Write(entrygroups)
		//l.PushBack(b)
		//this.Data["json"] = &jsongroup
		//this.ServeJSON()
	}
	fmt.Println(entrygroups)
	//type JsonResponse struct {
	//	Data []entrygroup `json:"data"`
	//}
	//
	//var jsonResponse JsonResponse
	//jsonResponse.Data = entrygroups
	this.Data["json"] = entrygroups
	//fmt.Println(&jsonResponse)
	this.ServeJSON()



}


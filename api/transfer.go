package api

import (
	"easypay/model"
	"easypay/service"
	"easypay/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func transfer(ctx *gin.Context) {

	name := ctx.PostForm("name")
	moneyNum := ctx.PostForm("num")
	money, _ := strconv.Atoi(moneyNum)

	if money < 0 {
		tool.RespErrorWithData(ctx, "不能转负数")
		return
	}

	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	flag, err := service.IsMoneyEnough(money, username)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println(err)
	}

	if !flag {
		tool.RespErrorWithData(ctx, "余额不足")
		return
	}
	//对象

	txt := username + "转了你" + moneyNum + "元"
	user := model.User{
		Username: name,
		Money:    money,
	}
	err = service.IncreaseMoney(user)
	OuserId, err := service.SelectIdByUsername(name)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	record := model.Record{
		Pid:  OuserId,
		Txt:  txt,
		Time: time.Now(),
	}
	//记录充值
	err = service.AddRecord(record)
	if err != nil {
		fmt.Println("add record err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	//自己
	myTxt := "你向" + name + "转了" + moneyNum + "元"
	Myuser := model.User{
		Username: username,
		Money:    -money,
	}
	err = service.IncreaseMoney(Myuser)
	userId, err := service.SelectIdByUsername(username)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	MyRecord := model.Record{
		Pid:  userId,
		Txt:  myTxt,
		Time: time.Now(),
	}

	//记录充值

	err = service.AddRecord(MyRecord)
	if err != nil {
		fmt.Println("add record err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "转账成功")
}

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

func topUp(ctx *gin.Context) {
	moneyNum := ctx.PostForm("money")
	money, err := strconv.ParseFloat(moneyNum, 32)
	if err != nil {
		tool.RespErrorWithData(ctx, "stou")
		fmt.Println("err is", err)
		return
	}

	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithData(ctx, "token有误")
		fmt.Println(err)
	}

	txt := username + "充值了" + moneyNum + "元"
	user := model.User{
		Username: username,
		Money:    money,
	}
	err = service.IncreaseMoney(user)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("err is", err)
		return
	}

	userId, err := service.SelectIdByUsername(username)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("err is", err)
		return
	}
	record := model.Record{
		Pid:  userId,
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

	tool.RespSuccessfulWithData(ctx, "你成功充值"+moneyNum)
}

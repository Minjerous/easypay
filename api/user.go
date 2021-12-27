package api

import (
	"easypay/dao"
	"easypay/model"
	"easypay/service"
	"easypay/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登入
func help(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"help接口(GET)": "http://106.55.225.88:8020/help",
		"注册接口(POST)":  "http://106.55.225.88:8020/register               (POSTFROM)有 password  username  ",
		"登入接口(POST)":  "http://106.55.225.88:8020/login                  (POSTFROM)有 password  username ",
		"改密接口(POST)":  "http://106.55.225.88:8020/user/changepassword ",
		"充值接口(POST)":  "http://106.55.225.88:8020/user/topup             (POSTFROM) key : num",
		"转账接口(POST)":  "http://106.55.225.88:8020/user/transfer          (POSTFROM) key : num；key : name",
		"备注接口(GET)h":  "http://106.55.225.88:8020/user/getrecord",
	})
}
func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}

	ctx.SetCookie("username", username, 60, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

//注册
func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if len(username) < 3 {
		tool.RespErrorWithData(ctx, "用户名至少两位")
		return
	}

	if len(password) < 8 {
		tool.RespErrorWithData(ctx, "密码必须大于8位")
		return
	}

	//加盐加密
	passWord := tool.HashWithSalted(password)
	user := model.User{
		Username: username,
		Password: passWord,
	}

	//判断是否用户名已经被注册
	flag, err := service.IsRepeatUsername(user.Username)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithData(ctx, "用户名重复")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	} else {
		tool.RespSuccessful(ctx)
	}
}
func getRecord(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	Id, err := dao.SelectIdByUsername(username)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Print(err)
		return
	}
	record, err := service.GetRecord(Id)
	if err != nil {
		fmt.Println("get record err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespAllRecord(ctx, record)
}
func changePassword(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	FirstNewPassword := ctx.PostForm("newpasswordOne")
	SecondNewPassword := ctx.PostForm("newpasswordTwo")

	user := model.User{
		Username: username,
		Password: tool.HashWithSalted(FirstNewPassword),
	}

	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println(err)
		tool.RespInternalError(ctx)
		return
	}

	if flag == false {
		tool.RespErrorWithData(ctx, "用户不存在")
		return
	}
	if flag {
		flag, err := service.IsPasswordCorrect(username, password)
		if err != nil {
			fmt.Println(err)
			tool.RespInternalError(ctx)
			return
		}

		if flag {
			if FirstNewPassword == SecondNewPassword {
				err := service.Password(user)
				if err == nil {
					tool.RespSuccessfulWithData(ctx, "修改成功")
					return
				} else {
					tool.RespErrorWithData(ctx, "修改失败")
					fmt.Println("err by change password is", err)
					return
				}
				return
			} else {
				tool.RespErrorWithData(ctx, "两次密码输入错误")
				return
			}
		}
		return
	}
}

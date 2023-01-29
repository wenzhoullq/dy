package controller

import (
	"github.com/gin-gonic/gin"
	"simple-demo/common"
	"simple-demo/response"
	"simple-demo/service"
)

const (
	Check_AccountParam_Fail string = "账号密码格式错误"
	UserRegister_Fail       string = "账号注册失败,请稍后再试"
	UserRegister_Success    string = "账号注册成功,即将跳转新的页面"
	UserInfo_Success        string = "获取用户成功"
	UserInfo_Fail           string = "获取用户失败"
	UserLogin_Success       string = "账号登陆成功,即将跳转新的页面"
	UserLogin_Fail          string = "账号不存在或者密码错误"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ok := common.Check_AccountParam(username, password)
	if !ok {
		response.Fail(c, Check_AccountParam_Fail, nil)
		return
	}
	mes, err := service.UserRegister(username, password)
	if err != nil {
		response.Fail(c, UserRegister_Fail, mes)
		return
	}
	response.Success(c, UserRegister_Success, mes)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ok := common.Check_AccountParam(username, password)
	if !ok {
		response.Fail(c, Check_AccountParam_Fail, nil)
		return
	}
	mes, err := service.Login(username, password)
	if err != nil {
		response.Fail(c, UserLogin_Fail, mes)
		return
	}
	response.Success(c, UserLogin_Success, mes)

}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	if mess, err := service.UserInfo(uid, token); err == nil {
		response.Success(c, UserInfo_Success, mess)
	} else {
		response.Fail(c, UserRegister_Fail, mess)
		return
	}
}

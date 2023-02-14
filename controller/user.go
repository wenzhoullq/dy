package controller

import (
	"github.com/gin-gonic/gin"
	"simple-demo/common"
	"simple-demo/response"
	"simple-demo/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

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
	username := c.PostForm("username")
	password := c.PostForm("password")
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
	err := common.CheckToken(token)
	if err != nil {
		response.Fail(c, CheckToken_Fail, nil)
		return
	}
	if mess, err := service.UserInfo(uid); err == nil {
		response.Success(c, UserInfo_Success, mess)
	} else {
		response.Fail(c, UserRegister_Fail, mess)
		return
	}
}

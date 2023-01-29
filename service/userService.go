package service

import (
	"simple-demo/common"
	"simple-demo/log"
	mes "simple-demo/proto/pkg"
	"simple-demo/repository"
	"strconv"
)

var (
	InitStatusCode      int32  = 0
	InitStatusMsg       string = "success"
	InitUserId          int64  = 0
	InitToken           string = ""
	Fail_StatusCode     int32  = -1
	UserNameExit_Msg    string = "该用户名已存在,请修改用户名后再注册"
	UserNameNotExit_Msg string = "该用户名不存在,请检查账号密码"
	User_Login_Failed   string = "账号密码错误,登陆失败"
	User_Login_Success  string = "登陆成功，即将为你跳转新的页面"
	Fail_Resigter       string = "注册失败,请重试"
	Fail_InsterUser     string = "插入用户失败"
	Fail_GetUid         string = "获得UID失败"
	Fail_GenerateToken  string = "创建Token失败"
	Fail_GenerateCache  string = "插入缓存失败"
	Fail_GetUserInfo    string = "获取用户信息失败"
	Key_UserInfo        string = "userInfo_uid_"
)

func UserRegister(username string, password string) (*mes.DouyinUserRegisterResponse, error) {

	mes := &mes.DouyinUserRegisterResponse{
		StatusCode: &InitStatusCode,
		StatusMsg:  &User_Login_Success,
		UserId:     &InitUserId,
		Token:      &InitToken,
	}
	var err error = nil
	//检查username是否存在
	ok := repository.UsernameExit(username)
	if ok {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &UserNameExit_Msg
		log.Info(Fail_Resigter, username)
		//err = errors.New(Fail_Resigter)
		return mes, err
	}
	//插入User
	err = repository.InsertUser(username, password)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//获得UID
	userId, err := repository.GetUserId(username)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//获得User
	user, err := repository.GetUser(userId)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//生成token
	token, err := common.GenToken(userId, username)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//user插入redis
	processUidKey := common.ProcessUid(Key_UserInfo, strconv.Itoa(int(userId)))
	err = common.InsertStringCache(processUidKey, user)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	mes.Token = &token
	mes.UserId = &userId
	log.Info(mes)
	return mes, nil
}

// 先判断redis是否有token,没有则返回
func UserInfo(uid string, token string) (*mes.DouyinUserResponse, error) {
	mess := &mes.DouyinUserResponse{
		StatusCode: &InitStatusCode,
		StatusMsg:  &InitStatusMsg,
		User:       nil,
	}
	err := common.CheckToken(uid, token)
	if err != nil {
		mess.StatusCode = &Fail_StatusCode
		mess.StatusMsg = &Fail_GetUserInfo
		log.Error(mess, err)
		return mess, err
	}
	processUidKey := common.ProcessUid(Key_UserInfo, uid)
	user, err := repository.GetCacheUser(processUidKey)
	if err != nil {
		mess.StatusCode = &Fail_StatusCode
		mess.StatusMsg = &Fail_GetUserInfo
		log.Error(mess, err)
		return mess, err
	}
	mess.User = user
	log.Info(mess)
	return mess, nil
}

// 账号登陆
func Login(username string, password string) (*mes.DouyinUserLoginResponse, error) {
	mes := &mes.DouyinUserLoginResponse{
		StatusCode: &InitStatusCode,
		StatusMsg:  &User_Login_Success,
	}
	var err error = nil
	ok := repository.UsernameExit(username)
	if !ok {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &UserNameNotExit_Msg
		log.Error(mes, err)
		return mes, err
	}
	//检查用户密码
	err = repository.CheckAccount(username, password)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &User_Login_Failed
		log.Error(mes, err)
		return mes, err
	}
	//获得UID
	userId, err := repository.GetUserId(username)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//获得User
	user, err := repository.GetUser(userId)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//生成token
	token, err := common.GenToken(userId, username)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	//处理uid
	processUidKey := common.ProcessUid(Key_UserInfo, strconv.Itoa(int(userId)))
	//插入redis
	err = common.InsertStringCache(processUidKey, user)
	if err != nil {
		mes.StatusCode = &Fail_StatusCode
		mes.StatusMsg = &Fail_Resigter
		log.Error(mes, err)
		return mes, err
	}
	mes.Token = &token
	mes.UserId = &userId
	log.Info(mes)
	return mes, nil
}

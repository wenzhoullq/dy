package service

import (
	"simple-demo/common"
	"simple-demo/log"
	mes "simple-demo/proto/pkg"
	"simple-demo/repository"
	"strconv"
)

func UserRegister(username string, password string) (*mes.DouyinUserRegisterResponse, error) {

	mes := &mes.DouyinUserRegisterResponse{
		StatusCode: InitStatusCode,
		StatusMsg:  UserLoginSuccess,
		UserId:     InitUserId,
		Token:      InitToken,
	}
	var err error = nil
	//检查username是否存在
	ok := repository.UsernameExit(username)
	if ok {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = UserNameExitMsg
		log.Info(FailRegisterUser, username)
		//err = errors.New(Fail_Resigter)
		return mes, err
	}
	//插入User
	err = repository.InsertUser(username, password)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//获得UID
	userId, err := repository.GetUserId(username)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//获得User
	user, err := repository.GetUser(userId)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//生成token
	token, err := common.GenToken(userId, username)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//user插入redis
	processUidKey := common.ProcessUid(KeyUserinfo, strconv.Itoa(int(userId)))
	err = common.InsertStringCache(processUidKey, user)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	mes.Token = token
	mes.UserId = userId
	log.Info(mes)
	return mes, nil
}

// 先判断redis是否有token,没有则返回
func UserInfo(uid string) (*mes.DouyinUserResponse, error) {
	mess := &mes.DouyinUserResponse{
		StatusCode: InitStatusCode,
		StatusMsg:  InitStatusMsg,
		User:       nil,
	}

	processUidKey := common.ProcessUid(KeyUserinfo, uid)
	user, err := repository.GetCacheUser(processUidKey)
	if err != nil {
		mess.StatusCode = FailStatusCode
		mess.StatusMsg = FailGetUserinfo
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
		StatusCode: InitStatusCode,
		StatusMsg:  UserLoginSuccess,
	}
	var err error = nil
	ok := repository.UsernameExit(username)
	if !ok {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = UserNameNotExitMsg
		log.Error(mes, err)
		return mes, err
	}
	//检查用户密码
	err = repository.CheckAccount(username, password)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = UserLoginFailed
		log.Error(mes, err)
		return mes, err
	}
	//获得UID
	userId, err := repository.GetUserId(username)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//获得User
	user, err := repository.GetUser(userId)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//生成token
	token, err := common.GenToken(userId, username)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	//处理uid
	processUidKey := common.ProcessUid(KeyUserinfo, strconv.Itoa(int(userId)))
	//插入redis
	err = common.InsertStringCache(processUidKey, user)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailRegisterUser
		log.Error(mes, err)
		return mes, err
	}
	mes.Token = token
	mes.UserId = userId
	log.Info(mes)
	return mes, nil
}

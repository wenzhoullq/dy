package service

import (
	"fmt"
	//"simple-demo/common"
	mes "simple-demo/proto/pkg"
	"simple-demo/repository"
)

func Register(username string, password string) (*mes.DouyinUserRegisterRes, error) {
	err := repository.UsernameEixt(username)
	if err != nil {
		mes := &mes.DouyinUserRegisterRes{
			StatusCode: -1,
			StatusMsg:  "failed",
		}
		return mes, fmt.Errorf("register is failed:%v", err)
	}
	token := ""
	//token, err := common.GenToken(username, password)
	if err != nil {
		mes := &mes.DouyinUserRegisterRes{
			StatusCode: -1,
			StatusMsg:  "failed",
		}
		return mes, fmt.Errorf("register is failed:%v", err)
	}
	userId, err := repository.InsertAccount(username, password)
	if err != nil {
		mes := &mes.DouyinUserRegisterRes{
			StatusCode: -1,
			StatusMsg:  "failed",
		}
		return mes, fmt.Errorf("register is failed:%v", err)
	}
	mes := &mes.DouyinUserRegisterRes{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     userId,
		Token:      token,
	}
	return mes, nil
}

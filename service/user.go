package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	message "github.com/RaymondCode/simple-demo/proto/pkg"
	"github.com/RaymondCode/simple-demo/repository"
)

func Register(username string, password string) (*message.DouyinUserRegisterRes, error) {
	err := repository.UsernameEixt(username)
	if err != nil {
		mes := &message.DouyinUserRegisterRes{
			StatusCode: -1,
			StatusMsg:  "failed",
		}
		return mes, fmt.Errorf("register is failed:%v", err)
	}
	token, err := common.GenToken(username, password)
	if err != nil {
		mes := &message.DouyinUserRegisterRes{
			StatusCode: -1,
			StatusMsg:  "failed",
		}
		return mes, fmt.Errorf("register is failed:%v", err)
	}
	userId, err := repository.InsertAccount(username, password)
	if err != nil {
		mes := &message.DouyinUserRegisterRes{
			StatusCode: -1,
			StatusMsg:  "failed",
		}
		return mes, fmt.Errorf("register is failed:%v", err)
	}
	mes := &message.DouyinUserRegisterRes{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     userId,
		Token:      token,
	}
	return mes, nil
}

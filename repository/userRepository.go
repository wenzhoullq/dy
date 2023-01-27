package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"simple-demo/common"
)

const (
	USERNAMEEXIT = "1"
)

type User struct {
	// gorm.Model
	Id       int64  `gorm:"column:user_id; primary_key;"`
	Name     string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
}

func UsernameExit(userName string) error {
	db := common.GetDB()
	user := User{}
	err := db.Where("user_name = ?", userName).Find(&user).Error
	if err == nil {
		return errors.New(USERNAMEEXIT)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
func InsertAccount(username string, password string) (int64, error) {
	var userId int64 = 0
	return userId, nil
}

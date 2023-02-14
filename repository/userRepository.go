package repository

import (
	"encoding/json"
	"simple-demo/common"
	"simple-demo/log"
	message "simple-demo/proto/pkg"
)

type User struct {
	Id             int64  `gorm:"column:userid"`
	Name           string `gorm:"column:username"`
	Password       string `gorm:"column:password"`
	Follow_count   int64  `gorm:"column:follow_count"`
	Follower_count int64  `gorm:"column:follower_count"`
	Is_follow      bool   `gorm:"column:is_follow"`
}

func UsernameExit(userName string) bool {
	db := common.GetDB()
	user := User{}
	err := db.Table("users").Where("username = ?", userName).Find(&user).Error
	if err != nil {
		return false
	}
	return true
}
func InsertUser(username string, password string) error {
	db := common.GetDB()
	user := User{
		0,
		username,
		password,
		0,
		0,
		false,
	}
	result := db.Table("users").Create(&user)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
func GetUser(userId int64) (User, error) {
	db := common.GetDB()
	user := User{}
	result := db.Table("users").Where("userid = ?", userId).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
func GetUserId(username string) (int64, error) {
	db := common.GetDB()
	user := User{}
	result := db.Table("users").Where("username = ?", username).Find(&user)
	if result.Error != nil {
		log.Error(result.Error)
		return 0, result.Error
	}
	log.Infof("获得userId: ", user.Id)
	return user.Id, nil
}
func GetCacheUser(token string) (*message.User, error) {
	//data是{}interface格式,需要做格式转换
	data, err := common.GetStringCache(token)
	if err != nil {
		return nil, err
	}
	user := &message.User{}
	str := data.([]byte)
	err = json.Unmarshal(str, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
func CheckAccount(username string, password string) error {
	db := common.GetDB()
	user := User{}
	result := db.Table("users").Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package config

import (
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	PassWord string
}
type Configs struct {
	MysqlConfig MysqlConfig
}

var Config Configs

func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config failed")
	}
	mysql := MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Database: viper.GetString("mysql.database"),
		Username: viper.GetString("mysql.username"),
		PassWord: viper.GetString("mysql.password"),
	}
	Config = Configs{
		MysqlConfig: mysql,
	}
}

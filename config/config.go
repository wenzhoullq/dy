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
type RedisConfig struct {
	Address string
	Network string
	Port    string
	Auth    string
}
type PathConfig struct {
	Videofile string
	Logfile   string
	Picfile   string
}

type MinIoConfig struct {
	Host            string
	Port            string
	AccessKeyID     string
	SecretAccessKey string
	Videobuckets    string
	Picbuckets      string
}

type Configs struct {
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
	Path        PathConfig
	MinIoConfig MinIoConfig
	Level       string
}

var Config Configs

func GetConfig() Configs {
	return Config
}
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
	redis := RedisConfig{
		Address: viper.GetString("redis.address"),
		Network: viper.GetString("redis.network"),
		Port:    viper.GetString("redis.port"),
		Auth:    viper.GetString("redis.auth"),
	}
	Minio := MinIoConfig{
		Host:            viper.GetString("minio.host"),
		Port:            viper.GetString("minio.port"),
		AccessKeyID:     viper.GetString("minio.accessKeyID"),
		SecretAccessKey: viper.GetString("minio.secretAccessKey"),
		Videobuckets:    viper.GetString("minio.videobuckets"),
		Picbuckets:      viper.GetString("minio.picbuckets"),
	}
	PathConfig := PathConfig{
		Videofile: viper.GetString("videofile"),
		Picfile:   viper.GetString("picfile"),
	}
	Config = Configs{
		MysqlConfig: mysql,
		RedisConfig: redis,
		MinIoConfig: Minio,
		Path:        PathConfig,
	}

}

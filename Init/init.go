package Init

import (
	"simple-demo/common"
	"simple-demo/config"
	"simple-demo/log"
)

func Init() {
	config.LoadConfig()
	common.InitMysqlDB()
	common.RedisInit()
	log.InitLog()
}

package Init

import "simple-demo/config"

func init() {
	config.LoadConfig()
	dbInit.InitMysqlDB()
}

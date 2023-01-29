package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"simple-demo/Init"
	"simple-demo/common"
	"syscall"
)

func main() {
	//Init中间件
	Init.Init()
	defer common.CloseDataBase()
	defer common.CloseRedis()
	r := gin.Default()
	//初始化路由
	Init.InitRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	q := make(chan os.Signal)
	//接收ctrl + c ，kill(排除 kill -9)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
}

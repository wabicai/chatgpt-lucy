package main

import (
	"chatgpt-lucy/config"
	"chatgpt-lucy/pkg/logger"
	"chatgpt-lucy/routes"
	"strconv"
	"sync"

	"github.com/alecthomas/kong"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var once sync.Once

func main() {

	kong.Parse(&config.CLI)
	// 路由注册
	once.Do(func() {
		router = gin.Default()
		routes.RegisterWebRoutes(router)
	})

	// 启动服务
	port := config.LoadConfig().Port
	portString := strconv.Itoa(port)
	// 自定义监听地址
	listen := config.LoadConfig().Listen
	err := router.Run(listen + ":" + portString)
	if err != nil {
		logger.Danger("run webserver error %s", err)
		return
	}
}

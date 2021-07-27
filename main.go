package main

import (

	"fmt"
	"go-example/bootstrap"
	"go-example/config"
	"go-example/middlewares"
	"go-example/pkg/gredis"
	"go-example/web/router"
)

func init() {
	gredis.SetUp()
}

func newApp() *bootstrap.Bootstrapper {
	middlewares.SystemLogsSetUp()
	app := bootstrap.New("gin-web", "gin-web-example")
	app.Bootstrap()
	app.Configure(router.Configure)
	return app
}

//增加项目地址
// @termsOfService https://github.com/likai1130/go-example
func main() {
	app := newApp()
	conf := config.AppConfig
	port := conf.Server.Port
	listenPort := fmt.Sprintf(":%v", port) //格式化监听端口
	app.Listen(listenPort)

}

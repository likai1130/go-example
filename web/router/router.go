package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-example/bootstrap"
	"go-example/config"
	"strings"
)

/**
 * 路由配置，并根据配置文件设置根路径
 * 参考url：https://github.com/gin-gonic/gin
 */
func Configure(r *bootstrap.Bootstrapper) {
	prefix := "/"
	//此处可以增加系统应用目录根路径
	pldConf := config.AppConfig
	contextPath := pldConf.Server.ContextPath
	if "" != contextPath && strings.HasPrefix(contextPath, "/") {
		//给拼接好的api ，增加前缀
		prefix = contextPath
	}
	rootRouter := r.Group(prefix) //设置访问的根目录
	concreteRouter(rootRouter)
	docs.SwaggerInfo.Title = "go-example:ONLINE API"
	docs.SwaggerInfo.Description = "This is Demo server online restFull api ."
	docs.SwaggerInfo.Version = "v0.1"
	rootRouter.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

/**
  配置具体的路由信息
*/
func concreteRouter(rootRouter *gin.RouterGroup) {


}

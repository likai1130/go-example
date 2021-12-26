package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-example/common/logger"
	"go-example/config"
	"io"
)

/**
日志中间件
*/
func LoggerToFile(appName string) gin.HandlerFunc {

	mode := config.AppConfig.Logger.Mode
	switch mode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			appName,
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

/**
系统启动日志
*/
func SystemLogsSetUp() {
	logger.GetLogger().SetTimeFormat("2006-01-02 15:04:05")
	//系统日志设置
	if config.AppConfig.Logger.IsOutPutFile {
		path := logger.CreateGinSysLogPath("access")
		write := logger.LogSplite(path)
		gin.DefaultWriter = io.MultiWriter(write)
	}
}

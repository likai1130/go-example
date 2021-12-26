package logger

import (
	"github.com/kataras/golog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go-example/config"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

var logInstance *golog.Logger

func logInit() *golog.Logger {
	logInstance = golog.Default
	logInstance.SetLevel(config.AppConfig.Logger.Level)
	logInstance.SetTimeFormat("2006-01-02 15:04:05")
	if config.AppConfig.Logger.IsOutPutFile == false {
		return logInstance
	}
	logInfoPath := CreateGinSysLogPath("go")
	writer := LogSplite(logInfoPath)
	//设置output
	logInstance.SetOutput(writer)
	return logInstance
}

func GetLogger() *golog.Logger {
	if logInstance == nil {
		logInstance = logInit()
	}
	return logInstance
}

/**
根据时间检测目录，不存在则创建
*/
func createDateDir(folderPath string) string {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// CreateGinSysLogPath 创建系统日志的名字/**
func CreateGinSysLogPath(filePrix string) string {
	baseLogPath := filepath.Join(config.AppConfig.Server.DataPath, "logs/")
	writePath := createDateDir(baseLogPath) //根据时间检测是否存在目录，不存在创建
	fileName := path.Join(writePath, filePrix)
	return fileName
}

// LogSplite 日志分割/**
func LogSplite(logInfoPath string) io.Writer {
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		logInfoPath+"_%Y%m%d.log",
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(logInfoPath),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(time.Duration(config.AppConfig.Logger.MaxAgeDay*24)*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Duration(config.AppConfig.Logger.RotationTime*24)*time.Hour),
	)
	return logWriter
}

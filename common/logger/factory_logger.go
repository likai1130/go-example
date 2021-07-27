package logger

import (
	"github.com/kataras/golog"
	"github.com/lestrrat-go/file-rotatelogs"
	"go-example/config"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Logger struct {
	logger *golog.Logger
}

var pldLoggerInstance *Logger

func NewInstance() *Logger {
	if pldLoggerInstance != nil {
		return pldLoggerInstance
	}
	pldLoggerInstance = &Logger{
		logger: golog.Default,
	}
	pldLoggerInstance.logger.SetLevel(config.AppConfig.Logger.Level)
	if config.AppConfig.Logger.IsOutPutFile == false {
		return pldLoggerInstance
	}
	pldLoggerInstance.logger.SetTimeFormat("2006-01-02 15:04:05")
	logInfoPath := CreateGinSysLogPath("go")

	writer := LogSplite(logInfoPath)
	//设置output
	pldLoggerInstance.logger.SetOutput(writer)
	return pldLoggerInstance
}

func (lf *Logger) GetLogger() *golog.Logger {
	if pldLoggerInstance == nil {
		instance := NewInstance()
		lf.logger = instance.logger
	}
	return lf.logger
}

/**
根据时间检测目录，不存在则创建
*/
func CreateDateDir(folderPath string) string {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

/**
创建系统日志的名字
*/
func CreateGinSysLogPath(filePrix string) string {
	baseLogPath := filepath.Join(config.AppConfig.Server.DataPath, "logs/")
	writePath := CreateDateDir(baseLogPath) //根据时间检测是否存在目录，不存在创建
	fileName := path.Join(writePath, filePrix)
	return fileName
}

/**
日志分割
*/
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

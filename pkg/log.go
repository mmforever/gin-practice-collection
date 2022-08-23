package pkg

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func LogSetup() {
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}

	writer, _ := rotatelogs.New(
		logFilePath+"%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Duration(30*24)*time.Hour),    //保留三分钟之内的文件
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), //每分钟切割一个文件出来
	)

	//设置输出
	logrus.SetOutput(writer)

	//设置在输出日志中添加文件名和方法信息
	// Logger.SetReportCaller(true)

	//设置日志级别
	logrus.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

package utils

import (
	"os"

	logrus "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logrus           = logrus.New()
	logrusLevel      = logrus.DebugLevel // level
	logrusUseJson    = true              // default is text
	logrusToFile     = false             // default to stdout
	logrusToFilePath = "go.log"          // log file
	logrusToFileSize = 10                // 文件轮换 大小以M为单位
)

type LogFields map[string]interface{}

func SetLogLevel(level string) {
	switch level {
	case "Info":
		logrusLevel = logrus.InfoLevel
	case "Warn":
		logrusLevel = logrus.WarnLevel
	case "Error":
		logrusLevel = logrus.ErrorLevel
	default:
		logrusLevel = logrus.DebugLevel // default "Debug"
	}
}

func SetLogUseJson(en bool) {
	logrusUseJson = en
}
func SetLogToFile(en bool, filePath string, fileSize int) {
	logrusToFile = en
	logrusToFilePath = filePath
	logrusToFileSize = fileSize
}

func InitLog() {
	Logrus.SetLevel(logrusLevel)

	if logrusUseJson {
		Logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000", // 自定义时间格式
		}) // json
	} else {
		Logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
			// TimestampFormat: "2006-01-02 15:04:05.000", // 自定义时间格式
		})
	}

	if logrusToFile {
		// 创建一个 RotateFile 钩子，用于轮换日志文件
		rotateFileHook := lumberjack.Logger{
			Filename:   logrusToFilePath, // 日志文件路径
			MaxSize:    logrusToFileSize, // 单个日志文件的最大大小，以 MB 为单位
			MaxBackups: 10,               // 最多保留的旧日志文件数
			MaxAge:     0,                // 保留日志文件的最大天数 0不限制时间
			Compress:   false,            // 是否压缩旧日志文件
		}

		Logrus.SetOutput(&rotateFileHook)
	} else {
		Logrus.SetOutput(os.Stdout) // 输出到控制台
	}
}

func Log(args ...map[string]interface{}) *logrus.Entry {
	if len(args) > 0 {
		return Logrus.WithFields(logrus.Fields(args[0]))
	}
	return Logrus.WithFields(logrus.Fields{})
}

// 使用示例

// // init log config
// utils.SetLogUseJson(true)
// // utils.SetLogToFile(true, "x.log", 1) // 输出到文件，以1M一个文件进行轮换
// utils.InitLog()
// s := []int{1, 2}
// m := map[string]interface{}{
// 	"name": "fish",
// 	"age":  2,
// }

// utils.Log(utils.LogFields{
// 	"name": "fish",
// 	"age":  24,
// }).Info("hello")

// utils.Log(utils.LogFields{
// 	"name": "fish",
// 	"age":  24,
// }).Info(1234)

// utils.Log(utils.LogFields{
// 	"name": "fish",
// 	"age":  24,
// }).Info(s)

// utils.Log().Info("not set fields")
// utils.Log().Debug(m)
// utils.Log().Error("error log")

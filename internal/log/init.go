package log

import (
	"fmt"
	"github.com/jianggujin/EasyWeb/internal/config"
	"github.com/rs/zerolog"
	"io"
	"os"
)

var logger zerolog.Logger

func Init() {
	// log level
	switch config.Config.Advanced.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default:
		panic(fmt.Sprintf("unknown log level: %s", config.Config.Advanced.LogLevel))
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	// 未配置日志文件，只初始化控制台日志
	if config.Config.Advanced.LogFile == "" {
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
		return
	}
	fileWriter, err := getFileWriter()
	if err != nil {
		panic(fmt.Sprintf("create log file failed: %s", err.Error()))
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	logger = zerolog.New(multi).With().Timestamp().Logger()
}

func getFileWriter() (io.Writer, error) {
	// 创建日志目录
	if _, err := os.Stat(config.Config.Advanced.LogDir); os.IsNotExist(err) {
		err := os.MkdirAll(config.Config.Advanced.LogDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	if config.Config.Advanced.LogMode == "single" {
		fileName := config.Config.Advanced.LogDir + "/" + config.Config.Advanced.LogFile + ".log"
		// 打开日志文件，如果不存在则创建
		return os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	return NewRollingWriter()
}

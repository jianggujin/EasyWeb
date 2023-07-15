package log

import (
	"github.com/jianggujin/EasyWeb/internal/config"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
	"time"
)

type rollingWriter struct {
	Out  io.Writer
	Date int
}

func NewRollingWriter() (io.Writer, error) {
	r := new(rollingWriter)
	err := r.resetWriter()
	return r, err
}

func (r *rollingWriter) resetWriter() error {
	r.Close()
	r.Date = time.Now().Day()
	write, err := r.getFileWriter()
	if err == nil {
		r.Out = write
	}
	go clearHistory()
	return err
}

func (r *rollingWriter) Write(p []byte) (n int, err error) {
	currentDate := time.Now().Day()
	if r.Date != currentDate {
		r.resetWriter()
	}
	return r.Out.Write(p)
}

// 关闭打开的文件
func (r *rollingWriter) Close() error {
	if r.Out != nil {
		return r.Out.(zerolog.ConsoleWriter).Out.(io.Closer).Close()
	}
	return nil
}

func (r *rollingWriter) getFileWriter() (io.Writer, error) {
	// 获取当前时间
	t := time.Now()

	// 拼接日志文件名，格式为 "app-20060102.log"
	dateStr := t.Format("20060102")
	fileName := config.Config.Advanced.LogDir + "/" + config.Config.Advanced.LogFile + "-" + dateStr + ".log"

	// 打开日志文件，如果不存在则创建
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// 返回 FileWriter，使用 UTC 时间格式
	return zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = file
		w.TimeFormat = "2006-01-02 15:04:05"
		w.NoColor = true
	}), nil
}

// 清除历史文件
func clearHistory() {
	if config.Config.Advanced.LogMaxHistory < 1 {
		return
	}
	dirs, err := os.ReadDir(config.Config.Advanced.LogDir)
	if err != nil {
		return
	}
	sevenDaysAgo := time.Now().AddDate(0, 0, -config.Config.Advanced.LogMaxHistory).Format("20060102")
	fileName := config.Config.Advanced.LogFile + "-" + sevenDaysAgo + ".log"
	for _, f := range dirs {
		if !f.IsDir() && strings.HasPrefix(f.Name(), config.Config.Advanced.LogFile) && strings.HasSuffix(f.Name(), ".log") && strings.Compare(f.Name(), fileName) < 0 {
			err = os.Remove(config.Config.Advanced.LogDir + "/" + f.Name())
			if err != nil {
				continue
			}
		}
	}
}

package main

import (
	"flag"
	"github.com/jianggujin/EasyWeb/internal/config"
	"github.com/jianggujin/EasyWeb/internal/log"
	"github.com/jianggujin/EasyWeb/internal/server"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// 构建版本信息
var (
	BuildVersion   = "unkown"
	BuildGoVersion = "unkown"
	BuildTime      = "unkown"
)

func printInfo() {
	log.Infof("Easy Web Version: %s", BuildVersion)
	log.Infof("BuildGoVersion: %s, BuildTime: %s", BuildGoVersion, BuildTime)
	log.Infof("GOOS: %s, GOARCH: %s", runtime.GOOS, runtime.GOARCH)
	log.Infof("pid: %d", os.Getpid())
}

// 等待系统信号，当进程被终止时保存必要的数据
func waitSystemSigle() {
	// 创建一个chan用于接收信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	<-sigChan
	log.Infof("Easy Web Goodbye.")
	os.Exit(0)
}

func init() {
	configFile := flag.String("config", "config.toml", "config file, eg: config.toml")
	flag.Parse()
	// load config
	config.LoadFromFile(*configFile)
	log.Init()
}

func main() {
	printInfo()
	go waitSystemSigle()
	server.StartServer()
}

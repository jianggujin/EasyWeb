package config

import (
	"bytes"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type tomlServer struct {
	Port     int               `toml:"port"`
	Static   string            `toml:"static"`
	MineType map[string]string `toml:"mine_type"`
}

type tomlAdvanced struct {
	// log
	LogFile       string `toml:"log_file"`
	LogLevel      string `toml:"log_level"`
	LogMaxHistory int    `toml:"log_max_history"`
	LogDir        string `toml:"log_dir"`
	LogMode       string `toml:"log_mode"`
}

type tomlEasyWebConfig struct {
	Server   tomlServer
	Advanced tomlAdvanced
}

var Config tomlEasyWebConfig

func init() {
	// Server
	Config.Server.Port = 8080
	Config.Server.Static = "static/"
	Config.Server.MineType = make(map[string]string)

	// advanced
	Config.Advanced.LogFile = "easy-web"
	Config.Advanced.LogLevel = "info"
	Config.Advanced.LogDir = "logs"
	Config.Advanced.LogMaxHistory = 7
	Config.Advanced.LogMode = "rolling"
}

func LoadFromFile(filename string) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		// log.Infof("config file [%s] not exists, use default config", filename)
		return
	}

	decoder := toml.NewDecoder(bytes.NewReader(buf))
	decoder.SetStrict(true)
	err = decoder.Decode(&Config)
	if err != nil {
		missingError, ok := err.(*toml.StrictMissingError)
		if ok {
			panic(fmt.Sprintf("decode config error:\n%s", missingError.String()))
		}
		panic(err.Error())
	}

}

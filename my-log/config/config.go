package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"time"
)

const (
	DEFAULT_DIR = "logs"

	DEFAULT_LEVEL = "DEBUG"

	DEFAULT_FORMAT = "YYYY-MM-DD hh:mm"
)

var (
	Conf              config
	DefaultConfigFile = "config/config.toml"
)

type config struct {
	LogConfig logconfig
}

type logconfig struct {
	OutputDir    string `toml:"output_dir"`
	OutputFormat string `toml:"output_format"`
	LogLevel     string `toml:"log_level"`
	LogPreFix    string `toml:"log_pre_fix"`
	LogSufFix    string `toml:"log_suf_fix"`
	Console      bool   `toml:"console"`
}

func init() {
	InitConfig("")
}

func LogName() string {

	fPath, err := os.Getwd()
	var bt bytes.Buffer
	if err != nil {
		fmt.Printf(err.Error())
		bt.WriteString(DEFAULT_DIR)

	} else {
		bt.WriteString(fPath)
		bt.WriteString("/")
		bt.WriteString(Conf.LogConfig.OutputDir)
	}

	bt.WriteString("/log-")
	bt.WriteString(time.Now().Format(Conf.LogConfig.LogPreFix))
	bt.WriteString(".")
	bt.WriteString(Conf.LogConfig.LogSufFix)
	return bt.String()
}

func InitConfig(configFile string) error {

	if configFile == "" {
		configFile = DefaultConfigFile
	}
	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config load err:" + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)

		if err != nil {
			return errors.New("config decode err:" + err.Error())
		}
	}
	return nil
}

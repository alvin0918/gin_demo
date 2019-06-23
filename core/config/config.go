package config

import (
	"gopkg.in/ini.v1"
	"github.com/alvin0918/gin_demo/core"
)

var Cfg *ini.File

// 读取配置文件
func init() {

	var (
		defaultConfigPath string
		cfg *ini.File
		err error
	)

	defaultConfigPath = "./config/app.ini"

	if !core.FileExists(defaultConfigPath) {
		panic("The default configuration file ["+defaultConfigPath+"] does not exist")
	}

	if cfg, err = ini.Load(defaultConfigPath); err != nil {
		panic("The configuration read failed!")
	}

	Cfg = cfg

}

// 获取字符串
func GetString(key string, section string) string {
	return Cfg.Section(section).Key(key).String()
}
















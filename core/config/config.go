package config

import (
	"gopkg.in/ini.v1"
	"github.com/alvin0918/gin_demo/core/commin"
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

	if !commin.FileExists(defaultConfigPath) {
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

// 获取系统IP Port
func GetIpAndPort(section string) string {
	switch section {
	case "Server":
		return GetString("IP", "Server") + ":" + GetString("Port", "Server")

	default:
		return GetString("IP", "Server") + ":" + GetString("Port", "Server")
	}

}

// 获取分区下的内容
func GetSection(section string) (*ini.Section, error) {
	return Cfg.GetSection(section)
}















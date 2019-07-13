package log

import (
	"github.com/aiwuTech/fileLogger"
	"github.com/alvin0918/gin_demo/core/config"
	"gopkg.in/ini.v1"
	"time"
	"encoding/json"
	"fmt"
)

var (
	TRACE   *fileLogger.FileLogger
	INFO    *fileLogger.FileLogger
	WARN    *fileLogger.FileLogger
	ERROR   *fileLogger.FileLogger
)

type Data struct {
	Mode 		string `json:"mode"`
	Ttl  		string `json:"@timestamp"`
	Data 		string `json:"data"`
	Code 		int64  `json:"code"`
	Uri 		string `json:"uri"`
	RemodeAddr 	string `json:"remode_addr"`
	Request 	string `json:"request"`
	Agent 		string `json:"agent"`
}


func init() {

	var (
		section *ini.Section
		err error
		LoggerPath string
		t time.Time
		t_str string
	)

	if section, err = config.GetSection("Logger"); err != nil {
		panic(err)
	}

	LoggerPath = section.Key("LoggerPath").String()

	t = time.Now()

	t_str = fmt.Sprintf("_%d_%d_%d_", t.Year(), t.Month(), t.Day())

	TRACE = fileLogger.NewDefaultLogger(LoggerPath, "trace"+t_str+".log")
	INFO  = fileLogger.NewDefaultLogger(LoggerPath, "info"+t_str+".log")
	WARN  = fileLogger.NewDefaultLogger(LoggerPath, "warn"+t_str+".log")
	ERROR = fileLogger.NewDefaultLogger(LoggerPath, "error"+t_str+".log")

}

func TracePrintf(data string, v ...interface{})  {


	var (
		res string
	)

	res = makeJson(data, "Trace", v)

	defer TRACE.Close()
	TRACE.Print(res)
}


func InfoPrintf(data string, v ...interface{})  {

	var (
		res string
	)

	res = makeJson(data, "Info", v)

	defer INFO.Close()
	INFO.Print(res)
}

func WarnPrintf(data string, v ...interface{})  {

	var (
		res string
	)

	res = makeJson(data, "Warn", v)

	defer WARN.Close()
	WARN.Print(res)
}

func ErrorPrintf(data string, v ...interface{})  {

	var (
		res string
	)

	res = makeJson(data, "Error", v)

	defer ERROR.Close()
	ERROR.E(res)
}

// code int64, uri string, remode_addr string, request string, agent string
func makeJson(str string, mode string, v ...interface{}) string {

	var (
		data *Data
		res []byte
		err error
	)

	data = &Data{
		Mode: mode,
		Ttl:  time.Unix(time.Now().Unix(), 0).String(),
		Data: str,
	}

	if res, err = json.Marshal(data); err != nil {
		panic(err)
	}

	return string(res)

}










package log

import (
	"time"
	"encoding/json"
	"gopkg.in/ini.v1"
	"os"
	"log"
	"github.com/alvin0918/gin_demo/core/config"
	"fmt"
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

func TracePrintf(f string, data string, v ...interface{})  {
	go func(f string, data string, v ...interface{}) {

		var (
			section *ini.Section
			err error
			LoggerPath string
			t time.Time
			t_str string
			name string
			file *os.File
			trace *log.Logger
		)

		if section, err = config.GetSection("Logger"); err != nil {
			panic(err)
		}

		LoggerPath = section.Key("LoggerPath").String()

		t = time.Now()

		t_str = fmt.Sprintf("_%d_%d_%d.log", t.Year(), t.Month(), t.Day())

		name = LoggerPath + "/" + f + t_str

		if file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err != nil {
			panic(err)
		}

		trace 	= log.New(file, "[" + f + "] ", log.Ldate|log.LUTC|log.Lmicroseconds|log.Llongfile)

		trace.Print(data)

		trace.Writer()

		defer file.Close()

	}(f, data, v)
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










package commin

import (
	"os"
	"fmt"
)

// 判断文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 错误处理
func RecoverName() {

	var (
		r interface{}
	)

	if r = recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}

package core

import (
	"runtime"
	"github.com/gin-gonic/gin"
	"github.com/alvin0918/gin_demo/routers"
	"github.com/alvin0918/gin_demo/core/config"
	_ "github.com/alvin0918/gin_demo/core/commin/log"
)

func Run(){

	var(
		r *gin.Engine
	)

	// 设置使用线程数
	initEnv()

	// 注册中间件
	r = gin.Default()

	// 加载路由
	routers.Init(r)

	// 开始运行
	r.Run(config.GetIpAndPort("Server"))

}

/**
	初始化线程数
 */
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

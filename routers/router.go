package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/alvin0918/gin_demo/controller"
)

func Init(r *gin.Engine)  {
	r.GET("/abc", controller.Index)
}

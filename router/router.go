package router

import (
	"github.com/GpsLypy/ginEssentail/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/Login", controller.Login)
	return r
}

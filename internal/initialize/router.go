package initialize

import (
	"github.com/duongha/go-ecommerce/global"
	"github.com/duongha/go-ecommerce/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	r.Use() //logger
	r.Use() // cross
	r.Use() // limit
	adminRouter := router.RouterGroupApp.Admin
	userRouter := router.RouterGroupApp.User

	mainGroup := r.Group("/v1")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.User.InitUserRouter(mainGroup)
		userRouter.Product.InitProductRouter(mainGroup)
	}
	{
		adminRouter.Admin.InitAdminRouter(mainGroup)
		adminRouter.User.InitUserRouter(mainGroup)

	}
	return r
}

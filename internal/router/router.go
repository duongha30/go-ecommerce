package router

import (
	"github.com/duongha/go-ecommerce/internal/controller"
	"github.com/duongha/go-ecommerce/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	//use the middleware
	r.Use(middleware.AuthenMiddleware())

	v1 := r.Group("/v1")
	{ // group endpoint for versioning v1
		v1.GET("/user", controller.NewUserController().GetUserById) // GET /v1/ping
	}
	return r
}

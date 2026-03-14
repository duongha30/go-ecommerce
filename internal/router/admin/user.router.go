package admin

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	// adminRouterPublic := Router.Group("/user")
	// {
	// 	adminRouterPublic.POST("/register")
	// 	adminRouterPublic.POST("/otp")
	// }
	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}

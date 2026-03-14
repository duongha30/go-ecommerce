package router

import (
	"github.com/duongha/go-ecommerce/internal/router/admin"
	"github.com/duongha/go-ecommerce/internal/router/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)

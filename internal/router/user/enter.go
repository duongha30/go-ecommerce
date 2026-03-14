package user

type UserRouterGroup struct {
	Product ProductRouter
	User    UserRouter
}

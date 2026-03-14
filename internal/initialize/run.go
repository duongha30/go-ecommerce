package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run()
}

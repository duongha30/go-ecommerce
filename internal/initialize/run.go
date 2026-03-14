package initialize

import (
	"github.com/duongha/go-ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run()
}

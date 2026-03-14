package initialize

import (
	"github.com/duongha/go-ecommerce/global"
	"github.com/duongha/go-ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}

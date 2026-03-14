package global

import (
	"github.com/duongha/go-ecommerce/pkg/logger"
	"github.com/duongha/go-ecommerce/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	// Mdb    *gorm.DB
)

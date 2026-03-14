package global

import (
	"github.com/duongha/go-ecommerce/pkg/logger"
	"github.com/duongha/go-ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

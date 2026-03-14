package initialize

import (
	"fmt"
	"time"

	"github.com/duongha/go-ecommerce/global"
	"github.com/duongha/go-ecommerce/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(fmt.Sprintf("%s: %v", errString, err))
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%v)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: true, // default is true, if data is not simultaneously updated, set to false will improve 30% perf
	})
	checkErrorPanic(err, "InitMySQL init error")
	global.Mdb = db

	//set connection pool
	SetPool()
	// migrate database
	migrateDatabase()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDB, err := global.Mdb.DB()
	checkErrorPanic(err, "SetPool init error")
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func migrateDatabase() {
	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	checkErrorPanic(err, "migrateDatabase error")
}

package initialize

import (
	"context"
	"fmt"

	"github.com/duongha/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})
	_, err := rdb.Ping(ctx).Result()
	checkErrorPanic(err, "Failed to connect to Redis")

	fmt.Println("Initialized Redis successfully")
	global.Rdb = rdb
}

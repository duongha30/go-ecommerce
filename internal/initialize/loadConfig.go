package initialize

import (
	"fmt"

	"github.com/duongha/go-ecommerce/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config %w \n", err)) // stop the program
	}
	// read config
	fmt.Printf("Server port: %d \n", viper.GetInt("server.port"))
	fmt.Printf("JWT key: %s \n", viper.GetString("security.jwt.key"))

	// config struct
	if err := viper.Unmarshal(&global.Config); err != nil { // if with init statement - notion
		fmt.Printf("Failed to unmarshal config %w \n", err)
	}
}

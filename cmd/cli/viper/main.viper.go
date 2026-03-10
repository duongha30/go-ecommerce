package main

import (
	"fmt"

	"github.com/spf13/viper" // viper is a library for reading configuration files in Go
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
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
	var config Config
	if err := viper.Unmarshal(&config); err != nil { // if with init statement - notion
		fmt.Printf("Failed to unmarshal config %w \n", err)
	}

	fmt.Println("Config port: ", config.Server.Port)

	for _, db := range config.Database { // loop through the database slice and print the configuration
		fmt.Printf("Database user: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}

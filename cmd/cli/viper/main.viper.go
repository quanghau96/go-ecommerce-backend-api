package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbName"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration file
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// get configuration values
	sever := viper.GetInt("server.port")

	fmt.Printf("Server port: %d\n", sever)

	// config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	fmt.Printf("Server port: %d\n", config.Server.Port)

	// fmt.Println("config::: ", config.Database[0].Username, config.Database[0].Password)
	for i, db := range config.Database {
		fmt.Printf("Database %d\n", i)
		fmt.Println("Username:", db.Username)
		fmt.Println("Password:", db.Password)
		fmt.Println("Host:", db.Host)
		fmt.Println("DBName:", db.DBName)
	}
}

// viper.SetConfigName("config") // name of config file (without extension)
// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
// viper.AddConfigPath(".")      // optionally look for config in the working directory

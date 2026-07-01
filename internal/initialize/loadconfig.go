package initialize

import (
	"github.com/quanghau96/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration file
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// get configuration values
	// sever := viper.GetInt("server.port")

	// fmt.Printf("Server port: %d\n", sever)

	// config structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}

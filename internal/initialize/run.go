package initialize

import (
	"fmt"

	"github.com/quanghau96/go-ecommerce-backend-api/global"
)

func Run() {
	// LoadConfig()
	LoadConfig()
	fmt.Println("global.Config.MySQL.Username: ", global.Config.MySql.Username)
	// InitLogger()
	InitLogger()
	global.Logger.Info("Logger initialized successfully")
	// InitMysql()
	InitMysql()
	// InitRedis()
	InitRedis()
	// InitRouter()
	r := InitRouter()
	r.Run(":8002")
}

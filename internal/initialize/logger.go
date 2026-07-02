package initialize

import (
	"github.com/quanghau96/go-ecommerce-backend-api/global"
	"github.com/quanghau96/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLoggerZap(global.Config.Logger)
}

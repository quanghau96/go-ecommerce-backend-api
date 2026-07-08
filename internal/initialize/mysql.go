package initialize

import (
	"fmt"
	"time"

	"github.com/quanghau96/go-ecommerce-backend-api/global"
	"github.com/quanghau96/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	checkErrorPanic(err, "Failed to connect to database")

	global.Logger.Info("Database connection established successfully")
	global.Mdb = db

	fmt.Println("Db::: ", db)

	// set Pool
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.MySql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("My sql error: %s::", err)
	}

	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}

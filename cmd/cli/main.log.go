package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "Tips go", 40)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Tipgo"), zap.Int("age", 40))

	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello New Developement")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core)

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format log

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts--> time
	encodeConfig.TimeKey = "time"

	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller "cli/main.log.go:20
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

//

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)

	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

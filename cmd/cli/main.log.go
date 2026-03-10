package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder { // Format log
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // Convert milliseconds to human-readable time format
	encoderConfig.TimeKey = "timestamp"                     // Change "ts" to "timestamp"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Convert log level to uppercase (info -> INFO)
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // Shorten caller information (e.g., main.go:10)
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr) // notion
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

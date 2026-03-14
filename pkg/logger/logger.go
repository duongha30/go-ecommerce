package logger

import (
	"os"

	"github.com/duongha/go-ecommerce/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level // debug > info > warn > error > fatal > panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "panic":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder { // Format log
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // Convert milliseconds to human-readable time format
	encoderConfig.TimeKey = "timestamp"                     // Change "ts" to "timestamp"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Convert log level to uppercase (info -> INFO)
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // Shorten caller information (e.g., main.go:10)
	return zapcore.NewJSONEncoder(encoderConfig)
}

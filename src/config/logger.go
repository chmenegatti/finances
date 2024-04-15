package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(loglevel zapcore.Level, fileName string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Capitalize the log level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC timestamp format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Short caller (file and line)
	}

	defer func() {
		err := logger.Sync()
		if err != nil {

		}
	}()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		loglevel,
	)

	var (
		file *os.File
		err  error
	)

	if _, err = os.Stat("log"); os.IsNotExist(err) {
		err = os.Mkdir("log", 0755)
		if err != nil {
			return
		}

		file, err = os.Create("log/" + fileName)
		if err != nil {
			if _, err = fmt.Fprintf(os.Stderr, "error on create log file: %v\n", err); err != nil {
				return
			}
			os.Exit(1)
		}
	} else {
		if file, err = os.OpenFile("log/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			if _, err = fmt.Fprintf(os.Stderr, "error on open log file: %v\n", err); err != nil {
				return
			}
			os.Exit(1)
		}
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(file)),
		loglevel,
	)

	logger = zap.New(zapcore.NewTee(core, fileCore))

}

// LogDebug debug message
func LogDebug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

// LogInfo info message
func LogInfo(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

// LogWarn warning message
func LogWarn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

// LogError error message
func LogError(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

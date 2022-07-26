package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New create a logger.
func New(config Config) *zap.Logger {
	level := prepLoggerLevel(&config) // Setting the zap logger level
	encoder := prepLoggerEncoder()    // Setting the zap logger encoder
	options := getLoggerOptions()     // Getting the zap logger options

	// Preparing default core of zap logger to write logs into write syncer
	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), level)
	cores := []zapcore.Core{
		defaultCore,
	}

	// Creating the logger based on the core and options
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, options...)

	return logger
}

func prepLoggerLevel(config *Config) zap.AtomicLevel {
	var level zapcore.Level

	// Level set error logging, set default logging level to warning level
	if err := level.Set(config.Level); err != nil {
		log.Printf("cannot parse log level %s: %s", config.Level, err)

		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	}

	return zap.NewAtomicLevelAt(level)
}

func prepLoggerEncoder() zapcore.Encoder { //nolint:ireturn
	// The default encoder configs are a console capital color level encoder
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLoggerOptions() []zap.Option {
	// Zap logger options
	options := make([]zap.Option, 0)

	// Two main logger options are caller and stacktrace
	options = append(options, zap.AddCaller())
	options = append(options, zap.AddStacktrace(zap.ErrorLevel))

	return options
}

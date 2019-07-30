package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	serviceConfig "internal/config"
	"log"
)

func CreateLogger() (logger *zap.Logger, err error) {

	globalConfig, err := serviceConfig.ReadConfig()
	if err != nil {
		return nil, err
	}

	var level zap.AtomicLevel
	switch globalConfig.Level {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warning":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	if len(globalConfig.OutputPaths) == 0 {
		globalConfig.OutputPaths = []string{"stdout"}
	}
	if len(globalConfig.ErrorOutputPaths) == 0 {
		globalConfig.ErrorOutputPaths = []string{"stderr"}
	}

	zapConfig := zap.Config{
		Encoding:         globalConfig.Encoding,
		Level:            level,
		OutputPaths:      globalConfig.OutputPaths,
		ErrorOutputPaths: globalConfig.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			MessageKey:   "message",
		},
	}

	logger, err = zapConfig.Build()
	if err != nil {
		log.Printf("Config build error: %v \n", err)
		return nil, err
	}
	return logger, nil

}

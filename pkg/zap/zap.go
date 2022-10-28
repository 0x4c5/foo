package zap

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (logger *zap.Logger, err error) {

	info := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})
	core := zapcore.NewCore(encoder("CONSOLE"), syncer(), info)
	logger = zap.New(core)
	zap.ReplaceGlobals(logger)
	return
}

func encoder(format string) zapcore.Encoder {
	conf := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ctime",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stack",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	if strings.ToLower(format) == "json" {
		return zapcore.NewJSONEncoder(conf)
	}
	return zapcore.NewConsoleEncoder(conf)
}

func syncer() zapcore.WriteSyncer {
	// only out to console
	return zapcore.AddSync(os.Stdout)
}

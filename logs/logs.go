package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = "" // disable stack trace

	var err error
	log, err = config.Build(zap.AddCallerSkip(1)) // skip caller 1 step
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch value := message.(type) {
	case error:
		log.Error(value.Error(), fields...)
	case string:
		log.Error(value)
	}

}

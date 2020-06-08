package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GCPConfig ...
func GCPConfig(level zapcore.Level) zap.Config {
	levelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		gcpLevel := level.CapitalString()

		switch level {
		case zap.WarnLevel:
			gcpLevel = "WARNING"
			break
		case zap.DPanicLevel:
		case zap.PanicLevel:
			gcpLevel = "CRITICAL"
			break
		case zap.FatalLevel:
			gcpLevel = "EMERGENCY"
			break
		}

		enc.AppendString(gcpLevel)
	}

	return zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:      "logName",
			MessageKey:   "textPayload",
			LevelKey:     "severity",
			EncodeLevel:  levelEncoder,
			TimeKey:      "timestamp",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
}

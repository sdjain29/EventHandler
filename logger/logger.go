package logger

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func SetupLogger() {
	var err error

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func GinInfoLog(c context.Context, s interface{}) {
	InfoLogWithType(c, s, "GIN")
}

func Debug(c context.Context, s interface{}) {
	DebugLogWithType(c, s, "DEBUG")
}

func Info(c context.Context, s interface{}) {
	InfoLogWithType(c, s, "INFO")
}

func Error(c context.Context, s interface{}) {
	ErrorLogWithType(c, s, "ERROR")
}

func HttpError(c context.Context, s interface{}) {
	ErrorLogWithType(c, s, "HTTPERROR")
}

func HttpClientInfo(c context.Context, s interface{}) {
	InfoLogWithType(c, s, "HTTPCLIENT")
}

func InfoLogWithType(c context.Context, s interface{}, loggerType string) {
	if c.Value("requestid") == nil {
		Logger.Info("", zap.String("requestId", "No Context RequestId"), zap.String("loggerType", loggerType), zap.Any("message", s))
		return
	}
	Logger.Info("", zap.String("requestId", c.Value("requestid").(string)), zap.String("loggerType", loggerType), zap.Any("message", s))
}

func DebugLogWithType(c context.Context, s interface{}, loggerType string) {
	if c.Value("requestid") == nil {
		Logger.Info("", zap.String("requestId", "No Context RequestId"), zap.String("loggerType", loggerType), zap.Any("message", s))
		return
	}
	Logger.Debug("", zap.String("requestId", c.Value("requestid").(string)), zap.String("loggerType", loggerType), zap.Any("message", s))
}

func ErrorLogWithType(c context.Context, s interface{}, loggerType string) {
	if c.Value("requestid") == nil {
		Logger.Info("", zap.String("requestId", "No Context RequestId"), zap.String("loggerType", loggerType), zap.Any("message", s))
		return
	}
	Logger.Error("", zap.String("requestId", c.Value("requestid").(string)), zap.String("loggerType", loggerType), zap.Any("message", s))
}

func filterLogBody(body string) string {
	if strings.Contains(strings.ToLower(body), "password") && os.Getenv("ENV") == "PROD" {
		return "CONTAINS SOME KIND OF PASSWORD"
	}
	if strings.Contains(strings.ToLower(body), "pan") && os.Getenv("ENV") == "PROD" {
		return "CONTAINS SENSIVE INFORMATION"
	}
	return body
}

func filterLogResponseBody(body string) string {
	if strings.Contains(strings.ToLower(body), "password") && os.Getenv("ENV") == "PROD" {
		return "CONTAINS SOME KIND OF PASSWORD"
	}
	return body
}

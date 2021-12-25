/*
 * File: zap.go
 * Created Date: 2021-12-25 02:02:28
 * Author: ysj
 * Description: zap 日志中间价
 */

package zap

import (
	"os"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"grpc-notes/conf"
)

var (
	runMode    = conf.Get("mode")
	zapConfig  zapcore.EncoderConfig
	stackLevel zapcore.Level
	ws         zapcore.WriteSyncer
)

// ZapInterceptor 返回zap.logger实例(把日志写到文件中)
func ZapInterceptor() *zap.Logger {
	if runMode == "dev" {
		ws = zapcore.AddSync(os.Stdout)
		zapConfig = zap.NewDevelopmentEncoderConfig()
		stackLevel = zapcore.WarnLevel
	} else {
		ws = zapcore.AddSync(&lumberjack.Logger{
			Filename:  "logs/log.log",
			MaxSize:   200, //MB
			LocalTime: true,
		})
		zapConfig = zap.NewProductionEncoderConfig()
		stackLevel = zapcore.ErrorLevel
	}

	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(zapConfig),
		zapcore.NewConsoleEncoder(zapConfig),
		ws,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core,
		zap.AddCaller(),               // 输出文件名和行号
		zap.AddCallerSkip(1),          // 让输出的文件名和行号是调用函数的位置
		zap.AddStacktrace(stackLevel), //输出调用堆栈
	)
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}

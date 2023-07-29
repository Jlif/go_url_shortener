package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger // SugarLogger 需要全局使用记录日志

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	// zapcore.DebugLevel: 日志级别
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller()) // zap.AddCaller(): 记录调用函数
	SugarLogger = Logger.Sugar()            // 通过主 logger 的方法获取 SugarLogger
}

// getEncoder 获取编码器 (如何写入日志)
// 时间使用人类可读的方式
// 使用大写字母标识日志级别
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 使用控制台或者 json 格式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 指定日志输出
func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("logs/zap.log")
	//fileSyncer := zapcore.AddSync(file)
	consoleSyncer := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(consoleSyncer)
}

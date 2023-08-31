package logger

import (
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var once sync.Once

// Initialize 初始化日志组件
func Initialize(outPutPath string, isDevelopment bool) *zap.Logger {
	once.Do(func() {
		logOutputFile := outPutPath
		logDevelopment := isDevelopment
		logger = createZapLogger(logOutputFile, logDevelopment)
	})

	return logger
}

// return a zap with lumberjack logger
func createZapLogger(logOutputFile string, logDevelopment bool) *zap.Logger {
	var zapLogger *zap.Logger

	hook := lumberjack.Logger{
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
		Filename:   logOutputFile,
	}

	write := zapcore.AddSync(&hook)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		write,
		zap.NewAtomicLevel(),
	)

	if logDevelopment {
		development := zap.Development()
		zapLogger = zap.New(core, development, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	} else {
		zapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	}

	return zapLogger
}

func Info(msg string, fields ...zap.Field) {
	if logger == nil {
		panic("logger not init.")
	}

	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	if logger == nil {
		panic("logger not init.")
	}

	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	if logger == nil {
		panic("logger not init.")
	}

	logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	if logger == nil {
		panic("logger not init.")
	}

	logger.Debug(msg, fields...)
}

func DatabaseError(err error, fields ...zap.Field) {
	if logger == nil {
		panic("logger not init.")
	}

	// fields的头部添加一个error字段
	fields = append([]zap.Field{zap.Error(err)}, fields...)

	logger.Error("Error encountered during database operation", fields...)
}

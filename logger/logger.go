package logger

import (
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/FengWuTech/commons/setting"
)

// error logger
var errorLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// Setup initialize the log instance
func Setup() {

	//获取日志的存储路径
	filePath := getLogFilePath()
	fileName := getLogFileName()
	fileFullName := filePath + fileName

	//获取配置的日志级别
	logLevel := setting.AppSetting.LogLevel
	level := getLoggerLevel(logLevel)
	debugLevel := zap.DebugLevel

	//日志切割配置
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileFullName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
		MaxAge:    30, //day
	})

	//日志编码配置
	var encoder zapcore.EncoderConfig
	runMode := setting.ServerSetting.RunMode
	//if runMode == "debug" {
	//	encoder = zap.NewDevelopmentEncoderConfig()
	//} else {
	encoder = zap.NewProductionEncoderConfig()
	//}

	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	//真正的来配置zap
	core := zapcore.NewTee(
		zapcore.NewCore( //文件，json
			zapcore.NewJSONEncoder(encoder),
			syncWriter,
			zap.NewAtomicLevelAt(level),
		),
		zapcore.NewCore( // 控制台
			zapcore.NewConsoleEncoder(encoder),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(debugLevel),
		),
	)

	var logger *zap.Logger
	additionalFields := zap.Fields(
		zap.Int("pid", os.Getpid()),
		zap.String("process", path.Base(os.Args[0])),
	)
	if runMode == "debug" {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development(), additionalFields)
	} else {
		logger = zap.New(core)
	}
	//logger := zap.New(core, zap.AddCaller(), zap.Development())
	errorLogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
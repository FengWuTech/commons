package logger

import (
	"crypto/md5"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"path"
	"runtime"
	"strings"

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

func NewRedisWriter(key string, cli *redis.Client) *redisWriter {
	return &redisWriter{
		cli: cli, listKey: key,
	}
}

// 为 logger 提供写入 redis 队列的 io 接口
type redisWriter struct {
	cli     *redis.Client
	listKey string
}

func (w *redisWriter) Write(p []byte) (int, error) {
	n, err := w.cli.RPush(w.listKey, p).Result()
	return int(n), err
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

	// 只线上环境推
	if setting.AppSetting.LogPushToRedis {
		redisCli := redis.NewClient(&redis.Options{
			Addr:     setting.RedisSetting.Host,
			Password: setting.RedisSetting.Password,
			DB:       3,
		})
		redisWriter := NewRedisWriter(setting.AppSetting.LogSaveName+"_list", redisCli)
		core = zapcore.NewTee(
			zapcore.NewCore( //文件，json
				zapcore.NewJSONEncoder(encoder),
				syncWriter,
				zap.NewAtomicLevelAt(level),
			),
			zapcore.NewCore( //redis队列
				zapcore.NewJSONEncoder(encoder),
				zapcore.AddSync(redisWriter),
				zap.NewAtomicLevelAt(level),
			),
			zapcore.NewCore( // 控制台
				zapcore.NewConsoleEncoder(encoder),
				zapcore.AddSync(os.Stdout),
				zap.NewAtomicLevelAt(debugLevel),
			),
		)
	}

	var logger *zap.Logger
	additionalFields := zap.Fields(
		zap.Int("pid", os.Getpid()),
		zap.String("process", path.Base(os.Args[0])),
	)
	if runMode == "debug" {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development(), additionalFields)
	} else {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), additionalFields)
	}
	//logger := zap.New(core, zap.AddCaller(), zap.Development())
	errorLogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.Debug(param...)
}

func Debugf(template string, args ...interface{}) {
	param := append([]interface{}{GID()}, args...)
	errorLogger.Debugf("%v "+template, param...)
}

func Info(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.Info(param...)
}

func Infof(template string, args ...interface{}) {
	param := append([]interface{}{GID()}, args...)
	errorLogger.Infof("%v "+template, param...)
}

func Warn(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.Warn(param...)
}

func Warnf(template string, args ...interface{}) {
	param := append([]interface{}{GID()}, args...)
	errorLogger.Warnf("%v "+template, param...)
}

func Error(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.Error(param...)
}

func Errorf(template string, args ...interface{}) {
	param := append([]interface{}{GID()}, args...)
	errorLogger.Errorf("%v "+template, param...)
}

func DPanic(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.DPanic(param...)
}

func DPanicf(template string, args ...interface{}) {
	param := append([]interface{}{GID()}, args...)
	errorLogger.DPanicf("%v "+template, param...)
}

func Panic(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.Panic(param...)
}

func Panicf(template string, args ...interface{}) {
	param := append([]interface{}{GID()}, args...)
	errorLogger.Panicf("%v"+template, param...)
}

func Fatal(args ...interface{}) {
	param := append([]interface{}{GID(), " "}, args...)
	errorLogger.Fatal(param...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}

func GID() string {
	var buff = make([]byte, 128)
	runtime.Stack(buff, true)

	gid := strings.Split(string(buff), " ")[1]
	return fmt.Sprintf("gid-%x", md5.Sum([]byte(gid)))
}

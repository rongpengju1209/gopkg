package logs

import (
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	initLogger(defaultOption())
}

// InitWithConfig 自定义日志的配置选项
func InitWithConfig(cfg *LogConfig) {
	initLogger(cfg)
}

func initLogger(cfg *LogConfig) {
	writeSyncer := zapcore.NewMultiWriteSyncer(getLogWriter(cfg), zapcore.AddSync(os.Stderr))
	encoder := getEncoder()
	var enab = new(zapcore.Level)
	err := enab.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		log.Fatal(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, enab)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getLogWriter(lc *LogConfig) zapcore.WriteSyncer {
	name := time.Now().Local().Format("2006-01-02-15") + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   lc.Path + name,
		MaxSize:    lc.MaxSize,
		MaxBackups: lc.MaxBackups,
		MaxAge:     lc.MaxAge,
		Compress:   lc.Compress,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Sugar().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Sugar().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Sugar().Panicf(template, args...)
}

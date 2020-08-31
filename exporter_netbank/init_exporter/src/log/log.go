package log

import (
	"flag"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger
var sugarLoggerWithOnceCallerAdd *zap.SugaredLogger
var (
	logLevel, logPath, logFileName string
	logSize                        int
	logBackups                     int
	logAge                         int
	isCallerVisible                bool
)

func init() {
	flag.StringVar(&logLevel, "log.level", "info", "log levels:debug/info/warn/error/dpanic/panic/fatal")
	flag.StringVar(&logPath, "log.path", "/tmp", "log save path")
	flag.IntVar(&logSize, "log.size", 10,
		"MaxSize is the maximum size in megabytes of the log file before it gets rotated. It defaults to 10 megabytes.")
	flag.IntVar(&logBackups, "log.backups", 5,
		"MaxBackups is the maximum number of old log files to retain.")
	flag.IntVar(&logAge, "log.age", 7,
		"MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename.")
	flag.StringVar(&logFileName, "log.filename", "default", "log file name")
	flag.BoolVar(&isCallerVisible, "log.caller", true, "log the caller or not")

}

func InitLogger() {

	encoder := getEncoder()

	var realLogLevel zapcore.Level
	if err := realLogLevel.Set(logLevel); err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), realLogLevel)

	if isCallerVisible {
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		sugarLogger = logger.Sugar()
		sugarLoggerWithOnceCallerAdd = zap.New(core, zap.AddCaller()).Sugar()
	} else {
		sugarLogger = zap.New(core).Sugar()
		sugarLoggerWithOnceCallerAdd = sugarLogger
	}

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args)
}
func Error(args ...interface{}) {
	sugarLogger.Error(args)
}
func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args)
}
func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func With(args ...interface{}) *zap.SugaredLogger {
	return sugarLoggerWithOnceCallerAdd.With(args...)
}

func WithField(args ...interface{}) *zap.SugaredLogger {
	return sugarLoggerWithOnceCallerAdd.With(args...)
}

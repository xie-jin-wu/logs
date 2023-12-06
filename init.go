package logs

import (
	"errors"
	"github.com/xie-jin-wu/program"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type logger struct {
	channel chan struct{}   //阻塞管道
	level   LogLevel        //输出日志等级
	target  LogOutputTarget //日志输出目标
	value   int             //日志分割值(时间分割模式:时间值,大小分割模式:文件大小值)
	dir     string          //日志文件目录名

	debug       *zap.SugaredLogger
	info        *zap.SugaredLogger
	error       *zap.SugaredLogger
	dpanic      *zap.SugaredLogger
	debugStack  *zap.SugaredLogger
	infoStack   *zap.SugaredLogger
	errorStack  *zap.SugaredLogger
	dpanicStack *zap.SugaredLogger
}

// NewLogger
// 输出到文件中默认按天进行分割日志
func NewLogger(level LogLevel, opt ...Options) (Logger, error) {
	if level < DebugLevel || level > DPanicLevel {
		return nil, errors.New("log level error")
	}
	var l = new(logger)
	l.level = level
	l.channel = make(chan struct{}, 1)
	for _, v := range opt {
		v.apply(l)
	}
	err := l.initLogger()
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Logger 日志接口
type Logger interface {
	Debug(msg ...any)
	Debugf(msg string, args ...any)
	StackDebug(msg ...any)
	StackDebugf(msg string, args ...any)
	Info(msg ...any)
	Infof(msg string, args ...any)
	StackInfo(msg ...any)
	StackInfof(msg string, args ...any)
	Error(msg ...any)
	Errorf(msg string, args ...any)
	StackError(msg ...any)
	StackErrorf(msg string, args ...any)
	DPanic(msg ...any)
	DPanicf(msg string, args ...any)
	StackDPanic(msg ...any)
	StackDPanicf(msg string, args ...any)
}

// 初始化日志
func (l *logger) initLogger() error {
	l.value = time.Now().Day()
	config := zap.NewProductionEncoderConfig()
	//将日志时间格式化
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	//将日志等级大写
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(config)
	return l.output(encoder)
}

func (l *logger) output(encoder zapcore.Encoder) error {
	switch l.target {
	case outputFile:
		err := os.MkdirAll(l.dir, os.ModePerm)
		if err != nil {
			return err
		}
		name, err := program.GetProgramName()
		if err != nil {
			return err
		}
		//debug
		filename := l.dir + "/" + name + "_" +
			time.Now().Format(time.DateOnly) +
			"_" + logLevelString[DebugLevel] + ".log"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(file))
		l.debug = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.debugStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		//info
		filename = l.dir + "/" + name + "_" +
			time.Now().Format(time.DateOnly) +
			"_" + logLevelString[InfoLevel] + ".log"
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(file))
		l.info = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.infoStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		//error
		filename = l.dir + "/" + name + "_" +
			time.Now().Format(time.DateOnly) +
			"_" + logLevelString[ErrorLevel] + ".log"
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(file))
		l.error = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.errorStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		//dpanic
		filename = l.dir + "/" + name + "_" +
			time.Now().Format(time.DateOnly) +
			"_" + logLevelString[DPanicLevel] + ".log"
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(file))
		l.dpanic = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.dpanicStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		return nil
	case outputTerminal:
		syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
		//debug
		l.debug = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.debugStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		//info
		l.info = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.infoStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		//error
		l.error = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.errorStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		//dpanic
		l.dpanic = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		l.dpanicStack = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),                       //打印行号
			zap.AddCallerSkip(1),                  //向上跳一层
			zap.AddStacktrace(zapcore.DebugLevel), //打印堆栈
		).Sugar()
		return nil
	default:
		return errors.New("unknown log output target... ")
	}
}

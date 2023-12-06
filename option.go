package logs

type Options interface {
	apply(*logger)
}

type outputFunction func(*logger)

func (o outputFunction) apply(l *logger) {
	o(l)
}

type LogLevel int8

// 日志等级
const (
	DebugLevel LogLevel = iota - 1
	InfoLevel
	ErrorLevel
	DPanicLevel
)

var logLevelString = map[LogLevel]string{
	DebugLevel:  "debug",
	InfoLevel:   "info",
	ErrorLevel:  "error",
	DPanicLevel: "panic",
}

type LogOutputTarget int8

// 日志输出目标
const (
	outputTerminal LogOutputTarget = iota //输出到终端
	outputFile                            //输出到文件
)

// LogOutputToFile 日志输出到文件
func LogOutputToFile(dir string) Options {
	return outputFunction(func(l *logger) {
		l.target = outputFile
		l.dir = dir
	})
}

// LogOutputToTerminal 日志输出到终端
func LogOutputToTerminal() Options {
	return outputFunction(func(l *logger) {
		l.target = outputTerminal
	})
}

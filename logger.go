package logs

import (
	"log"
	"time"
)

func (l *logger) Debug(msg ...any) {
	if l.level > DebugLevel {
		return
	}
	l.check()
	l.debug.Debug(msg...)
}
func (l *logger) Debugf(msg string, args ...any) {
	if l.level > DebugLevel {
		return
	}
	l.check()
	l.debug.Debugf(msg, args...)
}
func (l *logger) StackDebug(msg ...any) {
	if l.level > DebugLevel {
		return
	}
	l.check()
	l.debugStack.Debug(msg...)
}
func (l *logger) StackDebugf(msg string, args ...any) {
	if l.level > DebugLevel {
		return
	}
	l.check()
	l.debugStack.Debugf(msg, args...)
}
func (l *logger) Info(msg ...any) {
	if l.level > InfoLevel {
		return
	}
	l.check()
	l.info.Info(msg...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= DebugLevel {
		l.debug.Info(msg...)
	}
}
func (l *logger) Infof(msg string, args ...any) {
	if l.level > InfoLevel {
		return
	}
	l.check()
	l.info.Infof(msg, args...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= DebugLevel {
		l.debug.Infof(msg, args...)
	}
}
func (l *logger) StackInfo(msg ...any) {
	if l.level > InfoLevel {
		return
	}
	l.check()
	l.infoStack.Info(msg...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= DebugLevel {
		l.debugStack.Info(msg...)
	}
}
func (l *logger) StackInfof(msg string, args ...any) {
	if l.level > InfoLevel {
		return
	}
	l.check()
	l.infoStack.Infof(msg, args...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= DebugLevel {
		l.debugStack.Infof(msg, args...)
	}
}
func (l *logger) Error(msg ...any) {
	if l.level > ErrorLevel {
		return
	}
	l.check()
	l.error.Error(msg...)
	if l.target == outputTerminal {
		return
	}
	if l.level >= InfoLevel {
		l.info.Error(msg...)
	}
	if l.level <= DebugLevel {
		l.debug.Error(msg...)
	}
}
func (l *logger) Errorf(msg string, args ...any) {
	if l.level > ErrorLevel {
		return
	}
	l.check()
	l.error.Errorf(msg, args...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= InfoLevel {
		l.info.Errorf(msg, args...)
	}
	if l.level <= DebugLevel {
		l.debug.Errorf(msg, args...)
	}
}
func (l *logger) StackError(msg ...any) {
	if l.level > ErrorLevel {
		return
	}
	l.check()
	l.errorStack.Error(msg...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= InfoLevel {
		l.infoStack.Error(msg...)
	}
	if l.level <= DebugLevel {
		l.debugStack.Error(msg...)
	}
}
func (l *logger) StackErrorf(msg string, args ...any) {
	if l.level > ErrorLevel {
		return
	}
	l.check()
	l.errorStack.Errorf(msg, args...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= InfoLevel {
		l.infoStack.Errorf(msg, args...)
	}
	if l.level <= DebugLevel {
		l.debugStack.Errorf(msg, args...)
	}
}
func (l *logger) DPanic(msg ...any) {
	if l.level > DPanicLevel {
		return
	}
	l.check()
	l.dpanic.DPanic(msg...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= ErrorLevel {
		l.error.DPanic(msg...)
	}
	if l.level <= InfoLevel {
		l.info.DPanic(msg...)
	}
	if l.level <= DebugLevel {
		l.debug.DPanic(msg...)
	}
}
func (l *logger) DPanicf(msg string, args ...any) {
	if l.level > DPanicLevel {
		return
	}
	l.check()
	l.dpanic.DPanicf(msg, args...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= ErrorLevel {
		l.error.DPanicf(msg, args...)
	}
	if l.level <= InfoLevel {
		l.info.DPanicf(msg, args...)
	}
	if l.level <= DebugLevel {
		l.debug.DPanicf(msg, args...)
	}
}
func (l *logger) StackDPanic(msg ...any) {
	if l.level > DPanicLevel {
		return
	}
	l.check()
	l.dpanicStack.DPanic(msg...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= ErrorLevel {
		l.errorStack.DPanic(msg...)
	}
	if l.level <= InfoLevel {
		l.infoStack.DPanic(msg...)
	}
	if l.level <= DebugLevel {
		l.debugStack.DPanic(msg...)
	}
}
func (l *logger) StackDPanicf(msg string, args ...any) {
	if l.level > DPanicLevel {
		return
	}
	l.check()
	l.dpanicStack.DPanicf(msg, args...)
	if l.target == outputTerminal {
		return
	}
	if l.level <= ErrorLevel {
		l.errorStack.DPanicf(msg, args...)
	}
	if l.level <= InfoLevel {
		l.infoStack.DPanicf(msg, args...)
	}
	if l.level <= DebugLevel {
		l.debugStack.DPanicf(msg, args...)
	}
}

var last bool

// 校验
func (l *logger) check() {
	l.channel <- struct{}{}
	if l.checkOption() {
		err := l.initLogger()
		if err != nil && !last {
			log.Println(err)
			last = true
		}
	}
	<-l.channel
}

// 校验是否要重新创建日志记录
func (l *logger) checkOption() bool {
	if l.target == outputTerminal {
		return false
	}
	if l.value != time.Now().Day() {
		return true
	}
	return false
}

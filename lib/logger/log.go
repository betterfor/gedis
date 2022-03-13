package logger

import (
	"fmt"
	"github.com/betterfor/gedis/lib/file"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// logLevel log level
type logLevel int

// log levels
const (
	_ logLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	defaultLogger      *logger
	defaultPrefix      string
	flags              = log.LstdFlags
	level              = []string{"", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	defaultCallerDepth = 2
)

// Logger 日志接口
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

func init() {
	defaultLogger = &logger{logger: log.New(os.Stdout, defaultPrefix, flags), Locker: &sync.Mutex{}, level: DEBUG}
}

type logger struct {
	logger *log.Logger
	sync.Locker
	level logLevel
}

var _ Logger = (*logger)(nil)

func NewLogger(opts ...Option) *logger {
	opt := &option{
		path:     "",
		name:     "gedis",
		ext:      "log",
		logLevel: DEBUG,
	}

	for _, o := range opts {
		o(opt)
	}

	// 创建目录
	f, err := file.MustOpen(opt.path, fmt.Sprintf("%s.%s", opt.name, opt.ext))
	if err != nil {
		log.Fatal("logger failed: ", err)
	}

	return &logger{
		logger: log.New(io.MultiWriter(os.Stdout, f), defaultPrefix, flags),
		Locker: &sync.Mutex{},
		level:  opt.logLevel}
}

// Debug 打印 debug 日志
func Debug(v ...interface{}) { defaultLogger.Debug(v...) }
func (l *logger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.Lock()
		defer l.Unlock()
		l.setPrefix(DEBUG)
		l.logger.Println(v...)
	}
}

// Info 打印 info 日志
func Info(v ...interface{}) { defaultLogger.Info(v...) }
func (l *logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.Lock()
		defer l.Unlock()
		l.setPrefix(INFO)
		l.logger.Println(v...)
	}
}

// Warn 打印 warn 日志
func Warn(v ...interface{}) { defaultLogger.Warn(v...) }
func (l *logger) Warn(v ...interface{}) {
	if l.level <= WARNING {
		l.Lock()
		defer l.Unlock()
		l.setPrefix(WARNING)
		l.logger.Println(v...)
	}
}

// Error 打印 error 日志
func Error(v ...interface{}) { defaultLogger.Error(v...) }
func (l *logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.Lock()
		defer l.Unlock()
		l.setPrefix(ERROR)
		l.logger.Println(v...)
	}
}

// Fatal 打印 fatal 日志
func Fatal(v ...interface{}) { defaultLogger.Fatal(v...) }
func (l *logger) Fatal(v ...interface{}) {
	if l.level <= FATAL {
		l.Lock()
		defer l.Unlock()
		l.setPrefix(FATAL)
		l.logger.Fatalln(v...)
	}
}

func (l *logger) setPrefix(lvl logLevel) {
	var logPrefix string
	_, f, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", level[lvl], filepath.Base(f), line)
	} else {
		logPrefix = level[lvl]
	}
	l.logger.SetPrefix(logPrefix)
}

package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARMING
	ERROR
	UNKNOWN
)

type logger struct {
	level       LogLevel
	direct      bool
	targetFile  string
	handle      *os.File
	maxFileSize int64
}

func parseLog(s string) (logLevel LogLevel, err error) {
	switch strings.ToLower(s) {
	case "debug":
		logLevel = DEBUG
	case "trace":
		logLevel = TRACE
	case "info":
		logLevel = INFO
	case "warming":
		logLevel = WARMING
	case "error":
		logLevel = ERROR
	default:
		err = errors.New("不存在的日志等级")
		logLevel = UNKNOWN
	}
	return
}

func getInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller fail")
	}

	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file)
	return

}

func (l *logger) rotateFile() {
	stat, err := l.handle.Stat()
	if err != nil {
		panic(err)
		return
	}

	if stat.Size() >= l.maxFileSize {
		//关闭当前日志文件
		err := l.handle.Close()
		if err != nil {
			panic(err)
			return
		}
		//备份一下 rename
		err = os.Rename(l.targetFile, l.targetFile+time.Now().Format("20060102150405000"))
		if err != nil {
			panic(err)
			return
		}
		//打开一个新的文件
		handle, err := os.OpenFile(l.targetFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
			return
		}
		//将打开的文件赋值回 注意logger 要引用
		l.handle = handle
	}
}

func NewLogger(path string, level string) *logger {
	var handle *os.File
	var err error
	var direct = true
	if path != "" {
		handle, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("日志文件创建失败:" + fmt.Sprintf("%v", err)) //程序崩溃退出
		}
		direct = false
	}
	pLevel, err := parseLog(level)
	if err != nil {
		panic("日志文件创建失败:" + fmt.Sprintf("%v", err)) //程序崩溃退出
	}
	return &logger{
		level:       pLevel,
		direct:      direct,
		targetFile:  path,
		handle:      handle,
		maxFileSize: 2 * 1024 * 1024,
	}

}

func (l *logger) Close() {
	//fmt.Println("进来关掉了文件")
	l.handle.Close()
}

func (l *logger) doPrint(s string, logLevel LogLevel) {
	if l.level < logLevel {
		return
	}
	//检查文件大小 做切割
	l.rotateFile()

	funName, fileName, line := getInfo(3)
	//Mon Jan 2 15:04:05 -0700 MST 2006
	formatStr := fmt.Sprintf("[%s][%s:%s:%d]", time.Now().Format("2006-06-02 15:04:05.0000"), fileName, funName, line)
	s = formatStr + s + "\n"
	if l.direct == true {
		fmt.Fprintln(os.Stdout, s)
	} else {
		l.handle.WriteString(s)
	}
}

func init() {
	//fmt.Println("init")
}

func (l *logger) Debug(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	content := "[DEBUG]" + str
	l.doPrint(content, DEBUG)

}

func (l *logger) Trace(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	content := "[TRACE]" + str
	l.doPrint(content, TRACE)
}

func (l *logger) Info(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	content := "[INFO]" + str
	l.doPrint(content, INFO)

}

func (l *logger) Warming(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	content := "[WARMING]" + str
	l.doPrint(content, WARMING)

}

func (l *logger) Error(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	content := "[ERROR]" + str
	l.doPrint(content, ERROR)

}

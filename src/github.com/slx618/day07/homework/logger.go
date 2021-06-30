package main

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
	content     chan *logMsg
}

type logMsg struct {
	line       int
	fileName   string
	funcName   string
	timesStamp string
	msg        string
	level      LogLevel
}

type Log interface {
	Debug(string, ...interface{})
	Trace(string, ...interface{})
	Info(string, ...interface{})
	Warming(string, ...interface{})
	Error(string, ...interface{})
	Close()
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

func parseLogLevel(le LogLevel) (logLevel string, err error) {
	switch le {
	case DEBUG:
		logLevel = "DEBUG"
	case TRACE:
		logLevel = "TRACE"
	case INFO:
		logLevel = "INFO"
	case WARMING:
		logLevel = "WARMING"
	case ERROR:
		logLevel = "ERROR"
	default:
		err = errors.New("不存在的日志等级")
		logLevel = "UNKNOWN"
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

func (l *logger) writeFile() {
	if len(l.content) > 500 {
		var begin = 1
		for {
			l.handle.WriteString(l.formartLogStr())
			begin++
			if begin == 500 {
				break
			}
		}
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

	content := make(chan *logMsg, 50000)
	return &logger{
		level:       pLevel,
		direct:      direct,
		targetFile:  path,
		handle:      handle,
		maxFileSize: 2 * 1024 * 1024,
		content:     content,
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
	funcName, fileName, line := getInfo(3)

	msg := logMsg{
		line:     line,
		fileName: fileName,
		funcName: funcName,
		//Mon Jan 2 15:04:05 -0700 MST 2006
		timesStamp: time.Now().Format("2006-06-02 15:04:05.0000"),
		msg:        s,
		level:      DEBUG,
	}

	select {
	case l.content <- &msg:
	default:
	}

	//检查文件大小 做切割
	l.rotateFile()

	if l.direct == true {
		fmt.Fprintln(os.Stdout, l.formartLogStr())
	} else {
		go l.writeFile()
	}
}

func (l *logger) formartLogStr() string {
	logMsg := <-l.content
	return fmt.Sprintf("[%s][%s:%s.%s:%d][%s]%s\n", logMsg.timesStamp, logMsg.fileName, logMsg.funcName,
		logMsg.funcName, logMsg.line, logMsg.level, logMsg.msg)
}

func init() {
	//fmt.Println("init")
}

func (l *logger) Debug(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	l.doPrint(str, DEBUG)

}

func (l *logger) Trace(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	l.doPrint(str, TRACE)
}

func (l *logger) Info(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	l.doPrint(str, INFO)

}

func (l *logger) Warming(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	l.doPrint(str, WARMING)

}

func (l *logger) Error(s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	l.doPrint(str, ERROR)

}

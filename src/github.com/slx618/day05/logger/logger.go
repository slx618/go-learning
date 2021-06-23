package logger

import (
	"fmt"
	"os"
	"time"
)

type logger struct {
	direct     bool
	targetFile string
	handle     *os.File
}

func NewLogger(path string) *logger {
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
	return &logger{
		direct:     direct,
		targetFile: path,
		handle:     handle,
	}

}

func (l logger) doPrint(s string) {
	//检查文件大小 做切割

	//Mon Jan 2 15:04:05 -0700 MST 2006
	s = "[" + time.Now().Format("2006-06-02 15:04:05") + "]" + s + "\n"
	if l.direct == true {
		fmt.Fprintln(os.Stdout, s)
	} else {
		l.handle.WriteString(s)
	}
}

func init() {
	fmt.Println("init")
}

func (l logger) Debug(s interface{}) {
	var str = fmt.Sprintf("%#v", s)
	content := "[DEBUG]" + str
	l.doPrint(content)

}

func (l logger) Trace(s interface{}) {
	var str = fmt.Sprintf("%#v", s)
	content := "[TRACE]" + str
	l.doPrint(content)
}

func (l logger) Fatal(s interface{}) {
	var str = fmt.Sprintf("%#v", s)
	content := "[ERROR]" + str
	l.doPrint(content)

}

func (l logger) Info(s interface{}) {
	var str = fmt.Sprintf("%#v", s)
	content := "[INFO]" + str
	l.doPrint(content)

}

func (l logger) Warming(s interface{}) {
	var str = fmt.Sprintf("%#v", s)
	content := "[WARMING]" + str
	l.doPrint(content)

}

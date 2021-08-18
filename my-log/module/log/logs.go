package log

import (
	"bytes"
	"fmt"
	. "my-log/config"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	INFO = "INFO"

	DEBUG = "DEBUG"

	ERROR = "ERROR"

	WARN = "WARN"
)

var openConsole = true

func init() {

	openConsole = Conf.LogConfig.Console

}

var lock sync.RWMutex

// New 创建一条日志
func New(format string, a ...interface{}) {
	lock.Lock()
	defer lock.Unlock()
	fmtLog := getHEAD(Conf.LogConfig.LogLevel, format, a)
	outConsole(fmtLog, Conf.LogConfig.LogLevel)
	WriteLog(fmtLog)
}

// LogInfo INFO类型日志
func LogInfo(format string, a ...interface{}) {

	lock.Lock()
	defer lock.Unlock()
	fmtLog := getHEAD(INFO, format, a)
	outConsole(fmtLog, INFO)
	WriteLog(fmtLog)

}

// LogDEBUG LogDEBUG类型日志
func LogDEBUG(format string, a ...interface{}) {

	lock.Lock()
	defer lock.Unlock()
	fmtLog := getHEAD(DEBUG, format, a)
	outConsole(fmtLog, DEBUG)
	WriteLog(fmtLog)

}

// LogERROR LogERROR类型日志
func LogERROR(format string, a ...interface{}) {

	lock.Lock()
	defer lock.Unlock()
	fmtLog := getHEAD(ERROR, format, a)
	outConsole(fmtLog, ERROR)
	WriteLog(fmtLog)

}

// LogWARN LogERROR类型日志
func LogWARN(format string, a ...interface{}) {

	lock.Lock()
	defer lock.Unlock()
	fmtLog := getHEAD(WARN, format, a)
	outConsole(fmtLog, WARN)
	WriteLog(fmtLog)

}

// outConsole 控制台打印
func outConsole(msg string, t string) {

	if openConsole {
		switch t {
		case ERROR:
			fmt.Errorf(msg)
		default:
			fmt.Println(msg)
		}
	}
}

// fmtString 格式化字符
func fmtString(msg string) string {

	if len(msg) == 0 {
		return ""
	}
	return string([]byte(msg)[1 : len(msg)-1])
}

// getHEAD 标识拼写
func getHEAD(str string, msg string, a interface{}) string {

	msg = fmtString(fmt.Sprintf(msg, a))
	pc, file, line, ok := runtime.Caller(2)

	var bt bytes.Buffer
	bt.WriteString(time.Now().Format(Conf.LogConfig.OutputFormat))
	bt.WriteString("  [")
	bt.WriteString(str)
	if ok {
		bt.WriteString(" ,file: ")
		bt.WriteString(file)
		bt.WriteString(" ,FuncName: ")
		bt.WriteString(runtime.FuncForPC(pc).Name())
		bt.WriteString(" ,line: ")
		bt.WriteString(strconv.Itoa(line))
	}

	bt.WriteString("] --> ")
	bt.WriteString(msg)

	return bt.String()
}

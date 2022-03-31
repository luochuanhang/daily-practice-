package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

//设置前缀
func setPrefix(level Level) {
	/*
		调用者报告关于调用goroutine堆栈上的函数调用的文件和行号信息。参数跳过是要上升的堆栈帧数，
		其中e表示caller的调用方。(由于历史原因，skip的含义在Caller和Callers之间是不同的。)
		返回值报告相应调用文件中的程序计数器、文件名和行号。如果不能恢复信息，则布尔ok为false。
	*/
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		//如果可恢复前缀为对应信息类型+返回路劲最近一个元素+行号
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		//如果不可恢复前缀为错误信息
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	//设置log前缀
	logger.SetPrefix(logPrefix)
}

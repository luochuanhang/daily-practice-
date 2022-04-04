package logging

import (
	"fmt"
	"lianxi/blog/pkg/file"
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

// 安装程序初始化日志实例
func Setup() {
	var err error
	//log日志保存的地址
	filePath := getLogFilePath()
	//login日志文件名
	fileName := getLogFileName()
	//如果不存在就创建一个文件
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	/*
		New创建一个新的日志记录器。out变量设置要写入日志数据的目标。
		前缀出现在每个生成的日志行的开头，或者如果提供了Lmsgprefix标志，
		则出现在日志头的后面。flag参数定义了日志记录属性。
	*/
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug 输出调试级别的日志
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info 输出info级别的日志
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn 输出警告级别的日志
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error输出错误级别的日志
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal输出致命级别的错误
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

// setPrefix 设置日志输出信息的前缀
func setPrefix(level Level) {
	/*
		调用者报告关于调用goroutine堆栈上的函数调用的文件和行号信息。
		参数跳过是要上升的堆栈帧数，O表示caller的调用方。
		(由于历史原因，skip的含义在Caller和Callers之间是不同的。)
		返回值报告相应调用文件中的程序计数器、文件名和行号。
		如果不能恢复信息，则布尔ok为false。
	*/
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	//如果不能恢复信息，只打印错误类型，能恢复信息打印文件名和行号
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile) //
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
	//"\033[31m[error]\033[0m "\033[31m: 设置文本颜色为红色。
	//[error]: 显示的文本，标识为错误消息。
	//\033[0m: 重置文本样式，确保后续文本不受之前的样式影响。
	//下面的 同理
	//
	//log.new 中的flag控制日志记录的格式选项
	//log.LstdFlags表示选用标准的时间与日期
	//log.Lshortlife 表示日志中显示文件名与行号
)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)

	}
	if level > ErrorLevel {
		errorLog.SetOutput(io.Discard)
	}
	if level > InfoLevel {
		infoLog.SetOutput(io.Discard)
	}
}

// log methods
var (
	Error = errorLog.Println //以一行的形式记录错误信息

	Errorf = errorLog.Printf //格式化错误信息

	Info  = infoLog.Println
	Infof = infoLog.Printf
)

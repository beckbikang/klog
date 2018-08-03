package klog

//      Trace，Info,Warn,Error,Fatal
//定义基本的常量
type LOG_LEVEL int

const (
	TRACE LOG_LEVEL = iota
	INFO
	WARN
	ERROR
	FATAL
)
var(

	logTagName = map[LOG_LEVEL]string{
		TRACE:"TRACE",
		INFO:"INFO",
		WARN:"WARN",
		ERROR:"ERROR",
		FATAL:"FATAL",
	}
)

//检查level
func checkLevel(level LOG_LEVEL )bool{
	if level >= TRACE && level <=FATAL{
		return true
	}
	return false
}

//定义日志回调函数

type callLogFunc func (string)

type LOG_MSG struct{
	level LOG_LEVEL
	data string
	f callLogFunc
}

//日志通用功能接口
type KLogger interface {
	Level() LOG_LEVEL
	NewLog(interface{})error
	GetMsgChan(chan error)chan *LOG_MSG
	Start()
	Close()
	write(chan *LOG_MSG)
}

//类的通用的注册管理器
var LogMap = map[string]KLogger{}

func Reg(name string, logger KLogger ){
	if logger == nil{
		panic("register empty klogger")
	}
	if LogMap[name] != nil{
		panic("logger has been register ")
	}
	LogMap[name] = logger
}








//日志的通用的属性
type LogCommon struct {
	level LOG_LEVEL
	msg chan *LOG_MSG
	quitChan chan struct{}
	errorChan chan error
}



//定义标准库输出类






//定义文件输出类














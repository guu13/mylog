package mylog

import "os"

type Mylogger struct {
	AppName string
	Path    string
	Logfile string
}

var (
	mylogger = New()
)

func New() *Mylogger {
	return &Mylogger{AppName: "mylog", Path: "/var/log/", Logfile: "/var/log/mylog.log"}
}

func SetName(AppName string, LogName string) {
	mylogger.AppName = AppName
	mylogger.Logfile = mylogger.Path + LogName + ".log"
}

func SetAppName(AppName string) {
	mylogger.AppName = AppName
	mylogger.Logfile = mylogger.Path + AppName + ".log"
}

func Log2file(log string) {

	//创建日志文件
	f, err := os.OpenFile(mylogger.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	// 关闭文件
	defer f.Close()

	// 字符串写入
	_, err = f.WriteString(mylogger.AppName + ":::" + log + "\n")
	if err != nil {
		return
	}
}

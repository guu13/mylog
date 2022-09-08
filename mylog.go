package mylog

import "os"

type Mylogger struct {
	Name string
	Path string
}

var (
	mylogger = New()
)

func New() *Mylogger {
	return &Mylogger{Name: "mylog", Path: "/var/log/"}
}

func Mylog(Name string) {
	mylogger.Path += Name
	mylogger.Name = Name
}

func Log2file(log string) {

	//创建日志文件
	f, err := os.OpenFile(mylogger.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	// 关闭文件
	defer f.Close()

	// 字符串写入
	_, err = f.WriteString(mylogger.Name + " " + log + "\n")
	if err != nil {
		return
	}
}

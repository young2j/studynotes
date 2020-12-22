package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	// 设置日志格式
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	// 设置日志前缀
	log.SetPrefix("[日志前缀]")
	// 设置日志的输出目标
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile)
}

func main() {
	log.Println("这是一普通日志")
	v := "普通"
	log.Printf("这也是一条%s的日志\n", v)
	log.Fatalln("这是一条Fatal日志，会调用os.Exit(1)")
	log.Panicln("这是一条Panic日志，会触发panic")
}

/*
日志的设置flag
const(
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)

自定义log
func New(out io.Writer, prefix string, flag int) *Logger
logger:=log.New(...)
logger.Println("...")
*/

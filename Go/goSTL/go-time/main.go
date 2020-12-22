package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间类型-年月日时分秒
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("当前时间：%v\n", now)
	fmt.Printf("格式化时间：%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	nowAdd1h := now.Add(time.Hour)
	fmt.Printf("当前时间+1小时:%v\n", nowAdd1h)
	fmt.Printf("+1后时间间隔(Duration):%v\n", now.Sub(nowAdd1h))
	fmt.Printf("当前时间==+1h后的时间：%t\n", now.Equal(nowAdd1h)) // Equal会比较时区和地点信息，不同于==，尽量用Equal
	fmt.Printf("当前时间 Before +1h后的时间：%t\n", now.Before(nowAdd1h))
	fmt.Printf("当前时间 After +1h后的时间：%t\n\n", now.After(nowAdd1h))

	// 时间戳
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano() //纳秒时间戳
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)
	// 时间戳转时间格式
	unixTimeStamp := time.Unix(timeStamp1, 0)
	year = unixTimeStamp.Year()
	month = unixTimeStamp.Month()
	day = unixTimeStamp.Day()
	hour = unixTimeStamp.Hour()
	minute = unixTimeStamp.Minute()
	second = unixTimeStamp.Second()
	fmt.Printf("当前时间：%v\n", unixTimeStamp)
	fmt.Printf("格式化时间：%d-%02d-%02d %02d:%02d:%02d\n\n", year, month, day, hour, minute, second)

	// 定时器---定时器的本质是一个通道（channel）
	// ticker := time.Tick(time.Second)
	// for i := range ticker {
	// 	fmt.Println("计时：", i)
	// }

	// 时间格式化---Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）
	now = time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// 字符串时间格式解析
	loc, err := time.LoadLocation("Asia/Shanghai") //加载时区
	if err != nil {
		fmt.Println(err)
		return
	}
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc) //按照指定时区和格式解析 "2019/08/04 14:15:20"
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}

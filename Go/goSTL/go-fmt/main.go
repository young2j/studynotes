package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print 系列
	fmt.Print("直接打印无换行  ")
	str := string("字符串")
	fmt.Printf("格式化输出%s\n", str)
	fmt.Println("直接打印末尾带有换行")

	// Fprint 系列——向实现了io.Writer的对象写入内容
	fmt.Fprintf(os.Stdout, "打印，同时向标准输出写入\n")
	file, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	fmt.Fprintf(file, "向文件写入【%s】，不会在控制台打印", str)

	// Sprint 系列——将传入的数据生成返回一个字符串
	str = fmt.Sprint("打印，同时返回一个字符串\n")
	spt := fmt.Sprintf("打印，同时返回一个格式化%s\n", "字符串")
	fmt.Print(str, spt)

	// Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	err = fmt.Errorf("这是一个错误")

	// 通用格式化占位符
	o := struct{ name string }{"gopher"}
	fmt.Printf("%v\n", o)  //值，不带字段名
	fmt.Printf("%+v\n", o) // 值，带字段名
	fmt.Printf("%#v\n", o) // 类型定义+值
	fmt.Printf("%T\n", o)  // 类型定义
	fmt.Printf("100%%\n")  // 百分号

	// bool型格式化占位符
	fmt.Printf("%t\n", false)

	// string 和 []byte型格式化占位符
	fmt.Printf("%s\n", "直接输出字符串")
	fmt.Printf("%q\n", "输出双引号括起来的字符串")

	// 指针格式化占位符
	pstr := "输出指针地址，并带前导ox"
	fmt.Printf("%p\n", &pstr)

	// 宽度、精度占位符
	num := 66.66
	fmt.Printf("%f\n", num)     //默认
	fmt.Printf("%9f\n", num)    // 宽度9
	fmt.Printf("%.2f\n", num)   // 精度2
	fmt.Printf("%9.f\n", num)   // 宽度9，精度0,四舍五入，右对齐
	fmt.Printf("%-9.f\n", num)  // 宽度9，精度0,四舍五入，左对齐
	fmt.Printf("%+9.f\n", num)  // 宽度9，精度0,四舍五入，带正负号
	fmt.Printf("%09.2f\n", num) // 宽度9，精度2，根据宽度正负号后补0

	//获取输入
	// fmt.Scan()---从标准输入获取由空白符分隔的值，换行符也视为空白符
	// fmt.Scanln()---从标准输入获取由空白符分隔的值，遇到换行符结束扫描
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	// fmt.Scanln(&name, &age, &married)
	fmt.Printf("name:%s age:%d married:%t", name, age, married)

	// fmt.Scanf()---从标准输入获取指定格式的值
	var (
		height int
		weight int
		gender bool
	)
	fmt.Scanf("1:%d 2:%d 3:%t", &height, &weight, &gender)
	fmt.Printf("扫描结果 height:%d weight:%d gender:%t \n", height, weight, gender)

	// 获取完整输入
	// bufio.NewReader
	reader := bufio.NewReader(os.Stdin) //创建一个标准输入读取器
	fmt.Println("请输入内容:")
	input, err := reader.ReadString('\n') // 读取到换行符停止
	input = strings.TrimSpace(input)
	fmt.Printf("%#v\n", input)

	// 从io.Reader读取数据
	// func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	// func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	// func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

	// 从指定字符串读取数据
	// func Sscan(str string, a ...interface{}) (n int, err error)
	// func Sscanln(str string, a ...interface{}) (n int, err error)
	// func Sscanf(str string, format string, a ...interface{}) (n int, err error)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//标准io
	// 终端本质是一个文件
	var buf [16]byte
	os.Stdin.Read(buf[:])                 // 从标准输入读取，存放在buf中
	os.Stdout.WriteString(string(buf[:])) //将buf又写回标准输出

	// 文件io
	// 创建文件
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("创建练习用文件\n")
	file.Close()

	// 打开文件
	file, err = os.Open("./test.txt") //只读方式打开文件,自定义权限使用os.OpenFile()
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取文件
	contentBuf := make([]byte, 128)
	var content []byte
	for {

		n, err := file.Read(contentBuf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		content = append(content, contentBuf[:n]...)
	}
	fmt.Println(string(content))

	//bufio包实现了带缓冲区的读写，是对文件读写的封装
	//写入
	file, err = os.OpenFile("./bufio.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		line := fmt.Sprintf("通过bufio写入的内容%d\n", i)
		writer.WriteString(line)
	}
	writer.Flush()
	file.Close()

	// 读取
	file, err = os.Open("./bufio.txt")
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
	file.Close()

	// ioutil-更简便的文件读写
	err = ioutil.WriteFile("./ioutil.txt", []byte("ioutil写入的内容"), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err = ioutil.ReadFile("./ioutil.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

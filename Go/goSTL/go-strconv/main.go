package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string to int
	strNumber := "12345"
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		fmt.Println("can't convert string to int,err:", err)
	} else {
		fmt.Printf("type:%T value:%#v\n", number, number)
	}

	// int to string
	strNumber = strconv.Itoa(number)
	fmt.Printf("type:%T value:%#v\n", strNumber, strNumber)

	// parse string to others
	b, err := strconv.ParseBool("true")
	fmt.Printf("type:%T value:%#v\n", b, b)
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("type:%T value:%#v\n", f, f)
	i, err := strconv.ParseInt("-2", 10, 64)
	fmt.Printf("type:%T value:%#v\n", i, i)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Printf("type:%T value:%#v\n", u, u)

	// format others to string(一切皆可string，所以无err返回)

	strB := strconv.FormatBool(b)
	fmt.Printf("type:%T value:%#v\n", strB, strB)
	strI := strconv.FormatInt(i, 10)
	fmt.Printf("type:%T value:%#v\n", strI, strI)
	strU := strconv.FormatUint(u, 10)
	fmt.Printf("type:%T value:%#v\n", strU, strU)
	strF := strconv.FormatFloat(f, 'E', -1, 64)
	fmt.Printf("type:%T value:%#v\n", strF, strF) // Formatfloat()
	// fmt表示格式：
	// 	’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
	// prec控制精度（排除指数部分）：
	//	-f,e,E: 表示小数点后的数字个数；
	// 	-g,G: 控制总的数字个数。
	// 	-1：则代表使用最少数量的、但又必需的数字来表示f。
	// bitSize：表示f的来源类型（32：float32、64：float64），会据此进行舍入。

	// Append系列
	// Quote系列
}

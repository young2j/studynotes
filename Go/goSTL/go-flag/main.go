package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//os.Args []string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("arg[%d]=%v\n", index, arg)
		}
	}

	// flag.Type()
	name := flag.String("name", "张三", "名称") // 参数名，默认值，帮助信息
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否") // bool型参数需要以=号形式传递 如 --married=false

	flag.Parse() //解析命令行参数
	fmt.Println(*name, *age, *married)
	fmt.Println(flag.Args())  //返回命令行参数切片
	fmt.Println(flag.NArg())  // 返回命令行参数个数
	fmt.Println(flag.NFlag()) // 返回使用的flag参数个数

	//flag.TypeVar()
	var (
		height int
		weight int
		gender bool
	)
	flag.IntVar(&height, "height", 180, "身高") // *type，flagName,default,help
	flag.IntVar(&weight, "weight", 180, "体重")
	flag.BoolVar(&gender, "gender", true, "性别")
	flag.Parse() //解析命令行参数
	fmt.Println(height, weight, gender)
	fmt.Println(flag.Args())  //返回命令行参数切片
	fmt.Println(flag.NArg())  // 返回命令行参数个数
	fmt.Println(flag.NFlag()) // 返回使用的flag参数个数
}

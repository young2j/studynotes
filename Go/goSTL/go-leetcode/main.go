package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)


func main() {
	s := "abcdabcd1234ABCD"
	fmt.Println(s[0:0])
	fmt.Println(strings.Index(s[0:0], "a"))
	// arr2d:=make([][]int,5)
	// fmt.Printf("%T\n",arr2d)
	
	arr2d := [][]string{{"a"," ","b"},{"2","2"," "}}
	res:=""
	for _, v := range arr2d {
		res+="\n"
		for _, e := range v {
			res+=e
		}
	}
	fmt.Println(res)

	fmt.Println(int(math.Ceil(float64(10)/float64(3))))
	fmt.Println(float64(10)/float64(3))
	a:=strconv.Itoa(234)
	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i,string(v))
	}
	fmt.Println(strconv.Atoi("0000321"))
	fmt.Println(strings.ToLower(s))
	fmt.Println(reflect.TypeOf(fmt.Sprintf("%v","1234")))
	
	fmt.Printf("%q\n",strings.Split(s,"b"))

}
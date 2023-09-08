/*
 * File: main.go
 * Created Date: 2023-03-06 02:21:50
 * Author: ysj
 * Description:  golang 原型模式客户端调用
 */

package main

func main() {
	// 二狗
	ergou := &ClayBoy{
		Name:   "二狗",
		Eye:    "黑色",
		Nose:   "两个鼻孔",
		Skin:   "黄色",
		Hands:  2,
		Gender: "男",
		Temper: "很好",
	}
	ergou.Introduction()

	// 通过二狗子克隆出大壮, 大壮脾气不好
	dazhuang := ergou.Clone()
	dazhuang.SetName("大壮")
	dazhuang.SetTemper("暴躁")
	dazhuang.Introduction()

	// 小花
	xiaohua := &ClayGirl{
		Name:   "小花",
		Eye:    "黑色",
		Nose:   "两个鼻孔",
		Skin:   "黄色",
		Hands:  2,
		Gender: "女",
		Temper: "温柔",
	}
	xiaohua.Introduction()

	// 通过小花克隆出小朵，她们是双胞胎，只有名字不一样
	xiaoduo := xiaohua.Clone()
	xiaoduo.SetName("小朵")
	xiaoduo.Introduction()
}

/*
 * File: file.go
 * Created Date: 2023-03-27 01:30:45
 * Author: ysj
 * Description: 设计模式之组合模式--文件
 */

package main

import (
	"fmt"
	"strings"
)

type File struct {
	name   string // 文件名称
	isLeaf bool   // 是否叶子节点
	depth  int    // 层次深度
}

// 初始化
func NewFile(name string) Compositer {
	return &File{
		name:   name,
		isLeaf: true,
		depth:  1,
	}
}

// 添加一个目录
func (d *File) Add(c Compositer) {
	fmt.Println("a file can not add a compositer!")
}

// 移除一个目录
func (d *File) Remove(c Compositer) {
	fmt.Println("a file can not remove a compositer!")
}

// 显示目录信息
func (d *File) Info() {
	fmt.Println(strings.Repeat("-", d.depth) + " " + d.name)
}

// 增加深度
func (d *File) addDepth(depth int) {
	d.depth = depth
}

// 目录名称
func (d *File) getName() string {
	return d.name
}

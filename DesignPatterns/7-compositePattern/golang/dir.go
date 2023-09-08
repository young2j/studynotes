/*
 * File: dir.go
 * Created Date: 2023-03-27 01:30:17
 * Author: ysj
 * Description: 设计模式之组合模式--目录
 */

package main

import (
	"fmt"
	"strings"
)

type Dir struct {
	name    string       // 目录名称
	isLeaf  bool         // 是否叶子节点
	depth   int          // 层次深度
	subDirs []Compositer // 子目录树
}

// 初始化
func NewDir(name string) Compositer {
	return &Dir{
		name:    name,
		isLeaf:  false,
		depth:   1,
		subDirs: make([]Compositer, 0),
	}
}

// 添加一个目录
func (d *Dir) Add(c Compositer) {
	d.subDirs = append(d.subDirs, c)
	c.addDepth(d.depth + 1)
}

// 移除一个目录
func (d *Dir) Remove(c Compositer) {
	for i, subdir := range d.subDirs {
		if subdir.getName() == c.getName() {
			d.subDirs = append(d.subDirs[:i], d.subDirs[i+1:]...)
			break
		}
	}
}

// 显示目录信息
func (d *Dir) Info() {
	fmt.Println(strings.Repeat("-", d.depth) + " " + d.name)
	for _, subdir := range d.subDirs {
		subdir.Info()
	}
}

// 增加深度
func (d *Dir) addDepth(depth int) {
	d.depth = depth
}

// 目录名称
func (d *Dir) getName() string {
	return d.name
}

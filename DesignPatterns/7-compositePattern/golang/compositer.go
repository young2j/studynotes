/*
 * File: composite.go
 * Created Date: 2023-03-27 01:26:11
 * Author: ysj
 * Description: 设计模式之组合模式--组合接口
 */
package main

type Compositer interface {
	Add(c Compositer)
	Remove(c Compositer)
	Info()
	addDepth(depth int)
	getName() string
}

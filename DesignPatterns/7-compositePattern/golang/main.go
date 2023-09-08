/*
 * File: main.go
 * Created Date: 2023-03-27 01:25:47
 * Author: ysj
 * Description: 组合模式客户端调用
 */

package main

import "fmt"

func main() {
	// 创建根目录，在其下创建两个目录，两个文件
	rootDir := NewDir("根目录")
	subDir1 := NewDir("子目录1")
	subDir2 := NewDir("子目录2")
	file1 := NewFile("文件1")
	file2 := NewFile("文件2")
	rootDir.Add(subDir1)
	rootDir.Add(subDir2)
	rootDir.Add(file1)
	rootDir.Add(file2)

	// 在子目录1下创建一个目录和一个文件
	dir1_1 := NewDir("子目录1-1")
	file1_1 := NewFile("文件1-1")
	subDir1.Add(dir1_1)
	subDir1.Add(file1_1)

	rootDir.Info()

	fmt.Println()
	// 删除根目录下的文件2和子目录2
	rootDir.Remove(file2)
	rootDir.Remove(subDir2)

	rootDir.Info()
}

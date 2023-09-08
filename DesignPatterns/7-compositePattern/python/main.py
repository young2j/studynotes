#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2023-03-27 03:07:20
# Author: ysj
# Description: 组合模式客户端调用
###

from dir import Dir
from file import File

# 创建根目录，在其下创建两个目录，两个文件
root_dir = Dir("根目录")
sub_dir1 = Dir("子目录1")
sub_dir2 = Dir("子目录2")
file1 = File("文件1")
file2 = File("文件2")
root_dir.add(sub_dir1)
root_dir.add(sub_dir2)
root_dir.add(file1)
root_dir.add(file2)

# 在子目录1下创建一个目录和一个文件
dir1_1 = Dir("子目录1-1")
file1_1 = File("文件1-1")
sub_dir1.add(dir1_1)
sub_dir1.add(file1_1)

root_dir.info()

print()
# 删除根目录下的文件2和子目录2
root_dir.remove(file2)
root_dir.remove(sub_dir2)

root_dir.info()

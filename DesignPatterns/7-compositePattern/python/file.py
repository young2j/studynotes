#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: file.py
# Created Date: 2023-03-27 03:07:15
# Author: ysj
# Description: 设计模式之组合模式--文件
###
from compositer import Compositer


class File(Compositer):
    def __init__(self, name):
        self.name = name
        self.is_leaf = True
        self.depth = 1

    # 添加一个目录
    def add(self, c: Compositer):
        print("a file can not add a compositer!")

    # 移除一个目录
    def remove(self, c: Compositer):
        print("a file can not remove a compositer!")

    # 显示目录信息
    def info(self):
        print("-"*self.depth + " " + self.name)

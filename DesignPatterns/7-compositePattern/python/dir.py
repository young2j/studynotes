#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: dir.py
# Created Date: 2023-03-27 03:07:11
# Author: ysj
# Description: 设计模式之组合模式--目录
###

from typing import List
from compositer import Compositer


class Dir(Compositer):
    def __init__(self, name):
        self.name = name
        self.is_leaf = False
        self.depth = 1
        self.sub_dirs: List[Compositer] = []

    # 添加一个目录
    def add(self, c: Compositer):
        self.sub_dirs.append(c)
        c.depth = self.depth + 1

    # 移除一个目录
    def remove(self, c: Compositer):
        for i, subdir in enumerate(self.sub_dirs):
            if subdir.name == c.name:
                self.sub_dirs.pop(i)

    # 显示目录信息
    def info(self):
        print("-"*self.depth + " " + self.name)
        for subdir in self.sub_dirs:
            subdir.info()

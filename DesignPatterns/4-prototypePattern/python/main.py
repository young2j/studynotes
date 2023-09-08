#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2023-03-06 03:20:27
# Author: ysj
# Description:  python 原型模式客户端调用
###

from prototype import ClayBoy, ClayGirl

# 二狗
ergou = ClayBoy(name="二狗", eye="黑色", nose="两个鼻孔",
                skin="黄色", hands=2, gender="男", temper="很好",
                )
ergou.introduction()

# 通过二狗子克隆出大壮, 大壮脾气不好
dazhuang = ergou.clone()
dazhuang.name = "大壮"
dazhuang.temper = "暴躁"
dazhuang.introduction()

# 小花
xiaohua = ClayGirl(name="小花", eye="黑色", nose="两个鼻孔",
                   skin="黄色", hands=2, gender="女", temper="温柔",
                   )
xiaohua.introduction()

# 通过小花克隆出小朵，她们是双胞胎，只有名字不一样
xiaoduo = xiaohua.clone()
xiaoduo.name = "小朵"
xiaoduo.introduction()

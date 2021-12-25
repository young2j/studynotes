#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: context.py
# Created Date: 2021-12-19 06:12:20
# Author: ysj
# Description:  状态上下文类
###


class OrderContext(object):
    """订单状态上下文"""

    def __init__(self, state):
        """初始化"""
        super().__init__()
        self.__state = state
        self.has_paid = False
        self.elapsed_30min = False

    def set_state(self, state):
        """设置当前状态"""
        self.__state = state

    def get_state(self):
        """获取当前状态"""
        return self.__state

    def handle(self):
        """当前状态下的行为"""
        self.__state.handle(self)

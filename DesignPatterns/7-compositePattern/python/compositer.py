#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: compositer.py
# Created Date: 2023-03-27 03:07:02
# Author: ysj
# Description:  设计模式之组合模式--组合接口
###
from __future__ import annotations
# from typing import Self # python3.11
from abc import ABCMeta, abstractmethod


class Compositer(metaclass=ABCMeta):
    @abstractmethod
    def add(self, c: Compositer):
        pass

    @abstractmethod
    def remove(self, c: Compositer):
        pass

    @abstractmethod
    def info(self):
        pass

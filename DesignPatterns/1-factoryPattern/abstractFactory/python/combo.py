#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: combo.py
# Created Date: 2021-11-29 06:18:05
# Author: ysj
# Description:  套餐类
###

from abc import ABCMeta, abstractmethod


class Combo(metaclass=ABCMeta):
    """抽象基类"""

    name = "套餐名"
    price = 0

    @abstractmethod
    def task(self):
        pass

    @property
    @abstractmethod
    def name(self):
        return self.name

    @property
    @abstractmethod
    def price(self):
        return self.price


class TrialCombo(Combo):
    """具体实现类"""

    name = "体验版"
    price = 299

    def task(self):
        print(f"{self.name}-> 执行价值{self.price}元的服务")
        return


class BasicCombo(Combo):
    """具体实现类"""

    name = "基础版"
    price = 599

    def task(self):
        print(f"{self.name}-> 执行价值{self.price}元的服务")
        return


class PremiumCombo(Combo):
    """具体实现类"""

    name = "高级版"
    price = 1999

    def task(self):
        print(f"{self.name}-> 执行价值{self.price}元的服务")
        return

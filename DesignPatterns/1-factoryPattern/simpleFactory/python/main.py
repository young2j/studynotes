#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2021-11-24 12:11:07
# Author: ysj
# Description: python 简单工厂
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


class SimpleFactory(object):
    """简单工厂类"""

    def create_combo(self, combo_type=None):
        if combo_type == "trial":
            return TrialCombo()
        elif combo_type == "basic":
            return BasicCombo()
        elif combo_type == "premium":
            return PremiumCombo()
        else:
            return TrialCombo()


if __name__ == "__main__":
    factory = SimpleFactory()

    trial_combo = factory.create_combo('trial')
    trial_combo.task()

    basic_combo = factory.create_combo('basic')
    basic_combo.task()

    premium_combo = factory.create_combo('premium')
    premium_combo.task()

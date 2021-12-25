#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2021-11-26 12:13:10
# Author: ysj
# Description:  python 工厂方法
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


class Factory(metaclass=ABCMeta):
    """抽象工厂基类"""
    @abstractmethod
    def create_combo(self):
        pass


class TrialComboFactory(Factory):
    """具体工厂类"""

    def create_combo(self):
        return TrialCombo()


class BasicComboFactory(Factory):
    """具体工厂类"""

    def create_combo(self):
        return BasicCombo()


class PremiumComboFactory(Factory):
    """具体工厂类"""

    def create_combo(self):
        return PremiumCombo()


if __name__ == "__main__":
    trial_combo_factory = TrialComboFactory()
    trial_combo = trial_combo_factory.create_combo()
    trial_combo.task()

    basic_combo_factory = BasicComboFactory()
    basic_combo = basic_combo_factory.create_combo()
    basic_combo.task()

    premium_combo_factory = PremiumComboFactory()
    premium_combo = premium_combo_factory.create_combo()
    premium_combo.task()

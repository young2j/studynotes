#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: factory.py
# Created Date: 2021-11-29 06:19:19
# Author: ysj
# Description:  工厂类
###
from abc import ABCMeta, abstractmethod

from combo import TrialCombo, BasicCombo, PremiumCombo
from reporter import TrialReporter, BasicReporter, PremiumReporter


class Factory(metaclass=ABCMeta):
    """抽象基类"""
    @abstractmethod
    def create_combo(self):
        pass

    @abstractmethod
    def create_reporter(self):
        pass


class TrialFactory(Factory):
    """具体实现类"""

    def create_combo(self):
        return TrialCombo()

    def create_reporter(self):
        return TrialReporter()


class BasicFactory(Factory):
    """具体实现类"""

    def create_combo(self):
        return BasicCombo()

    def create_reporter(self):
        return BasicReporter()


class PremiumFactory(Factory):
    """具体实现类"""

    def create_combo(self):
        return PremiumCombo()

    def create_reporter(self):
        return PremiumReporter()

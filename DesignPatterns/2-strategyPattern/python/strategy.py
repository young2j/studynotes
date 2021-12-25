#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: strategy.py
# Created Date: 2021-12-07 02:21:10
# Author: ysj
# Description: 策略类
###

from abc import ABCMeta, abstractmethod


class PriceCalculationStrategy(metaclass=ABCMeta):
    """抽象基类"""
    @abstractmethod
    def calculate_price(self, origin_price):
        pass


class KeepNormalStrategy(PriceCalculationStrategy):
    """原价策略"""

    def __init__(self):
        super().__init__()

    def calculate_price(self, origin_price):
        total_price = round(origin_price, 2)
        return total_price


class DiscountStrategy(PriceCalculationStrategy):
    """打折策略"""

    def __init__(self, discount):
        self.__discount = discount

    def calculate_price(self, origin_price):
        total_price = origin_price * self.__discount
        total_price = round(total_price, 2)
        return total_price


class FullReductionStrategy(PriceCalculationStrategy):
    """满减策略"""

    def __init__(self, full, reduction):
        self.__full = full
        self.__reduction = reduction

    def calculate_price(self, origin_price):
        if origin_price >= self.__full:
            origin_price -= self.__reduction
        total_price = round(origin_price, 2)
        return total_price


class DirectReductionStrategy(PriceCalculationStrategy):
    """无门槛策略"""

    def __init__(self, reduction):
        self.__reduction = reduction

    def calculate_price(self, origin_price):
        total_price = max(0, origin_price-self.__reduction)
        total_price = round(total_price, 2)
        return total_price

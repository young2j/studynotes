#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: context1.py
# Created Date: 2021-12-07 03:03:37
# Author: ysj
# Description:  策略上下文
###

from strategy import PriceCalculationStrategy


class PriceCalculationContext(object):
    """策略上下文-由客户端判断选择策略"""

    def __init__(self, strategy: PriceCalculationStrategy):
        self.__strategy = strategy

    def calculate_price(self, origin_price):
        total_price = self.__strategy.calculate_price(origin_price)
        return total_price

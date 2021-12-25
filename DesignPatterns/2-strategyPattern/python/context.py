#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: context.py
# Created Date: 2021-12-07 03:43:19
# Author: ysj
# Description:  策略上下文 + 简单工厂
###
from strategy import (
    KeepNormalStrategy, DiscountStrategy,
    FullReductionStrategy, DirectReductionStrategy,
)


class PriceCalculationContext(object):
    """策略上下文-由简单工厂判断创建策略"""

    def __init__(self, strategy_type):
        if strategy_type == "keep_normal":
            self.__strategy = KeepNormalStrategy()
        elif strategy_type == "discount":
            self.__strategy = DiscountStrategy(discount=0.8)
        elif strategy_type == "full_reduction":
            self.__strategy = FullReductionStrategy(full=500, reduction=200)
        elif strategy_type == "direct_reduction":
            self.__strategy = DirectReductionStrategy(reduction=100)
        else:
            self.__strategy = KeepNormalStrategy()

    def calculate_price(self, origin_price):
        total_price = self.__strategy.calculate_price(origin_price)
        return total_price

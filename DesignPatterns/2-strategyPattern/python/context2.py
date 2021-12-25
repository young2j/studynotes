#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: context2.py
# Created Date: 2021-12-13 03:55:57
# Author: ysj
# Description:  策略上下文 + HashMap
###

class PriceCalculationContext(object):
    """策略上下文-由HashMap存储策略算法"""

    def __init__(self):
        self.__strategyMap = {
            "keep_normal": self.keep_normal,
            "discount": self.discount,
            "full_reduction": self.full_reduction,
            "direct_reduction": self.direct_reduction,
        }

    @staticmethod
    def keep_normal(origin_price):
        return round(origin_price, 2)

    @staticmethod
    def discount(origin_price, discount):
        return round(origin_price * discount, 2)

    @staticmethod
    def full_reduction(origin_price, full, reduction):
        if origin_price < full:
            return round(origin_price, 2)
        return round(origin_price-reduction, 2)

    @staticmethod
    def direct_reduction(origin_price, reduction):
        return round(max(0, origin_price-reduction), 2)

    def calculate_price(self, strategy_type, origin_price, *args, **kwargs):
        calculate_func = self.__strategyMap[strategy_type]
        total_price = calculate_func(origin_price, *args, **kwargs)
        return total_price

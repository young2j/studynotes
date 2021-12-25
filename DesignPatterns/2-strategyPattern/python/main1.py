#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main1.py
# Created Date: 2021-12-07 03:14:12
# Author: ysj
# Description:  python 策略模式
###

from strategy import (
    KeepNormalStrategy, DiscountStrategy,
    FullReductionStrategy, DirectReductionStrategy,
)

from context1 import PriceCalculationContext


# 购买价格、数量
price = 599
quantity = 1
origin_price = price * quantity
print(f"单价:{price} 数量:{quantity} 原价:{origin_price}")

# 正常计算策略
keep_normal = KeepNormalStrategy()
ctx = PriceCalculationContext(keep_normal)
total_price = ctx.calculate_price(origin_price)
print(f"正常价格:{total_price}")

# 打折计算策略
discount = DiscountStrategy(discount=0.8)
ctx = PriceCalculationContext(discount)
total_price = ctx.calculate_price(origin_price)
print(f"八折价格:{total_price}")

# 满减策略
full_reduction = FullReductionStrategy(full=500, reduction=200)
ctx = PriceCalculationContext(full_reduction)
total_price = ctx.calculate_price(origin_price)
print(f"满500减200价格:{total_price}")

# 无门槛策略
direct_reduction = DirectReductionStrategy(reduction=100)
ctx = PriceCalculationContext(direct_reduction)
total_price = ctx.calculate_price(origin_price)
print(f"直减100价格:{total_price}")

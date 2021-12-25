#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2021-12-07 03:47:06
# Author: ysj
# Description:  python 策略模式 + 简单工厂
###

from context import PriceCalculationContext

# 购买价格、数量
price = 599
quantity = 1
origin_price = price * quantity
print(f"单价:{price} 数量:{quantity} 原价:{origin_price}")

# 正常计算策略
keep_normal = PriceCalculationContext(strategy_type="keep_normal")
total_price = keep_normal.calculate_price(origin_price)
print(f"正常价格:{total_price}")

# 打折计算策略
discount = PriceCalculationContext(strategy_type="discount")
total_price = discount.calculate_price(origin_price)
print(f"打折价格:{total_price}")

# 满减策略
full_reduction = PriceCalculationContext(strategy_type="full_reduction")
total_price = full_reduction.calculate_price(origin_price)
print(f"满减价格:{total_price}")

# 无门槛策略
direct_reduction = PriceCalculationContext(strategy_type="direct_reduction")
total_price = direct_reduction.calculate_price(origin_price)
print(f"直减价格:{total_price}")

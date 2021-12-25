#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2021-12-19 06:11:39
# Author: ysj
# Description:  python 状态模式
###

from context import OrderContext
from state import OrderPaying

init_state = OrderPaying()
# 支付中->未支付
ctx = OrderContext(init_state)
ctx.handle()

# 未支付->已过期
ctx.elapsed_30min = True
ctx.handle()

# 已过期->X 已支付
ctx.has_paid = True
ctx.handle()

# 支付中-> 已支付
ctx.set_state(init_state)
ctx.has_paid = True
ctx.handle()

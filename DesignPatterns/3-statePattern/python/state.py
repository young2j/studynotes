#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: state.py
# Created Date: 2021-12-19 06:12:44
# Author: ysj
# Description:  状态类
###

from abc import ABCMeta, abstractmethod
from context import OrderContext


class OrderState(metaclass=ABCMeta):
    """抽象基类"""
    @abstractmethod
    def handle(self, ctx: OrderContext):
        pass


class OrderPaying(OrderState):
    """订单支付中"""

    def handle(self, ctx: OrderContext):
        print("订单支付中...")
        if ctx.has_paid:
            ctx.set_state(OrderPaid())
            ctx.handle()
        elif ctx.elapsed_30min:
            ctx.set_state(OrderExpired())
            ctx.handle()
        else:
            ctx.set_state(OrderUnpay())
            ctx.handle()


class OrderUnpay(OrderState):
    """订单未支付"""

    def handle(self, ctx: OrderContext):
        if ctx.has_paid:
            ctx.set_state(OrderPaid())
            ctx.handle()
        elif ctx.elapsed_30min:
            ctx.set_state(OrderExpired())
            ctx.handle()
        else:
            print("订单未支付...")


class OrderPaid(OrderState):
    """订单已支付"""

    def handle(self, ctx: OrderContext):
        print("订单已支付。")


class OrderExpired(OrderState):
    """订单已过期"""

    def handle(self, ctx: OrderContext):
        print("订单已过期。")

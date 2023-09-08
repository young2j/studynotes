#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: caretaker.py
# Created Date: 2023-03-19 03:30:18
# Author: ysj
# Description:  备忘录模式-备忘录托管人
###

from memento import Memento


class CareTaker(object):
    mementos = []

    def put(self, m: Memento):
        self.mementos = [
            m,
            *self.mementos[:2]
        ]

    def take(self, i: int) -> Memento:
        if i < 0 or i > 2:
            i = 0
        return self.mementos[i]

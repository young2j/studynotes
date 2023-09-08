#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: memento.py
# Created Date: 2023-03-19 03:30:40
# Author: ysj
# Description:  备忘录模式-备忘录(状态保存者)
###

class Memento(object):
    def __init__(
        self,
        skills="大威天龙",
        hp=100,
        mp=100,
        attack=100,
        defense=100,
    ):
        self.skills = skills
        self.hp = hp
        self.mp = mp
        self.attack = attack
        self.defense = defense

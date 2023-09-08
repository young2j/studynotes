#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: originator.py
# Created Date: 2023-03-19 03:30:53
# Author: ysj
# Description:  备忘录模式-状态保存发起人
###
from memento import Memento


class RoleOriginator(object):
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

    # 保存状态
    def create_memento(self) -> Memento:
        mem = Memento(skills=self.skills,
                      hp=self.hp,
                      mp=self.mp,
                      attack=self.attack,
                      defense=self.defense,
                      )
        return mem

    # 恢复状态
    def set_memento(self, m: Memento):
        self.skills = m.skills
        self.hp = m.hp
        self.mp = m.mp
        self.attack = m.attack
        self.defense = m.defense

    # 显示状态
    def show_state(self, msg):
        print('''
=======================
%s
========角色属性========
技能：%s
血量：%d
蓝量：%d
攻击力：%d
防御力：%d
''' % (msg, self.skills, self.hp, self.mp, self.attack, self.defense))

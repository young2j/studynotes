#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2023-03-19 03:30:05
# Author: ysj
# Description:  备忘录模式客户端调用
###
from originator import RoleOriginator
from caretaker import CareTaker

# 主角
role = RoleOriginator(
    skills="大威天龙",
    hp=100,
    mp=100,
    attack=100,
    defense=100,
)

# 游戏状态托管人
caretaker = CareTaker()

# 保存进度
role.show_state("初始状态")
memento0 = role.create_memento()
caretaker.put(memento0)

# 攻打boss失败
role.hp = 0
role.mp = 10
role.attack = 1
role.defense = 20
role.show_state("攻打boss失败")

# 恢复之前的状态
memento0 = caretaker.take(0)
role.set_memento(memento0)
role.show_state("恢复初始状态")

# 打怪升级
role.skills = "超·大威天龙"
role.hp = 1000
role.mp = 1000
role.attack = 1000
role.defense = 1000

# 保存进度
role.show_state("打怪升级")
memento1 = role.create_memento()
caretaker.put(memento1)

# 挑战boss成功
role.hp = 2000
role.mp = 2000
role.attack = 2000
role.defense = 2000
role.show_state("挑战boss成功,属性翻倍")

# 保存进度
memento2 = role.create_memento()
caretaker.put(memento2)

# 想重新挑战boss
memento1 = caretaker.take(1)
role.set_memento(memento1)
role.show_state("重新挑战boss, 恢复之前状态")

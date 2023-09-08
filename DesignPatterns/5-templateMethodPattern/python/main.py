#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2023-03-13 03:22:16
# Author: ysj
# Description:  模板方式模式——客户端调用
###

from old_master import OldMasterForgeSword
from little_apprentice import LittleApprenticeForgeSword

print("=============老师傅铸剑============")
old_master = OldMasterForgeSword()
old_master.forge_sword()

print("=============小徒弟铸剑============")
little_apprentice = LittleApprenticeForgeSword()
little_apprentice.forge_sword()

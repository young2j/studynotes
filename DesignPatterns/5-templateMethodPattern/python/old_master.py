#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: old_master.py
# Created Date: 2023-03-13 03:22:45
# Author: ysj
# Description:  模板方式模式——老师傅铸剑
###

from template import ForgeSwordTemplate


class OldMasterForgeSword(ForgeSwordTemplate):
    def make_clay_model(self):
        '''制做泥范, 每道工序20分'''
        super().make_clay_model()

    def dispense_materials(self):
        '''调剂材料, 每道工序20分'''
        super().dispense_materials()

    def smelting_materials(self):
        '''熔炼原料, 每道工序20分'''
        super().smelting_materials()

    def water_shaping(self):
        '''浇灌成形, 每道工序20分'''
        super().water_shaping()

    def repair_processing(self):
        '''修治加工, 每道工序20分'''
        super().repair_processing()

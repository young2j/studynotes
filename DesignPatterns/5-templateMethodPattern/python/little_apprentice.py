#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: little_apprentice.py
# Created Date: 2023-03-13 03:23:08
# Author: ysj
# Description:  模板方式模式——小徒弟铸剑
###

from template import ForgeSwordTemplate


class LittleApprenticeForgeSword(ForgeSwordTemplate):
    def make_clay_model(self):
        '''制做泥范, 每道工序20分'''
        super().make_clay_model()

    def dispense_materials(self):
        print("调剂材料比例不均 +10分")
        self.score += 10

    def smelting_materials(self):
        print("熔炼原料火候不够 +10分")
        self.score += 10

    def water_shaping(self):
        '''浇灌成形, 每道工序20分'''
        super().water_shaping()

    def repair_processing(self):
        '''修治加工, 每道工序20分'''
        super().repair_processing()

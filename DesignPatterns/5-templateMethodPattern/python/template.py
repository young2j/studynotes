#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: template.py
# Created Date: 2023-03-13 03:21:58
# Author: ysj
# Description:  python 模板方法
###

from abc import ABCMeta, abstractmethod


class ForgeSwordTemplate(metaclass=ABCMeta):
    score = 0

    def forge_sword(self):
        '''铸剑模板方法'''
        self.make_clay_model()
        self.dispense_materials()
        self.smelting_materials()
        self.water_shaping()
        self.repair_processing()

        self.result()

    @abstractmethod
    def make_clay_model(self):
        '''制做泥范, 每道工序20分'''
        print("制做泥范完美 +20分")
        self.score += 20

    @abstractmethod
    def dispense_materials(self):
        '''调剂材料, 每道工序20分'''
        print("调剂材料完美 +20分")
        self.score += 20

    @abstractmethod
    def smelting_materials(self):
        '''熔炼原料, 每道工序20分'''
        print("熔炼原料完美 +20分")
        self.score += 20

    @abstractmethod
    def water_shaping(self):
        '''浇灌成形, 每道工序20分'''
        print("浇灌成形完美 +20分")
        self.score += 20

    @abstractmethod
    def repair_processing(self):
        '''修治加工, 每道工序20分'''
        print("修治加工完美 +20分")
        self.score += 20

    def result(self):
        print("总分:", self.score)
        if self.score == 100:
            print("获得绝世好剑!!!")
        elif self.score > 80:
            print("获得一把好剑!!")
        else:
            print("获得一把村好剑!")

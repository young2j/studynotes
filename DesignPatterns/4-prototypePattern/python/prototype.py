#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: prototype.py
# Created Date: 2023-03-06 03:09:23
# Author: ysj
# Description: 原型类接口及其实现
###

from abc import ABCMeta, abstractmethod
# from typing import Self


class ClayPerson(metaclass=ABCMeta):
    """抽象原型类"""
    # @abstractmethod
    # def clone() -> Self: # python3.11
    #     pass

    @abstractmethod
    def clone():
        pass


class ClayBoy(ClayPerson):
    """男娃娃小泥人儿"""

    def __init__(self,
                 name,
                 eye,
                 nose,
                 skin,
                 hands,
                 gender,
                 temper,
                 ):
        self.name = name
        self.eye = eye
        self.nose = nose
        self.skin = skin
        self.hands = hands
        self.gender = gender
        self.temper = temper

    def clone(self):
        self_clone = ClayBoy(
            name=self.name,
            eye=self.eye,
            nose=self.nose,
            skin=self.skin,
            hands=self.hands,
            gender=self.gender,
            temper=self.temper,
        )
        return self_clone

    def introduction(self):
        print("%s: 性别%s, 有%s的眼睛, %s, %s的皮肤,%d只手,脾气%s。" %
              (self.name, self.gender, self.eye, self.nose,
               self.skin, self.hands, self.temper)
              )


class ClayGirl(ClayPerson):
    """女娃娃小泥人儿"""

    def __init__(self,
                 name,
                 eye,
                 nose,
                 skin,
                 hands,
                 gender,
                 temper,
                 ):
        self.name = name
        self.eye = eye
        self.nose = nose
        self.skin = skin
        self.hands = hands
        self.gender = gender
        self.temper = temper

    def clone(self):
        self_clone = ClayGirl(
            name=self.name,
            eye=self.eye,
            nose=self.nose,
            skin=self.skin,
            hands=self.hands,
            gender=self.gender,
            temper=self.temper,
        )
        return self_clone

    def introduction(self):
        print("%s: 性别%s, 有%s的眼睛, %s, %s的皮肤,%d只手,脾气%s。" %
              (self.name, self.gender, self.eye, self.nose,
               self.skin, self.hands, self.temper)
              )

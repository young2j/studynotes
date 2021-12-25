#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: reporter.py
# Created Date: 2021-11-29 06:18:27
# Author: ysj
# Description:  报告类
###
from abc import ABCMeta, abstractmethod


class Reporter(metaclass=ABCMeta):
    """抽象基类"""
    report_type = "报告类型"

    @abstractmethod
    def export_report(self):
        return

    @property
    @abstractmethod
    def report_type(self):
        return self.report_type


class TrialReporter(Reporter):
    """具体实现类"""
    report_type = "简要excel报告"

    def export_report(self):
        print(f"导出{self.report_type}")
        return


class BasicReporter(Reporter):
    """具体实现类"""
    report_type = "详细excel报告"

    def export_report(self):
        print(f"导出{self.report_type}")
        return


class PremiumReporter(Reporter):
    """具体实现类"""
    report_type = "精美pdf报告"

    def export_report(self):
        print(f"导出{self.report_type}")
        return

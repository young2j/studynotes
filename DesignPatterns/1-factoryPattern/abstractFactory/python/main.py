#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: main.py
# Created Date: 2021-11-28 07:20:56
# Author: ysj
# Description:  python 抽象工厂
###

from factory import TrialFactory, BasicFactory, PremiumFactory

trial_factory = TrialFactory()
trial_combo = trial_factory.create_combo()
trial_combo.task()
trial_reporter = trial_factory.create_reporter()
trial_reporter.export_report()

basic_factory = BasicFactory()
basic_combo = basic_factory.create_combo()
basic_combo.task()
basic_reporter = basic_factory.create_reporter()
basic_reporter.export_report()

premium_factory = PremiumFactory()
premium_combo = premium_factory.create_combo()
premium_combo.task()
premium_reporter = premium_factory.create_reporter()
premium_reporter.export_report()

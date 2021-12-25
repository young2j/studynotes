/**
 * -------------------------------------------------------
 * File: index.ts
 * Created Date: 2021-11-29 04:20:50
 * Author: ysj
 * Description: ts 抽象工厂
 * -------------------------------------------------------
 */

import { TrialFactory, BasicFactory, PremiumFactory } from './factory';

const trialFactory = new TrialFactory();
const trialCombo = trialFactory.createCombo();
trialCombo.task();
const trialReporter = trialFactory.createReporter();
trialReporter.exportReport();

const basicFactory = new BasicFactory();
const basicCombo = basicFactory.createCombo();
basicCombo.task();
const basicReporter = basicFactory.createReporter();
basicReporter.exportReport();

const premiumFactory = new PremiumFactory();
const premiumCombo = premiumFactory.createCombo();
premiumCombo.task();
const premiumReporter = premiumFactory.createReporter();
premiumReporter.exportReport();

/*
 * File: main.cpp
 * Created Date: 2021-11-29 06:29:10
 * Author: ysj
 * Description:  cpp 抽象工厂
 */
#include <iostream>
#include "factory.hpp"
using namespace std;

// 客户端调用
int main()
{
    Factory *trialFactory = new TrialFactory();
    Combo *trialCombo = trialFactory->createCombo();
    trialCombo->task();
    Reporter *trialReporter = trialFactory->createReporter();
    trialReporter->exportReport();

    Factory *basicFactory = new BasicFactory();
    Combo *basicCombo = basicFactory->createCombo();
    basicCombo->task();
    Reporter *basicReporter = basicFactory->createReporter();
    basicReporter->exportReport();

    Factory *premiumFactory = new PremiumFactory();
    Combo *premiumCombo = premiumFactory->createCombo();
    premiumCombo->task();
    Reporter *premiumReporter = premiumFactory->createReporter();
    premiumReporter->exportReport();

    return 0;
};

/*
 * File: include/factory.hpp
 * Created Date: 2021-12-01 06:16:13
 * Author: ysj
 * Description:  工厂类头文件(含实现)
 */

#pragma once
#include "combo.hpp"
#include "reporter.hpp"

// 抽象类
class Factory
{
public:
    virtual Combo *createCombo() = 0;
    virtual Reporter *createReporter() = 0;
};

// 具体实现类
class TrialFactory : public Factory
{
public:
    Combo *createCombo()
    {
        return new TrialCombo();
    };
    Reporter *createReporter()
    {
        return new TrialReporter();
    };
};

// 具体实现类
class BasicFactory : public Factory
{
public:
    Combo *createCombo()
    {
        return new BasicCombo();
    };
    Reporter *createReporter()
    {
        return new BasicReporter();
    };
};

// 具体实现类
class PremiumFactory : public Factory
{
public:
    Combo *createCombo()
    {
        return new PremiumCombo();
    };
    Reporter *createReporter()
    {
        return new PremiumReporter();
    };
};

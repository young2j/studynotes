/*
 * File: include/reporter.hpp
 * Created Date: 2021-12-01 05:22:23
 * Author: ysj
 * Description:  报告类头文件(含实现)
 */

#pragma once

#include <iostream>
#include <string>
using namespace std;

// 抽象类
class Reporter
{
public:
    string reportType;

public:
    virtual void exportReport() = 0;
};

// 具体实现类
class TrialReporter : public Reporter
{
public:
    TrialReporter()
    {
        reportType = "简要excel报告";
    };

public:
    void exportReport()
    {
        cout << "导出" << this->reportType << endl;
    };
};

// 具体实现类
class BasicReporter : public Reporter
{
public:
    BasicReporter()
    {
        reportType = "详细excel报告";
    };

public:
    void exportReport()
    {
        cout << "导出" << this->reportType << endl;
    };
};

// 具体实现类
class PremiumReporter : public Reporter
{
public:
    PremiumReporter()
    {
        reportType = "精美pdf报告";
    };

public:
    void exportReport()
    {
        cout << "导出" << this->reportType << endl;
    };
};

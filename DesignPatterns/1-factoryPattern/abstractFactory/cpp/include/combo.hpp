/*
 * File: include/combo.hpp
 * Created Date: 2021-12-01 02:46:27
 * Author: ysj
 * Description:  套餐类头文件(含实现)
 */

#pragma once

#include <iostream>
#include <string>
using namespace std;

// 抽象类
class Combo
{
public:
    string name;
    int price;

public:
    virtual void task() = 0;
};

// 具体实现类
class TrialCombo : public Combo
{
public:
    TrialCombo()
    {
        name = "体验版";
        price = 299;
    };
    void task()
    {
        cout << name << "-> 执行价值" << price << "元的服务" << endl;
    };
};

// 具体实现类
class BasicCombo : public Combo
{
public:
    BasicCombo()
    {
        name = "基础版";
        price = 599;
    };
    void task()
    {
        cout << name << "-> 执行价值" << price << "元的服务" << endl;
    };
};

// 具体实现类
class PremiumCombo : public Combo
{
public:
    PremiumCombo()
    {
        name = "高级版";
        price = 1999;
    };
    void task()
    {
        cout << name << "-> 执行价值" << price << "元的服务" << endl;
    };
};

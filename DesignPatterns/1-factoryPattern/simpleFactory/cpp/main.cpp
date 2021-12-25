/*
 * File: main.cpp
 * Created Date: 2021-11-25 10:10:45
 * Author: ysj
 * Description:  cpp 简单工厂
 */

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

//具体实现类
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

//具体实现类
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

enum ComboType
{
    trial,
    basic,
    premium
};

// 简单工厂类
class SimpleFactory
{
public:
    static Combo *createCombo(ComboType comboType)
    {
        switch (comboType)
        {
        case trial:
            return new TrialCombo();
        case basic:
            return new BasicCombo();
        case premium:
            return new PremiumCombo();
        default:
            return new TrialCombo();
        }
    };
};

// 客户端调用
int main()
{
    SimpleFactory factory = SimpleFactory();
    Combo *trialCombo = factory.createCombo(trial);
    trialCombo->task();

    Combo *basicCombo = factory.createCombo(basic);
    basicCombo->task();

    Combo *premiumCombo = factory.createCombo(premium);
    premiumCombo->task();

    return 0;
}

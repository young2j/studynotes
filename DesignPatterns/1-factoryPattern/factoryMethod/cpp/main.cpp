/*
 * Created Date: 2021-11-26 12:38:36
 * Author: ysj
 * Description:  cpp 工厂方法
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

enum ComboType
{
    trial,
    basic,
    premium
};

// 抽象工厂类
class Factory
{
public:
    virtual Combo *createCombo() = 0;
};

// 具体工厂类
class TrialComboFactory : public Factory
{
public:
    Combo *createCombo()
    {
        return new TrialCombo();
    };
};

// 具体工厂类
class BasicComboFactory : public Factory
{
public:
    Combo *createCombo()
    {
        return new BasicCombo();
    };
};

// 具体工厂类
class PremiumComboFactory : public Factory
{
public:
    Combo *createCombo()
    {
        return new PremiumCombo();
    };
};

// 客户端调用
int main()
{
    Factory *trialComboFactory = new TrialComboFactory();
    Combo *trialCombo = trialComboFactory->createCombo();
    trialCombo->task();

    Factory *basicComboFactory = new BasicComboFactory();
    Combo *basicCombo = basicComboFactory->createCombo();
    basicCombo->task();

    Factory *premiumComboFactory = new PremiumComboFactory();
    Combo *premiumCombo = premiumComboFactory->createCombo();
    premiumCombo->task();

    return 0;
}

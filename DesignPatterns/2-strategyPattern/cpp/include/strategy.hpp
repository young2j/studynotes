/*
 * File: strategy.hpp
 * Created Date: 2021-12-12 05:08:43
 * Author: ysj
 * Description:  策略类
 */
#pragma once
#include <iostream>
using namespace std;

// 价格计算抽象策略类
class PriceCalculationStrategy
{
public:
    virtual float calculatePrice(float originPrice) = 0;
};

// 原价策略
class KeepNormalStrategy : public PriceCalculationStrategy
{
public:
    virtual float calculatePrice(float originPrice)
    {
        return originPrice;
    }
};

// 打折策略
class DiscountStrategy : public PriceCalculationStrategy
{
private:
    float discount;

public:
    DiscountStrategy(float discount)
    {
        this->discount = discount;
    }
    virtual float calculatePrice(float originPrice)
    {
        return originPrice * discount;
    }
};

// 满减策略
class FullReductionStrategy : public PriceCalculationStrategy
{
private:
    float full;
    float reduction;

public:
    FullReductionStrategy(float full, float reduction)
    {
        this->full = full;
        this->reduction = reduction;
    }
    virtual float calculatePrice(float originPrice)
    {
        float totalPrice = originPrice;
        if (originPrice >= full)
        {
            totalPrice = originPrice - reduction;
        }
        return totalPrice;
    }
};

// 无门槛策略
class DirectReductionStrategy : public PriceCalculationStrategy
{
private:
    float reduction;

public:
    DirectReductionStrategy(float reduction)
    {
        this->reduction = reduction;
    }
    virtual float calculatePrice(float originPrice)
    {
        float totalPrice = originPrice - reduction;
        if (totalPrice < 0)
        {
            return 0;
        }
        return totalPrice;
    }
};

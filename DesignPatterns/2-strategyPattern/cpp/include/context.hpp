/*
 * File: context.hpp
 * Created Date: 2021-12-12 05:08:58
 * Author: ysj
 * Description:  策略上下文类
 */
#pragma once
#include <iostream>
#include "strategy.hpp"
using namespace std;

class PriceCalculationContext
{
private:
    PriceCalculationStrategy *strategy;

public:
    PriceCalculationContext(PriceCalculationStrategy *strategy)
    {
        this->strategy = strategy;
    }
    float calculatePrice(float originPrice)
    {
        float totalPrice = strategy->calculatePrice(originPrice);
        return totalPrice;
    }
};

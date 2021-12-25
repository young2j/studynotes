/*
 * File: main.cpp
 * Created Date: 2021-12-12 05:06:54
 * Author: ysj
 * Description:  cpp策略模式
 */
#include <iostream>
#include "strategy.hpp"
#include "context.hpp"
using namespace std;

int main()
{
    // 购买价格、数量
    float price = 599;
    float quantity = 1;
    float originPrice = price * quantity;
    cout << "单价:" << price << " 数量:" << quantity << " 原价:" << originPrice << endl;

    // 正常计算策略
    PriceCalculationStrategy *keepNormal = new KeepNormalStrategy();
    PriceCalculationContext *ctx = new PriceCalculationContext(keepNormal);
    float totalPrice = ctx->calculatePrice(originPrice);
    cout << "正常价格:" << totalPrice << endl;

    // 打折计算策略
    PriceCalculationStrategy *discount = new DiscountStrategy(0.8);
    ctx = new PriceCalculationContext(discount);
    totalPrice = ctx->calculatePrice(originPrice);
    cout << "八折价格:" << totalPrice << endl;

    // 满减策略
    PriceCalculationStrategy *fullReduction = new FullReductionStrategy(500, 200);
    ctx = new PriceCalculationContext(fullReduction);
    totalPrice = ctx->calculatePrice(originPrice);
    cout << "满500减200价格:" << totalPrice << endl;

    // 无门槛策略
    PriceCalculationStrategy *directReduction = new DirectReductionStrategy(100);
    ctx = new PriceCalculationContext(directReduction);
    totalPrice = ctx->calculatePrice(originPrice);
    cout << "直减100价格:" << totalPrice << endl;
    
    return 0;
}

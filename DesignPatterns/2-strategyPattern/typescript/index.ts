/**
 * -------------------------------------------------------
 * File: index.ts
 * Created Date: 2021-12-12 01:45:20
 * Author: ysj
 * Description: ts 策略模式
 * -------------------------------------------------------
 */

import PriceCalculationContext from './context';
import {
  KeepNormalStrategy,
  DiscountStrategy,
  FullReductionStrategy,
  DirectReductionStrategy,
} from './strategy';

// 购买价格、数量
const price = 599;
const quantity = 1;
const originPrice = price * quantity;
console.log(`单价:${price} 数量:${quantity} 原价:${originPrice}`);

// 正常计算策略
const keepNormal = new KeepNormalStrategy();
let ctx = new PriceCalculationContext(keepNormal);
let totalPrice = ctx.calculatePrice(originPrice);
console.log(`正常价格:${totalPrice}`);

// 打折计算策略
const discount = new DiscountStrategy(0.8);
ctx = new PriceCalculationContext(discount);
totalPrice = ctx.calculatePrice(originPrice);
console.log(`八折价格:${totalPrice}`);

// 满减策略
const fullReduction = new FullReductionStrategy(500, 200);
ctx = new PriceCalculationContext(fullReduction);
totalPrice = ctx.calculatePrice(originPrice);
console.log(`满500减200价格:${totalPrice}`);

// 无门槛策略
const directReduction = new DirectReductionStrategy(100);
ctx = new PriceCalculationContext(directReduction);
totalPrice = ctx.calculatePrice(originPrice);
console.log(`直减100价格:${totalPrice}`);

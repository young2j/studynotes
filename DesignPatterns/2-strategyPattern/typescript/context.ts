/**
 * -------------------------------------------------------
 * File: context.ts
 * Created Date: 2021-12-12 01:46:28
 * Author: ysj
 * Description: 策略上下文
 * -------------------------------------------------------
 */

import { IPriceCalculationStrategy } from './strategy';

export default class PriceCalculationContext {
  private strategy: IPriceCalculationStrategy;
  constructor(strategy: IPriceCalculationStrategy) {
    this.strategy = strategy;
  }
  calculatePrice(originPrice: number): number {
    const totalPrice = this.strategy.calculatePrice(originPrice);
    return totalPrice;
  }
}

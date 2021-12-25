/**
 * -------------------------------------------------------
 * File: strategy.ts
 * Created Date: 2021-12-12 01:46:00
 * Author: ysj
 * Description: 策略类
 * -------------------------------------------------------
 */

/**策略接口 */
export interface IPriceCalculationStrategy {
  calculatePrice(originPrice: number): number;
}

/**原价策略 */
export class KeepNormalStrategy implements IPriceCalculationStrategy {
  calculatePrice(originPrice: number) {
    const totalPrice = Number(originPrice.toFixed(2));
    return totalPrice;
  }
}

/**打折策略 */
export class DiscountStrategy implements IPriceCalculationStrategy {
  private discount: number;
  constructor(discount: number) {
    this.discount = discount;
  }
  calculatePrice(originPrice: number) {
    let totalPrice = originPrice * this.discount;
    totalPrice = Number(totalPrice.toFixed(2));
    return totalPrice;
  }
}

/**满减策略 */
export class FullReductionStrategy implements IPriceCalculationStrategy {
  private full: number;
  private reduction: number;
  constructor(full: number, reduction: number) {
    this.full = full;
    this.reduction = reduction;
  }
  calculatePrice(originPrice: number) {
    let totalPrice =
      originPrice >= this.full ? originPrice - this.reduction : originPrice;
    totalPrice = Number(totalPrice.toFixed(2));
    return totalPrice;
  }
}

/**无门槛策略 */
export class DirectReductionStrategy implements IPriceCalculationStrategy {
  private reduction: number;
  constructor(reduction: number) {
    this.reduction = reduction;
  }
  calculatePrice(originPrice: number) {
    let totalPrice = Math.max(0, originPrice - this.reduction);
    totalPrice = Number(totalPrice.toFixed(2));
    return totalPrice;
  }
}

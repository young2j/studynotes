/*
 * File: context.go
 * Created Date: 2021-12-12 04:25:39
 * Author: ysj
 * Description:  策略上下文
 */

package main

type PriceCalculationContext struct {
	strategy IPriceCalculationStrategy
}

func NewPriceCalculationContext(strategy IPriceCalculationStrategy) *PriceCalculationContext {
	return &PriceCalculationContext{
		strategy: strategy,
	}
}

func (p *PriceCalculationContext) calculatePrice(originPrice float64) float64 {
	return p.strategy.calculatePrice(originPrice)
}

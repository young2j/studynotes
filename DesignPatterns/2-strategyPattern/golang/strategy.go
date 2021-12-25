/*
 * File: strategy.go
 * Created Date: 2021-12-12 04:25:14
 * Author: ysj
 * Description:  策略
 */
package main

// 策略接口
type IPriceCalculationStrategy interface {
	calculatePrice(originPrice float64) float64
}

// 原价策略
type KeepNormalStrategy struct{}

func NewKeepNormalStrategy() IPriceCalculationStrategy {
	return &KeepNormalStrategy{}
}

func (k *KeepNormalStrategy) calculatePrice(originPrice float64) float64 {
	return originPrice
}

// 打折策略
type DiscountStrategy struct {
	discount float64
}

func NewDiscountStrategy(discount float64) IPriceCalculationStrategy {
	return &DiscountStrategy{
		discount: discount,
	}
}
func (d *DiscountStrategy) calculatePrice(originPrice float64) float64 {
	return originPrice * d.discount
}

// 满减策略
type FullReductionStrategy struct {
	full      float64
	reduction float64
}

func NewFullReductionStrategy(full, reduction float64) IPriceCalculationStrategy {
	return &FullReductionStrategy{
		full:      full,
		reduction: reduction,
	}
}
func (f *FullReductionStrategy) calculatePrice(originPrice float64) float64 {
	totalPrice := originPrice
	if originPrice >= f.full {
		totalPrice = originPrice - f.reduction
	}
	return totalPrice
}

// 无门槛策略
type DirectReductionStrategy struct {
	reduction float64
}

func NewDirectReductionStrategy(reduction float64) IPriceCalculationStrategy {
	return &DirectReductionStrategy{
		reduction: reduction,
	}
}
func (d *DirectReductionStrategy) calculatePrice(originPrice float64) float64 {
	totalPrice := originPrice - d.reduction
	if totalPrice < 0 {
		return 0
	}
	return totalPrice
}

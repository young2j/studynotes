/*
 * File: main.go
 * Created Date: 2021-12-12 04:26:03
 * Author: ysj
 * Description:  golang策略模式
 */

package main

import "fmt"

func main() {
	// 购买价格、数量
	price := 599.0
	quantity := 1.0
	originPrice := price * quantity
	fmt.Printf("单价:%.2f 数量:%.2f 原价:%.2f\n", price, quantity, originPrice)

	// 正常计算策略
	keepNormal := NewKeepNormalStrategy()
	ctx := NewPriceCalculationContext(keepNormal)
	totalPrice := ctx.calculatePrice(originPrice)
	fmt.Printf("正常价格:%.2f\n", totalPrice)

	// 打折计算策略
	discount := NewDiscountStrategy(0.8)
	ctx = NewPriceCalculationContext(discount)
	totalPrice = ctx.calculatePrice(originPrice)
	fmt.Printf("八折价格:%.2f\n", totalPrice)

	// 满减策略
	fullReduction := NewFullReductionStrategy(500, 200)
	ctx = NewPriceCalculationContext(fullReduction)
	totalPrice = ctx.calculatePrice(originPrice)
	fmt.Printf("满500减200价格:%.2f\n", totalPrice)

	// 无门槛策略
	directReduction := NewDirectReductionStrategy(100)
	ctx = NewPriceCalculationContext(directReduction)
	totalPrice = ctx.calculatePrice(originPrice)
	fmt.Printf("直减100价格:%.2f\n", totalPrice)
}

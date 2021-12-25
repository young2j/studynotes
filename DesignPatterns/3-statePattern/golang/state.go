/*
 * File: state.go
 * Created Date: 2021-12-19 11:00:08
 * Author: ysj
 * Description:  状态类
 */

package main

import "fmt"

type IOrderState interface {
	Handle(ctx *OrderContext)
}

/* 订单支付中 */
type OrderPaying struct{}

// 初始化
func NewOrderPaying() IOrderState {
	return OrderPaying{}
}

// 行为
func (o OrderPaying) Handle(ctx *OrderContext) {
	fmt.Println("订单支付中...")
	if ctx.HasPaid {
		ctx.SetState(NewOrderPaid())
		ctx.Handle()
	} else if ctx.Elapsed30min {
		ctx.SetState(NewOrderExpired())
		ctx.Handle()
	} else {
		ctx.SetState(NewOrderUnpay())
		ctx.Handle()
	}
}

/* 订单未支付 */
type OrderUnpay struct{}

// 初始化
func NewOrderUnpay() IOrderState {
	return OrderUnpay{}
}

// 行为
func (o OrderUnpay) Handle(ctx *OrderContext) {
	if ctx.HasPaid {
		ctx.SetState(NewOrderPaid())
		ctx.Handle()
	} else if ctx.Elapsed30min {
		ctx.SetState(NewOrderExpired())
		ctx.Handle()
	} else {
		fmt.Println("订单未支付...")
	}

}

/* 订单已支付 */
type OrderPaid struct{}

// 初始化
func NewOrderPaid() IOrderState {
	return OrderPaid{}
}

// 行为
func (o OrderPaid) Handle(ctx *OrderContext) {
	fmt.Println("订单已支付。")
}

/* 订单已过期 */
type OrderExpired struct{}

// 初始化
func NewOrderExpired() IOrderState {
	return OrderExpired{}
}

// 行为
func (o OrderExpired) Handle(ctx *OrderContext) {
	fmt.Println("订单已过期。")
}

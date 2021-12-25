/*
 * File: main.go
 * Created Date: 2021-12-19 10:59:45
 * Author: ysj
 * Description:  golang 状态模式
 */

package main

func main() {
	initState := NewOrderPaying()
	// 支付中->未支付
	ctx := NewOrderContext(initState)
	ctx.Handle()

	// 未支付->已过期
	ctx.Elapsed30min = true
	ctx.Handle()

	// 已过期->X 已支付
	ctx.HasPaid = true
	ctx.Handle()

	// 支付中-> 已支付
	ctx.SetState(initState)
	ctx.HasPaid = true
	ctx.Handle()
}

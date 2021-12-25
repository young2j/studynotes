/*
 * File: context.go
 * Created Date: 2021-12-19 11:00:25
 * Author: ysj
 * Description:  状态上下文类
 */

package main

type OrderContext struct {
	state        IOrderState
	HasPaid      bool
	Elapsed30min bool
}

// 初始化
func NewOrderContext(state IOrderState) *OrderContext {
	return &OrderContext{
		state:        state,
		HasPaid:      false,
		Elapsed30min: false,
	}
}

// 获取当前状态
func (o *OrderContext) GetState() IOrderState {
	return o.state
}

// 设置当前状态
func (o *OrderContext) SetState(state IOrderState) {
	o.state = state
}

// 当前状态下的行为
func (o *OrderContext) Handle() {
	o.state.Handle(o)
}

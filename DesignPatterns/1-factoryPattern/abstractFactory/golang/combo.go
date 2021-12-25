/*
 * File: combo.go
 * Created Date: 2021-11-29 06:24:53
 * Author: ysj
 * Description:  套餐
 */

package main

import "fmt"

// 接口
type ICombo interface {
	task()
}

type Combo struct {
	Name  string
	Price int
}

func (t *Combo) task() {
	fmt.Printf("%s-> 执行价值%d元的服务\n", t.Name, t.Price)
}

// 具体实现
func NewTrialCombo() ICombo {
	return &Combo{
		Name:  "体验版",
		Price: 299,
	}
}

// 具体实现
func NewBasicCombo() ICombo {
	return &Combo{
		Name:  "基础版",
		Price: 599,
	}
}

// 具体实现
func NewPremiumCombo() ICombo {
	return &Combo{
		Name:  "高级版",
		Price: 1999,
	}
}

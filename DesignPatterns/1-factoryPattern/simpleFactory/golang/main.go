/*
 * File: main.go
 * Created Date: 2021-11-24 12:34:23
 * Author: ysj
 * Description: golang 简单工厂
 */
package main

import (
	"fmt"
)

// 接口
type ICombo interface {
	task()
}

type Combo struct {
	Name  string
	Price int
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
func NewPremuimCombo() ICombo {
	return &Combo{
		Name:  "高级版",
		Price: 1999,
	}
}

func (t *Combo) task() {
	fmt.Printf("%s-> 执行价值%d元的服务\n", t.Name, t.Price)
}

// 简单工厂
type SimpleFactory struct{}

func (s *SimpleFactory) createCombo(comboType string) ICombo {
	switch comboType {
	case "trial":
		return NewTrialCombo()
	case "basic":
		return NewBasicCombo()
	case "premium":
		return NewPremuimCombo()
	default:
		return NewTrialCombo()
	}
}

// 客户端调用
func main() {
	factory := &SimpleFactory{}

	trialCombo := factory.createCombo("trial")
	trialCombo.task()

	basicCombo := factory.createCombo("basic")
	basicCombo.task()

	premiumCombo := factory.createCombo("premium")
	premiumCombo.task()
}

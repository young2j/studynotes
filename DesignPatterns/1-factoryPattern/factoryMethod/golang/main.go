/*
 * Created Date: 2021-11-26 12:31:54
 * Author: ysj
 * Description:  golang 工厂方法
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

func (t *Combo) task() {
	fmt.Printf("%s-> 执行价值%d元的服务\n", t.Name, t.Price)
}

// 工厂接口
type IFactory interface {
	createCombo() ICombo
}

// 具体工厂
type TrialComboFactory struct{}

func (t *TrialComboFactory) createCombo() ICombo {
	return NewTrialCombo()
}

// 具体工厂
type BasicComboFactory struct{}

func (t *BasicComboFactory) createCombo() ICombo {
	return NewBasicCombo()
}

// 具体工厂
type PremiumComboFactory struct{}

func (t *PremiumComboFactory) createCombo() ICombo {
	return NewPremiumCombo()
}

// 客户端调用
func main() {
	trialComboFactory := &TrialComboFactory{}
	trialCombo := trialComboFactory.createCombo()
	trialCombo.task()

	basicComboFactory := &BasicComboFactory{}
	basicCombo := basicComboFactory.createCombo()
	basicCombo.task()

	premiumComboFactory := &PremiumComboFactory{}
	premiumCombo := premiumComboFactory.createCombo()
	premiumCombo.task()

}

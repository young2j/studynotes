/*
 * File: factory.go
 * Created Date: 2021-11-29 06:25:22
 * Author: ysj
 * Description:  工厂
 */
package main

// 接口
type IFactory interface {
	createCombo() ICombo
	createReporter() IReporter
}

// "体验版"工厂
type TrialFactory struct{}

func (t *TrialFactory) createCombo() ICombo {
	return NewTrialCombo()
}
func (t *TrialFactory) createReporter() IReporter {
	return NewTrialReporter()
}

// "基础版"工厂
type BasicFactory struct{}

func (t *BasicFactory) createCombo() ICombo {
	return NewBasicCombo()
}
func (t *BasicFactory) createReporter() IReporter {
	return NewBasicReporter()
}

// "高级版"工厂
type PremiumFactory struct{}

func (t *PremiumFactory) createCombo() ICombo {
	return NewPremiumCombo()
}
func (t *PremiumFactory) createReporter() IReporter {
	return NewPremiumReporter()
}

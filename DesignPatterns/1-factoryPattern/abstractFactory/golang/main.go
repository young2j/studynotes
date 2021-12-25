/*
 * File: main.go
 * Created Date: 2021-11-29 02:03:08
 * Author: ysj
 * Description:  golang 抽象工厂
 */

package main

func main() {
	trialFactory := &TrialFactory{}
	trialCombo := trialFactory.createCombo()
	trialCombo.task()
	trialReporter := trialFactory.createReporter()
	trialReporter.exportReport()

	basicFactory := &BasicFactory{}
	basicCombo := basicFactory.createCombo()
	basicCombo.task()
	basicReporter := basicFactory.createReporter()
	basicReporter.exportReport()

	premiumFactory := &PremiumFactory{}
	premiumCombo := premiumFactory.createCombo()
	premiumCombo.task()
	premiumReporter := premiumFactory.createReporter()
	premiumReporter.exportReport()

}

/*
 * File: main.go
 * Created Date: 2023-03-19 02:31:23
 * Author: ysj
 * Description:  备忘录模式客户端调用
 */

package main

func main() {
	// 主角
	role := &RoleOriginator{
		Skills:  "大威天龙",
		HP:      100,
		MP:      100,
		Attack:  100,
		Defense: 100,
	}
	// 游戏状态托管人
	caretaker := &CareTaker{}

	// 保存进度
	role.ShowState("初始状态")
	memento0 := role.CreateMemento()
	caretaker.Put(memento0)

	// 攻打boss失败
	role.HP = 0
	role.MP = 10
	role.Attack = 1
	role.Defense = 20
	role.ShowState("攻打boss失败")

	// 恢复之前的状态
	memento0 = caretaker.Take(0)
	role.SetMemento(memento0)
	role.ShowState("恢复初始状态")

	// 打怪升级
	role.Skills = "超·大威天龙"
	role.HP = 1000
	role.MP = 1000
	role.Attack = 1000
	role.Defense = 1000

	// 保存进度
	role.ShowState("打怪升级")
	memento1 := role.CreateMemento()
	caretaker.Put(memento1)

	// 挑战boss成功
	role.HP = 2000
	role.MP = 2000
	role.Attack = 2000
	role.Defense = 2000
	role.ShowState("挑战boss成功,属性翻倍")

	// 保存进度
	memento2 := role.CreateMemento()
	caretaker.Put(memento2)

	// 想重新挑战boss
	memento1 = caretaker.Take(1)
	role.SetMemento(memento1)
	role.ShowState("重新挑战boss, 恢复之前状态")
}

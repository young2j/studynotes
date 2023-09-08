/*
 * File: originator.go
 * Created Date: 2023-03-19 02:31:55
 * Author: ysj
 * Description:  备忘录模式-状态保存发起人
 */

package main

import "fmt"

type RoleOriginator struct {
	Skills  string // 技能
	HP      int    // 血量
	MP      int    // 蓝量
	Attack  int    // 攻击力
	Defense int    // 防御力
}

// 保存状态
func (ro *RoleOriginator) CreateMemento() Memento {
	return Memento{
		Skills:  ro.Skills,
		HP:      ro.HP,
		MP:      ro.MP,
		Attack:  ro.Attack,
		Defense: ro.Defense,
	}
}

// 恢复状态
func (ro *RoleOriginator) SetMemento(m Memento) {
	ro.Skills = m.Skills
	ro.HP = m.HP
	ro.MP = m.MP
	ro.Attack = m.Attack
	ro.Defense = m.Defense
}

// 显示当前状态
func (ro *RoleOriginator) ShowState(msg string) {
	fmt.Printf(`
=======================
%s
========角色属性========
技能：%s
血量：%d
蓝量：%d
攻击力：%d
防御力：%d
`, msg, ro.Skills, ro.HP, ro.MP, ro.Attack, ro.Defense)
}

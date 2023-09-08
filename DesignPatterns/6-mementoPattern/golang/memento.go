/*
 * File: memento.go
 * Created Date: 2023-03-19 02:32:35
 * Author: ysj
 * Description:  备忘录模式-备忘录(状态保存者)
 */
package main

type Memento struct {
	Skills  string // 技能
	HP      int    // 血量
	MP      int    // 蓝量
	Attack  int    // 攻击力
	Defense int    // 防御力
}

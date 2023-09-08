/*
 * File: caretaker.go
 * Created Date: 2023-03-19 02:34:02
 * Author: ysj
 * Description:  备忘录模式-备忘录托管人
 */
package main

type CareTaker struct {
	Mementos [3]Memento // 备忘录，最多保存最近三次进度
}

// 放入1个Memento
func (c *CareTaker) Put(m Memento) {
	c.Mementos = [3]Memento{
		m,
		c.Mementos[0],
		c.Mementos[1],
	}
}

// 取出第i个Memento
func (c *CareTaker) Take(i int) Memento {
	if i < 0 || i > 2 {
		i = 0
	}
	return c.Mementos[i]
}

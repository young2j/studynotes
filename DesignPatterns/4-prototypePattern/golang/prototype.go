/*
 * File: prototype.go
 * Created Date: 2023-03-06 02:18:44
 * Author: ysj
 * Description:  原型类接口及其实现
 */

package main

import "fmt"

// 定义泥人儿，需要实现接口Clone
type ClayPerson interface {
	SetName(name string)
	SetTemper(temper string)
	SetEye(eye string)
	SetNose(nose string)
	SetSkin(Skin string)
	SetHands(hands int)
	SetGender(gender string)
	Introduction()
	Clone() ClayPerson
}

// 男娃娃小泥人儿
type ClayBoy struct {
	Name   string
	Eye    string
	Nose   string
	Skin   string
	Hands  int
	Gender string
	Temper string
}

// 克隆男娃娃小泥人儿
func (c *ClayBoy) Clone() ClayPerson {
	return &ClayBoy{
		Name:   c.Name,
		Eye:    c.Eye,
		Nose:   c.Nose,
		Skin:   c.Skin,
		Hands:  c.Hands,
		Gender: c.Gender,
		Temper: c.Temper,
	}
}

func (c *ClayBoy) SetName(name string) {
	c.Name = name
}
func (c *ClayBoy) SetTemper(temper string) {
	c.Temper = temper
}
func (c *ClayBoy) SetEye(eye string) {
	c.Eye = eye
}
func (c *ClayBoy) SetNose(nose string) {
	c.Nose = nose
}
func (c *ClayBoy) SetSkin(Skin string) {
	c.Skin = Skin
}
func (c *ClayBoy) SetHands(hands int) {
	c.Hands = hands
}
func (c *ClayBoy) SetGender(gender string) {
	c.Gender = gender
}

func (c *ClayBoy) Introduction() {
	fmt.Printf("%s: 性别%s, 有%s的眼睛, %s, %s的皮肤,%d只手,脾气%s。\n",
		c.Name, c.Gender, c.Eye, c.Nose, c.Skin, c.Hands, c.Temper,
	)
}

// 女娃娃小泥人儿
type ClayGirl struct {
	Name   string
	Eye    string
	Nose   string
	Skin   string
	Hands  int
	Gender string
	Temper string
}

// 克隆女娃娃小泥人儿
func (c *ClayGirl) Clone() ClayPerson {
	return &ClayGirl{
		Name:   c.Name,
		Eye:    c.Eye,
		Nose:   c.Nose,
		Skin:   c.Skin,
		Hands:  c.Hands,
		Gender: c.Gender,
		Temper: c.Temper,
	}
}

func (c *ClayGirl) SetName(name string) {
	c.Name = name
}
func (c *ClayGirl) SetTemper(temper string) {
	c.Temper = temper
}
func (c *ClayGirl) SetEye(eye string) {
	c.Eye = eye
}
func (c *ClayGirl) SetNose(nose string) {
	c.Nose = nose
}
func (c *ClayGirl) SetSkin(Skin string) {
	c.Skin = Skin
}
func (c *ClayGirl) SetHands(hands int) {
	c.Hands = hands
}
func (c *ClayGirl) SetGender(gender string) {
	c.Gender = gender
}

func (c *ClayGirl) Introduction() {
	fmt.Printf("%s: 性别%s, 有%s的眼睛, %s, %s的皮肤,%d只手,脾气%s。\n",
		c.Name, c.Gender, c.Eye, c.Nose, c.Skin, c.Hands, c.Temper,
	)
}

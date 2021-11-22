package main

import (
	"fmt"
)

type Animal struct {
	Name   string  //名称
	Color  string  //颜色
	Height float32 //身高
	Weight float32 //体重
	Age    int     //年龄
}

//奔跑
func (a Animal) Run() {
	fmt.Println(a.Name + "is running")
}

//吃东西
func (a Animal) Eat() {
	fmt.Println(a.Name + "is eating")
}

type Cat struct {
	a Animal
}

func main() {
	var c = Cat{
		a: Animal{
			Name:   "猫猫",
			Color:  "橙色",
			Weight: 10,
			Height: 30,
			Age:    5,
		},
	}
	fmt.Println(c.a.Name)
	c.a.Run()
}

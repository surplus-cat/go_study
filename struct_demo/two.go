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

type Lion struct {
	Animal // 匿名字段
}

//奔跑
func (a Animal) Run() {
	fmt.Println(a.Name + "is running")
}

//吃东西
func (a Animal) Eat() {
	fmt.Println(a.Name + "is eating")
}

func main() {
	var lion = Lion{
		Animal{
			Name:  "狮子",
			Color: "红色",
		},
	}

	lion.Run()

	fmt.Println(lion.Name)
}

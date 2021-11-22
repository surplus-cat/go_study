package main

import "fmt"

var name = "雷公"

func changeName() {
	name = "电母"

	fmt.Println("gogogo,", name)
}

func sayHi() {
	fmt.Println("hi,", name)
	go changeName() // 协程

	// go func() { // 匿名函数执行协程
	// 	name = "change"
	// }()

	// fmt.Println("gogogo,", name)
}

func main() {
	sayHi()
}

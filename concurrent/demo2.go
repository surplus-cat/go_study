package main

import "fmt"

var name = "yiyi"
var flag = make(chan bool) // 创建了bool类型的channel

func changeName() {
	name = "change"
	flag <- true // 发送
}

func sayHi() {
	go changeName() // 协程
	<-flag          //接收
	fmt.Println("hi", name)
}

func main() {
	sayHi() //
}

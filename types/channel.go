package main

import (
	"fmt"
)

func main() {
	// 通道
	var c chan int

	if c == nil {
		fmt.Println("通道c是 nil的，不能使用，需要先创建通道")
		c = make(chan int)
		fmt.Printf("c的数据类型是: %T", c)
	}
}

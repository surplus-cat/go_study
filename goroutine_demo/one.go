package main

import (
	"fmt"
)

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine ing...")
		c <- 666 //将666写入channel
	}()

	num := <-c // 从c中接受数据，并赋值给num

	fmt.Println("num 值为", num)
	fmt.Println("main gorotine 结束...")
}

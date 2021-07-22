package main;

import (
	"fmt"
	"time"
)

var sum = 0;

func main() {

	//开启100个协程来让 sum + 1
	for i := 1; i <= 100; i++ {
		go add();

		// 睡眠两秒防止程序提前退出
		time.Sleep(2 * time.Second)
		fmt.Println("sum:", sum);
	}
}

func add() {
	sum += 1;
}
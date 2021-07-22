package main

import (
	"fmt"
	"time"
	"sync"
)

var sum = 0;

var mutex = sync.Mutex{};

func main()  {
	//开启100个协程来让 sum + 1
	for i := 1; i < 10; i++ {
		go add();

		// 睡眠两秒防止程序提前退出
		time.Sleep(2 * time.Second)
		fmt.Println("sum:", sum)
	}
}

func add()  {
	mutex.Lock();
	defer mutex.Unlock(); //使用defer语句，确保锁一定会被释放
	sum += 1;
}
package main

import (
	"fmt"
	"sync"
	"time"
)

var sum = 0;
var mutex = sync.Mutex{}
var rwmutex = sync.RWMutex{}
func main() {
	//开启100个协程来让 sum + 1
	for i := 1; i <= 100; i++ {
		go add()
	}
	for i := 1; i<= 10; i++ {
		go fmt.Println("sum:",getSum())
	}
	// 睡眠两秒防止程序提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("sum:", sum)
}
func add(){
	mutex.Lock()
	defer mutex.Unlock() //使用defer语句，确保锁一定会被释放
	sum += 1
}
func getSum() int {
	rwmutex.RLock() //使用读写锁
	defer  rwmutex.RUnlock()
	return sum
}

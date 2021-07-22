package main

import (
	"fmt"
	"sync"
)

var sum = 0
var mutex = sync.Mutex{}
var rwmutex = sync.RWMutex{}

func run() {
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 1; i <= 100; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			add()
		}()
	}
	for i := 1; i <= 10; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("sum:", getSum())
		}()
	}
	//一直等待，只要计数器值为0
	wg.Wait()
}
func main() {
	run()
}
func add() {
	mutex.Lock()
	defer mutex.Unlock() //使用defer语句，确保锁一定会被释放
	sum += 1
}
func getSum() int {
	rwmutex.RLock() //使用读写锁
	defer rwmutex.RUnlock()
	return sum
}

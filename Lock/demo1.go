package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64

var lock sync.Mutex // 新增互斥锁

func f() {
	for i := 0; i < 100000; i++ {
		lock.Lock() // 加互斥锁
		x = x + 1
		lock.Unlock() // 解锁
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	fmt.Println(x)
}

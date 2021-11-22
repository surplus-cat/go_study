package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x1     int64
	wg1    sync.WaitGroup
	lock1  sync.Mutex
	rwlock sync.RWMutex
)

func main() {
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		wg1.Add(1)
		write()
	}

	for i := 0; i < 10000; i++ {
		wg1.Add(1)
		read()
	}
	wg1.Wait()
	fmt.Println(time.Now().Sub(startTime))
}

func read() {
	defer wg1.Done()
	lock1.Lock() // 互斥锁
	// rwlock.RLock() // 读写互斥锁
	fmt.Println(x1)
	lock1.Unlock() // 互斥锁
	// rwlock.RUnlock() // 读写互斥锁
}

func write() {

	defer wg1.Done()
	lock1.Lock() //
	x1 += 1
	lock1.Unlock() //
}

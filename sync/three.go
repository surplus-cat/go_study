package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {

	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(time.Millisecond)
	rwlock.Unlock() // 解锁

	wg.Done()
}

func read() {

	rwlock.Lock()                // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.Unlock()              // 解互斥锁

	wg.Done()
}

// 读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

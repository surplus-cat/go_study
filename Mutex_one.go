package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	// 主 goroutine 先获取锁
	fmt.Println("Locking  (G0)")
	mutex.Lock()
	fmt.Println("locked (G0)")

	wg.Add(3)

	for i := 0; i < 4; i++ {
		go func(i int) {

			fmt.Println("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Println("unlocked (G%d)\n", i)

			wg.Done()
		}(i)
	}

	// 主 goroutine 5 秒后释放锁
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")

	wg.Wait()
}

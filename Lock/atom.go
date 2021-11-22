package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var x int64

func main() {
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	fmt.Println(x)
}

func f() {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&x, 1)
	}
	wg.Done()
}

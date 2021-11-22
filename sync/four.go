package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(n int) {
			key := strconv.Itoa(i)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Println("k=%v, v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

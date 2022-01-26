package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go func(ch chan bool, t int) {
		time.Sleep(time.Second * time.Duration(t))
		ch <- true
	}(timeout, 2)

	select {
	case ms := <-ch1:
		fmt.Println("ch1 received data", ms)
	case ms2 := <-ch2:
		fmt.Println("ch2 received data", ms2)
	case <-timeout:
		fmt.Println("Timeout, exit")
	}
}

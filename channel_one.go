package main

import "fmt"

func Add(x, y int, ch chan int) {
	z := x + y
	ch <- z
}

func main() {

	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go Add(i, i, ch)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

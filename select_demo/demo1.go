package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "xiaomo"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received", msg2)
	default:
		fmt.Println("No data receive.")
	}
}

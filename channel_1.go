package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum // 把sum 发送到通道C
}

func main() {
	s := []int{2, 3, 4, 5, 7, 20}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)

}

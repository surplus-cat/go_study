package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go handlerTask1(&wg)
	go handlerTask2(&wg)
	go handlerTask3(&wg)

	wg.Wait()

	fmt.Println("全部任务执行完毕")
}

func handlerTask1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("执行任务 1")
}

func handlerTask2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("执行任务 2")
}

func handlerTask3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("执行任务 3")
}

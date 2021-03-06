package main

import (
	"time"
	"strconv"
	"fmt"
)

//go goroutine-基本介绍
//进程和线程介绍
//1.进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的最小单元
//2.线程是进程的一个执行实例，是程序执行的最小单元，它是比进程更小的能独立运行的最小单元
//3.一个进程可以创建和销毁多个线程，同一个进程可以有多个线程并发执行
//4.一个程序至少有一个进程，一个进程至少有一个线程
//并发和并行
//1.多线程程序在单核上运行，就是并发
//2.多线程程序在多核上运行，就是并行

//Go 协程和 Go 主线程
//Go 主线程(有程序员直接称为线程/也可以理解成进程): 一个 Go 线程上，可以起多个协程，你可以 这样理解，协程是轻量级的线程[编译器做优化]。

//Go 协程的特点
//1.有独立的栈空间
//2.共享程序堆空间
//3.调度由用户控制
//4.协程是轻量级的线程

//栗子
//1.在主线程(可以理解成进程)中，开启一个 goroutine, 该协程每隔 1 秒输出 "hello,world"
//2.在主线程中也每隔一秒输出"hello,golang", 输出 10 次后，退出程序
//3.要求主线程和 goroutine 同时执行.

func main()  {
	// 开启一个协程
	go test();

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second);
	}
}

func test()  {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello,world" + strconv.Itoa(i));
		time.Sleep(time.Second);
	}
}

//细节说明
//1.主线程是一个物理线程，直接作用在 cpu 上的。是重量级的，非常耗费 cpu 资源。
//2.协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小。
//3.Golang 的协程机制是重要的特点，可以轻松的开启上万个协程。其它编程语言的并发机制是一
//般基于线程的，开启过多的线程，资源耗费大，这里就突显 Golang 在并发上的优势了
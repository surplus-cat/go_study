package main

import (
	"fmt"
)

func main() {
	// 指针
	var i int = 20                   // 声明实际变量
	var ip *int                      // 声明指针变量
	ip = &i                          // 指针变量的存储地址
	fmt.Println("i变量的地址是: ", &i)     // 变量的存储地址
	fmt.Println("ip变量储存的指针地址: ", ip) // 指针变量的存储地址
	fmt.Println("*ip变量的值: ", *ip)    // 使用指针访问值
}

package main

import (
	"fmt"
)

func main() {
	//为了说明类型，我采用了显性的变量定义方法，实际开发中更多的是用“:=”自动获取类型变量类型
	var mystr string = "Hello!"
	var mystrPointer *string = &mystr

	fmt.Println(mystrPointer)

	fmt.Println(*mystrPointer)

}

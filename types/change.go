package main

import (
	"fmt"
)

func main() {
	// 类型转换
	var i16 int16 = 10000
	i32 := int32(i16)
	fmt.Print(i32)
}

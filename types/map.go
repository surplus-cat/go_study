package main

import (
	"fmt"
)

func main() {
	// Map
	personMap := make(map[string]string)
	personMap["zhangsan"] = "张三"
	personMap["lisi"] = "李四"
	personMap["wanger"] = "王二"
	personMap["zhaowu"] = "赵五"
	for p := range personMap {
		fmt.Println(p, "的中文名是", personMap[p])
	}

	fmt.Println(personMap)
}

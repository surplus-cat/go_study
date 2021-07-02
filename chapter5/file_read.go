package main

import (
	"io/ioutil"
	"fmt"
)

//ioutil.ReadFile一次性读取文件
func main()  {
	file := "/pic/240x120.png";

	content, err := ioutil.ReadFile(file);

	if err != nil {
		fmt.Println("read file err=%v", err)
	}

	fmt.Println("%v", string(content))
}

//细节说明
//不需要显式的open文件，close文件
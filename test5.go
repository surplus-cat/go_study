package main

import "fmt"

//定义一个结构体类型
type myStruct struct {
	Name string
}

//定义这个结构体的改名方法
func (m *myStruct) ChangeName(newName string) {
	m.Name = newName
}

func main() {
	//创建这个结构体变量
	mystruct := myStruct{
		Name: "zeta",
	}

	//调用改名函数
	mystruct.ChangeName("Chow")

	//没改成功
	fmt.Println(mystruct.Name)
}

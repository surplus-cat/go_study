package main

import (
	"fmt"
)

func main()  {
	name := "🐏";
	nameP := &name; // 取地址
	
	fmt.Println("name变量值为：", name)
	fmt.Println("name变量的内存地址为：", nameP);

	nameV := *nameP
	fmt.Println("nameP指针指向的值为:",nameV) //nameP指针指向的值为: 

	*nameP = "小行星";  //修改指针指向的值
	fmt.Println("nameP指针指向的值为:",*nameP)
	fmt.Println("name变量的值为:", name);

	modify(&name); // 传参传内存地址

	fmt.Println("name的值为:", name)
}

func modify(name *string)  {
	*name = "anywhere"; // 修改指针的值 
}
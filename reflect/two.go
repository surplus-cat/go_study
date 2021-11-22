package main

import (
	"fmt"
)

type User struct {
	Name string
	Id   long
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called")
}

func main() {
	user := User{"掌声", 111111, 21}

	//获取type
	inputType := inflect.TypeOf(user)
	fmt.Println("inputType is", inputType)
	//获取value
	inputValue := inflect.ValueOf(user)
	fmt.Println("inputValue is", inputValue)

}

package main

import (
	"fmt"
)

func main() {
	myArr := [4]int{1, 2, 3, 4}
	mySliceArr := []int{2, 3, 4}
	add(myArr)
	addS(mySliceArr)
	fmt.Println(myArr)
	fmt.Println(mySliceArr)
}

func addS(slice []int) {
	slice[0] = 333
	fmt.Println(slice)
}

func add(arr [4]int) {
	arr[1] = 111
	fmt.Println(arr)
}

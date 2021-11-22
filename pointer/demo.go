package main

import (
	"fmt"
)

func main() {
	a := 1
	changeValue(&a)
	changeVal(a)
	fmt.Println("a =", a)
}

func changeVal(b int) {
	b = 10
}

func changeValue(b *int) {
	*b = 10
}

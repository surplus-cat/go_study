package main

import "fmt"

func main() {
	var whatever [6]struct{}

	for i := range whatever {
		// defer func() {
		// 	fmt.Println(i)
		// }()

		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

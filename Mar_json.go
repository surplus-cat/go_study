package main

import (
	"encoding/json"
	"fmt"
)

// type Student struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }

// type Address struct {
// 	Road     string
// 	Street   string
// 	City     string
// 	Province string
// 	Country  string
// }

type Student map[string]interface{}
type Address map[string]interface{}

func main() {
	zhange3 := Student{
		"Name": "",
		"Age":  12,
		"Address": Address{
			"Road":     "renmin south road",
			"Street":   "123 street",
			"City":     "cs",
			"Province": "hn",
			"Country":  "CN",
		},
	}

	Info_of_zhang, err := json.Marshal(zhange3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(Info_of_zhang))
}

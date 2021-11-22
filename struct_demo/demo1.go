package main

import (
	"fmt"
)

type Member struct {
	name string
	age  int64
}

func Change(m1 Member, m2 *Member) {
	m1.name = "猎手"
	m2.name = "水手"
}

func main() {
	m1 := Member{}
	m2 := new(Member)
	Change(m1, m2)
	fmt.Println(m1, m2)
}

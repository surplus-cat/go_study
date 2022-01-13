package main

import (
	"YourProjectName/gin_demo/routers"
	"fmt"
	"gin_demo/routers"
)

func main() {
	r := routers.SetupRouter()

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

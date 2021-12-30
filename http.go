package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CodeInfo struct {
	Message string
	Code    string
	data    DataStruct
}

type DataStruct struct {
	total int
	list  ListStruct
}

type ListStruct struct {
	total int
	list  []ListObj
}

type ListObj struct {
	Title    string
	Path     string
	Id       string
	Icon     string
	hidden   string
	Pfather  string
	Ptype    string
	Children []ListObj
}

type Student struct {
	Name string `json:"name"` // 序列化 体现在输出的的时候 就变成小写了
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1 map

		data := map[string]interface{}{
			"msg":  "请求成功",
			"code": "10001200",
			"age":  20,
		}

		c.JSON(http.StatusOK, data)

	})

	r.GET("/anotherjson", func(c *gin.Context) {

		dataList := map[string]interface{}{
			"msg":  "请求成功",
			"code": "10001200",
			"age":  20,
		}
		//方法二 结构体
		data := &CodeInfo{
			Message: "wahahah",
			Code:    "10001200",
		}

		c.JSON(http.StatusOK, data)

	})

	// json序列化
	r.GET("/test", func(c *gin.Context) {
		data := &Student{
			Name: "xiaoming",
			Age:  23,
		}
		c.JSON(http.StatusOK, data)
	})

	/* 输出
	{"name":"xiaoming","age":23}
	*/

	r.Run(":7979")
}

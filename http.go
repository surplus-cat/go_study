package main

import (
	"fmt"
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

		// dataList := map[string]interface{}{
		// 	"msg":  "请求成功",
		// 	"code": "10001200",
		// 	"age":  20,
		// }
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

	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.GET("/topgoer", handle)

	r.Run(":7979")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "yang")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "doll")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello www.ccc.com",
	})
}

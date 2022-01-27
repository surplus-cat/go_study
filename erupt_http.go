package main

import (
	"fmt"
	"io/ioutil"

	// "encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/logs"
)

// type Result struct {
//   Msg string `json:"msg"`
//   Rescode string `json:"rescode"`
// }
// func parseBody(body io.Reader) {
// 	var v Result
// 	err := json.NewDecoder(body).Decode(&v)
// 	if err != nil {
// 		return nil, fmt.Errorf("DecodeJsonFailed:%s", err.Error())
// 	}
// }

// 请求函数
func onespider(num int, ch chan int) {
	url := "https://www.makeapie.com/chart/list?author=all&sortBy=star&builtinTags[]=category-work&pageSize=2&pageNumber" + strconv.Itoa(num)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("get error")
	}
	defer response.Body.Close()
	// reg, err1 := regexp.Compile(`<a href="/detail-\d{8}.html">(?s:(.*?))</a>`)
	// if err1 != nil {
	// 	log.Fatal("compile error")
	// }
	// body, err := ioutil.ReadAll(response.Body)

	// str := make([]byte, 2048)
	// size, err := response.Body.Read(str)
	// if size == 0 {
	// 	panic(err)
	// }
	log.Println(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error("ioutil.ReadAll failed ,err:%v", err)
	}
	// fmt.Print(string(body))
	// jsonStu, err := json.Marshal(body)
	// if err != nil {
	// 	fmt.Println("生成json字符串错误")
	// }
	// fmt.Print(string(jsonStu))

	// //jsonStu是[]byte类型，转化成string类型便于查看
	// fmt.Println(jsonStu, string(jsonStu))
	// fmt.Fprintf(body)
	defer response.Body.Close()

	os.Create("my.json")

	file, err := os.OpenFile("my.json", os.O_RDWR, 0666)
	if err != nil {
		println("error occuer")
		panic(err)
	} else {
		println("write data ok")
		fmt.Printf("第%d条数据写入完毕\n", num)
		file.Write(body)
	}
	defer file.Close()

	// var respstring string
	// buf := make([]byte, 1024)
	// for {
	// 	num, _ := resp.Body.Read(buf)
	// 	if num == 0 {
	// 		break
	// 	}
	// 	respstring += string(buf[:num])
	// }

	// cont := reg.FindAllStringSubmatch(respstring, -1)
	// file, _ := os.OpenFile("./爬虫/"+"第"+strconv.Itoa(num)+"页爬虫.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)

	// defer file.CLose()
	// var i int
	// for _, value := range cont {
	// 	if len(value[1]) < 100 {
	// 		continue
	// 	}
	// 	value[1] = strings.Replace(value[1], "<br />", "\n", -1)
	// 	index := strconv.Itoa(i + 1)
	// 	file.Write([]byte("第" + index + "则段:\n" + value[1] + "\n\n\n"))
	// 	i++
	// }
	ch <- num
}
func Spider(start, end int) {
	ch := make(chan int)
	for i := start; i <= end; i++ {
		time.Sleep(time.Second * 1)
		go onespider(i, ch)
	}
	// for i := s; i <= e; i++ {
	// 	num := <-ch
	// 	fmt.Printf("第%d页爬取完毕\n", num)
	// }
}
func main() {
	var start int = 0
	var end int = 1 //1000000
	Spider(start, end)
}

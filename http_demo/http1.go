package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 绑定 url 和对应的处理方法
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// 在9666端口监听
	log.Fatal(http.ListenAndServe(":9666", nil))
}

// 访问 / 的处理方法
func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// 访问 /hello 的处理方法
func helloHandler(w http.ResponseWriter, req *http.Request) {

	for k, v := range req.Header {
		_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9777", engine))
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

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		indexHandler(w, req)
	case "/hello":
		helloHandler(w, req)
	default:
		_, err := fmt.Fprintf(w, "404 Not Found, %s\n", req.URL)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{}) //使用handle  需要implement  http.Handler.ServeHTTP()
	mux.HandleFunc("/hello", sayHello)

	//实现静态文件的访问
	wd, err := os.Getwd() //get work direct 获得工作目录
	if err != nil {
		log.Fatal(err)
	}
	//StripPrefix脱去前缀
	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //implement  http.Handler.ServeHTTP() method
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version 2.")
}

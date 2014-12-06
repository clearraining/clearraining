package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //return ServeMux

	//方式1
	mux.Handle("/", &myHandler{})

	//方式2
	mux.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", mux)

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

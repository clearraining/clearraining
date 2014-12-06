package gotest

import (
	"fmt"
)

func Hello() {
	fmt.Println("hello from gotest!")
	hi()
}

func hi() {
	fmt.Println("hi from gotest")
}

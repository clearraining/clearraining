package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("go go go")
		c <- true
		close(c)
	}()
	for v := range c { //遍历c
		fmt.Println(v)
	}
}

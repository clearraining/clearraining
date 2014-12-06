package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!")
		c <- false //true也可以
		fmt.Println("end")
	}()
	<-c

}

package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("o1")
					o <- true
					break
				}
				fmt.Println("c1:", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("o2")
					o <- true
					break
				}
				fmt.Println("c1:", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "int"
	c1 <- 2
	c2 <- "string"

	close(c1)
	close(c2)
	<-o

}

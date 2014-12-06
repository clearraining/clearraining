package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 2)
	go func() {
		fmt.Println("go go go")
		<-c
	}()
	c <- false //由于有缓冲2 所以c中有内容小于2的时候就会直接执行下一步   不会阻塞等待。所以没有输出
}

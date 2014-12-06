package main

import (
	"fmt"
)

var (
	Id   string
	Flag *bool
	Name string
)

func main() {
	flage := true
	Flag = &flage
	Id = "123455"
	Name = "xia"
	if *Flag {
		fmt.Println(Id, Name)
	}
}

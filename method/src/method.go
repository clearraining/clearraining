package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type TZ int
type SS int
type TT int

func main() {
	// a1 := A{}
	var a1 A
	a1.Print()
	fmt.Println(a1.Name)

	b := B{}
	b.Print()

	var tz TZ //类型别名
	tz.Print()
	(*TZ).Print(&tz)

	var ss SS
	ss = 33
	fmt.Println(ss)

	//0变换成100
	var tt TT
	tt.Increace(100)
	fmt.Println(tt)
}

func (a *A) Print() {
	a.Name = "naizui"
	fmt.Println("A")
}

func (b B) Print() {
	fmt.Println("B")
}

func (tz *TZ) Print() {
	fmt.Println("TZ")
}

func (tt *TT) Increace(num int) {
	*tt += TT(num)
}

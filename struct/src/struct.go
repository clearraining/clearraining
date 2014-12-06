package main

import (
	"fmt"
)

type human struct {
	Sex string
}

type teacher struct {
	human
	Name string
	Age  int
}

type person struct {
	Name string
	Age  int
}

//嵌套struct
type student struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func main() {
	lisi := person{
		Name: "lisi",
	}
	lisi.Age = 22
	fmt.Println(lisi)

	wangwu := person{
		Name: "wangwu",
		Age:  23,
	}
	changeAge(&wangwu) //指针
	fmt.Println("wangwu_Age:", wangwu.Age)

	zhangsan := &person{
		Name: "zhangsan",
		Age:  33,
	}
	zhangsan.Name = "张三哥" // 可以不需要  *zhangsan
	fmt.Println(zhangsan.Name)
	changeName(zhangsan)
	fmt.Println(zhangsan.Name)

	//匿名结构
	niming := struct {
		Name string
		Age  int
	}{
		Name: "niming",
		Age:  12,
	}
	fmt.Println("niming:", niming)

	//嵌套struct
	naizui := student{
		Name: "naizui",
		Age:  23,
	}
	naizui.Contact.Phone = "1233445667" //不能写在里面
	naizui.Contact.City = "fuding"
	fmt.Println(naizui)

	var raining student
	raining = naizui
	fmt.Println("raining:", raining)

	clearraining := student{}
	clearraining = raining
	fmt.Println("clearraining:", clearraining)

	//嵌入struct
	jiangming := teacher{
		Name:  "jiangming",
		Age:   22,
		human: human{Sex: "man"},
	}
	fmt.Println("jiangming:", jiangming)
	jiangming.human.Sex = "woman"
	fmt.Println("jiangming:", jiangming)

}
func changeAge(per *person) {
	per.Age = 24
}

func changeName(per *person) {
	per.Name = "张三"
}

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("hello world.")
}

func main() {
	user := User{1, "zhangsan", 33}
	user_1 := User{2, "lisi", 22}
	Info(user)
	Info(&user_1)

	manager := Manager{User: User{3, "wangwu", 44}, title: "123"}
	m := reflect.TypeOf(manager)
	fmt.Printf("%#v\n", m.FieldByIndex([]int{0, 1})) //得到user的第二个字段名Name

	/*for i := 0; i < m.NumField(); i++ {
		f := m.Field(i)
		val := m.FieldByIndex([]int{0, 0})
		fmt.Println(f.Name, f.Type, val)
	}*/

}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)

	if k := t.Kind(); k != reflect.Struct { //对传进来的类型进行判断，如果不是struct返回
		fmt.Println("XX")
		return
	}

	fmt.Println("Type:", t.Name())
	fmt.Println("Field:")

	for i := 0; i < v.NumField(); i++ { //t.NumField()  也一样
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

}

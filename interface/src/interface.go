package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect
}

type Connect interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

//只要在类型中使用了接口的所有方法就是实现了接口
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect(usb USB) { //USB 可以修改为interface{}
	/*
		if pc, ok := usb.(PhoneConnecter); ok { //return pc ok
			fmt.Println("disconnect:", pc.name)
			return
		}
		fmt.Println("unknown!")
	*/
	switch v := usb.(type) { //让系统自己判断类型
	case PhoneConnecter:
		fmt.Println("disconnect:", v.name)
	default:
		fmt.Println("unknown!")
	}

}
func main() {
	var usb USB
	usb = PhoneConnecter{"PhoneConnecter"}
	usb.Connect()
	Disconnect(usb)

	usb_1 := PhoneConnecter{"phoneConnecter_usb_1"}
	Disconnect(usb_1)

	var a Connect
	a = Connect(usb_1)
	a.Connect()

}

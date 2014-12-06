package main

import (
	//"fmt"
	//"github.com/clearraining/gotest"
	"github.com/astaxie/beego"
	//"time"
)

type HomeControl struct {
	beego.Controller //包括了get post等方法
}

func (this *HomeControl) Get() { //隐藏了beego.Controller中的get post等方法
	this.Ctx.WriteString("hello clearraining")
}

func main() {
	beego.Router("/", &HomeControl{})
	beego.Run()

	/*gotest.Hello()
	gotest.Hello()
	m := time.Now()
	fmt.Println(m.Format(time.ANSIC))

	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	select {}*/
}

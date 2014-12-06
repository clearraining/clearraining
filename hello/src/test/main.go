package main

// godoc -http=:8080
// localhost:8080
import (
	"fmt"
	"sort"
	"strconv"
)

//定义常量
const (
	w = iota
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PI   = 3.4
	age  = 33
	name = "naizui"
	z    = len(name)
)

//类型别名
type (
	byte     int8
	rune     int32
	Bytesize int64
	文本       string
)

func main() {
	fmt.Println("Hello, GO!下雨了！")
	fmt.Println(age)
	fmt.Println("z:", z)
	fmt.Println("B KB", B, KB)

	//初始零值   0  false  空值
	var a int
	var a1 [2]int

	var b bool
	var b1 [2]bool

	var c string
	var c1 [2]string

	fmt.Println(a, a1, b, b1, c, c1)

	//变量申明
	var d int
	d = 2
	var e int = 2
	var f = 2
	g := 2
	fmt.Println(d, e, f, g)

	var h, i, j, k int = 2, 3, 4, 5
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	//类型转换
	var l float32 = 100.1
	m := int(l)
	fmt.Println(m)

	var n int = 65
	o := string(n)
	fmt.Println(o) //A

	o = strconv.Itoa(n) // 65
	fmt.Println(o)

	fmt.Println(1 << 10 << 10)
	fmt.Println(1 << 20)
	fmt.Println(1 >> 10)
	/*
	   位运算符
	   6: 0110
	   11:1011
	   ------------------
	   &  0010   2
	   |  1111   15
	   ^  1101   13    只有一个是1的就是1
	   &^ 0100   4      第二个中是1的就把第一个变成0
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	a = 0
	if a != 0 && (10/a) > 1 {
		fmt.Println("OK!")
	} else {
		fmt.Println("NO!")
	}

	a = 1
	var p *int = &a
	var r *int
	r = &a
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(*r)

	for {
		a++
		if a > 3 {
			fmt.Println(a)
			break
		}
	}
	fmt.Println("over")

	a = 1
	for a < 3 {
		a++
		fmt.Println(a)
	}
	fmt.Println("over")

	a = 1
	for i := 1; i < 4; i++ {
		a++
		fmt.Println(a)
	}

LABEL:
	for i := 1; i < 4; i++ {
		for {
			fmt.Println("i:", i)
			continue LABEL
			//goto  LABEL 调整循环位置
			//break LABEL  //如果没有加Label，跳出里面一个for循环
		}
	}

	a = 1
	switch a {
	case 0:
		fmt.Println("case 0")
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}

	switch {
	case a == 1:
		fmt.Println("case 1")
		fallthrough
	case a > 1:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}

	switch a := 1; {
	case a == 1:
		fmt.Println("case 1")
	}

	var v [3]int
	fmt.Println(v)

	x := [20]int{15: 1} //第15个元素为1 其他为0
	fmt.Println(x)

	q := [...]int{1, 2, 3, 4, 5}
	fmt.Println(q)

	var aq *[5]int = &q
	aqq := &q
	fmt.Println(*aq)
	fmt.Println(*aqq)

	ax, ay := 1, 2
	axy := [...]*int{&ax, &ay}
	fmt.Println(axy)

	//指针数组
	var xy [3]int
	xy[1] = 1
	fmt.Println(xy)

	var xp = new([3]int)
	xp[1] = 1
	fmt.Println(xp)

	ab := [...][3]int{
		{1, 2, 3},
		{2: 1}}
	fmt.Println(ab)

	pai := [...]int{3, 5, 7, 2, 1, 9, 10, 8}
	fmt.Println(pai)

	num := len(pai)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if pai[i] < pai[j] {
				temp := pai[i] + pai[j]
				pai[i] = temp - pai[i]
				pai[j] = temp - pai[i]
			}
		}
	}
	fmt.Println(pai)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sli := arr[5:10]
	fmt.Println(sli)
	sli = arr[5:len(arr)]
	fmt.Println(sli)
	sli = arr[5:]
	fmt.Println(sli)

	strsli := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	strsli_a := strsli[2:5]
	fmt.Println("byte:", strsli_a, len(strsli_a), cap(strsli_a))
	strsli_b := strsli_a[3:5] //超出srtsli_a范围，从strsli取  共享地址
	fmt.Println(strsli_b)

	strappend := make([]int, 3, 6)
	strappend = append(strappend, 1, 2, 3)
	fmt.Printf("value:%v\naddress:%p\n", strappend, strappend)
	strappend = append(strappend, 1, 2, 3)
	fmt.Printf("value:%v\naddress:%p\n", strappend, strappend) //地址改变

	sli_a := []int{1, 2, 3, 4, 5}
	sli_b := sli_a[2:5]
	sli_c := sli_a[3:5]
	sli_c = append(sli_c, 1, 1, 1, 1, 1, 1)
	sli_b[1] = 9
	fmt.Println(sli_b, sli_c)

	copy(sli_b, sli_c)
	fmt.Println(sli_b, sli_c)

	//map
	var map_1 map[int]string
	map_1 = map[int]string{}
	fmt.Println(map_1)

	var map_2 map[int]string = map[int]string{}
	fmt.Println(map_2)

	var map_3 map[int]string = make(map[int]string)
	fmt.Println(map_3)

	map_3_1 := make(map[int]string)
	fmt.Println(map_3_1)

	map_4 := map[int]string{}
	map_4[2] = "OK"
	fmt.Println("OK", map_4)

	delete(map_4, 2)
	fmt.Println(map_4)

	// var map_5 map[int]map[int]string
	// map_5 = make(map[int]map[int]string)
	map_5 := map[int]map[int]string{}
	map_a, ok := map_5[2][1] //返回ok为false
	if !ok {
		map_5[2] = make(map[int]string)
	}
	map_5[2][1] = "GOOD"
	map_a = map_5[2][1]
	fmt.Println(map_a, map_5, ok)

	map_6 := make([]map[int]string, 5)
	for i := range map_6 {
		map_6[i] = make(map[int]string, 1)
		map_6[i][1] = "OK"
	}
	fmt.Println(map_6)

	/**
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	var tags map[string]interface{}
	tags = make(map[string]interface{})

	var tagsLocal map[string]string
	tagsLocal = make(map[string]string)
	var tagsTest map[string]string
	tagsTest = make(map[string]string)
	var tagsProduction map[string]string
	tagsProduction = make(map[string]string)

	tagsLocal["dev.www.9178.us"] = "dev.www.9178.us"
	tagsLocal["dev.static.9178.us"] = "dev.static.9178.us"

	tagsTest["dev.www.9178.us"] = "www.ninja911.com"
	tagsTest["dev.static.9178.us"] = "static.ninja911.com"

	tagsProduction["dev.www.9178.us"] = "ipx.www.ninja911.com"
	tagsProduction["dev.static.9178.us"] = "ipx.static.ninja911.com"

	tags["local"] = tagsLocal
	tags["test"] = tagsLocal
	tags["production"] = tagsLocal

	fmt.Println("tags:", tags)
	fmt.Println(len(tags))

	map_7 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	map_s := make([]int, len(map_7))
	i = 0
	for k, _ := range map_7 { //k = key    value
		map_s[i] = k
		i++
	}
	sort.Ints(map_s)
	fmt.Println("map_s:", map_s)

	map_8 := map[int]string{1: "a", 2: "b", 3: "c"}
	map_9 := map[string]int{}
	for k, v := range map_8 {
		map_9[v] = k
	}
	fmt.Println(map_8, "\n", map_9)

	//func 实现直接传值
	fun_1, fun_2 := 1, 2
	A(fun_1, fun_2)
	fmt.Println("fun_1_2:", fun_1, fun_2)

	//引用传递
	fun_3 := []int{1, 2, 3, 4, 5}
	li(fun_3)
	fmt.Println("fun_3:", fun_3)

	//指针传递
	fun_4 := 2
	ki(&fun_4)
	fmt.Println(fun_4)

	fun_5 := ji //ji is a func
	fun_5()

	fun_6 := func() { //内嵌函数
		fmt.Println("fun_6")
	}
	fun_6()

	//闭包
	clo := closure(10)
	fmt.Println(clo(2))
	fmt.Println(clo(3))

	for i := 0; i < 3; i++ {
		//defer fmt.Println(i) //defer print 2 1 0  顺序相反
		defer func() {
			fmt.Println("defer:", i) //print 3 3 3
		}()
	}

	pan_1()
	pan_2()
	pan_3()

}

func A(s ...int) {
	s[0] = 3
	s[1] = 4
	fmt.Println("s:", s)
}

func li(sli []int) {
	sli[0] = 11
	sli[3] = 23
	fmt.Println("sli:", sli)
}

func ki(s *int) {
	*s = 22
	fmt.Println(*s)
}

func ji() {
	fmt.Println("func_5")
}

// 闭包
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func pan_1() {
	fmt.Println("pan_1")
}

func pan_2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("pan_2_1")
		}
	}()
	panic("panic in pan_2") //不执行
	fmt.Println("pan_2_2")
}

func pan_3() {
	fmt.Println("pan_3")
}

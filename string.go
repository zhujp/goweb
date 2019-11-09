package main

import (
	"fmt"
)

var s string = "hello"

var (
	i int
	b string
) //分组定义

func main() {
	//	s[0] = "c"  //字符串不可变，编译会产生错误
	c := []byte(s) //字符串数组
	//	fmt.Println(c)  //[104 101 108 108 111]
	//	fmt.Println(string(c[0]))  //h
	c[0] = 'c' //byte 用单引号，byte，rune比较特殊,byte 等同于int8，常用来处理ascii字符,rune 等同于int32,常用来处理unicode或utf-8字符
	fmt.Println(string(c))
	//	s2 := string(c)
	//	fmt.Println(s2)
}

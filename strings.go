package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("My name is vilay", "vil")) // 字符串 s 中是否包含 substr，返回 bool 值 return true
	fmt.Println(strings.Contains("My name is vilay", "ci"))  //  return false
	fmt.Println(strings.Contains("vilay", ""))               //return true
	fmt.Println(strings.Contains("", ""))                    //return true
	s := []string{"vilay", "cicy", "didy"}
	fmt.Println(strings.Join(s, ","))                          // 字符串链接，把 slice a 通过 sep 链接起来 类似于php implode
	fmt.Println(strings.Index("vilay", "i"))                   // 在字符串 s 中查找 sep 所在的位置，返回位置值，找不到返回 -1return 1
	fmt.Println(strings.Index("vilay", "d"))                   //return -1
	fmt.Println("vilay repeat:" + strings.Repeat("vilay", 2))  // 重复 s 字符串 count 次，最后返回重复的字符串return vilay repeat:vilayvilay
	fmt.Println(strings.Replace("vilay vilay", "la", "lb", 1)) //在 s 字符串中，把 old 字符串替换为 new 字符串，n 表示替换的次数，小于 0 表示全部替换
	fmt.Println(strings.Replace("vilay vilay", "la", "lb", -1))
	fmt.Println(strings.Split("vilay vilay vilay", " ")) //把 s 字符串按照 sep 分割，返回 slice
	fmt.Println(strings.Trim("!!vilay!a!", "!"))         //在 s 字符串的头部和尾部去除 cutset 指定的字符串
	fmt.Println(strings.Fields(" vilay is a name "))
	fmt.Println(strings.Fields("  foo bar  baz   "))
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}

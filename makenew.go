package main

import "fmt"

//make 用于内建类型（map、slice 和 channel）的内存分配。new 用于各种类型的内存分配。
//new 返回指针。
//make 返回初始化后的（非零）值。
//关于 “零值”，所指并非是空值，而是一种 “变量未填充前” 的默认值，通常为 0。
//int     0
//int8    0
//int32   0
//int64   0
//uint    0x0
//rune    0 // rune 的实际类型是 int32
//byte    0x0 // byte 的实际类型是 uint8
//float32 0 // 长度为 4 byte
//float64 0 // 长度为 8 byte
//bool    false
//string  ""
func main() {

}

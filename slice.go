package main

import "fmt"

//slice 并不是真正意义上的动态数组，而是一个引用类型。slice 总是指向一个底层 array，slice 的声明也可以像 array 一样，只是不需要长度。
// slice 的index只能是int

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//	var a, b []int
	//	a = arr[0:3]
	//	b = arr[5:8]
	c := arr[2:4]
	fmt.Println(c)
}

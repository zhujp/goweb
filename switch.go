package main

import "fmt"

func main() {
	x := 6
	//go switch 默认自带了break 匹配了就会终止
	//go 可以使用fallthrough，强制继续执行后面的逻辑
	switch x {
	case 1:
		fmt.Println(" x is 1")
	case 6:
		fmt.Println("x is 6")
		fallthrough
	default:
		fmt.Println("default x is %d", x)
	}
}

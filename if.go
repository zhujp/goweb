package main

import "fmt"

func computeValue() int {
	return 5
}

func main() {
	//	x := 10
	//	if x > 5 {
	//		fmt.Println("x is more five")
	//	} else {
	//		fmt.Println("x is less ten")
	//	}
	//	if x := computeValue(); x > 10 { //if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，
	//		fmt.Println("x is grater ten")
	//	} else {
	//		fmt.Println("x is less ten")
	//	}

	x := 10
	if x > 10 {
		fmt.Println("x is grate 10")
	} else if x == 10 {
		fmt.Println("x is equal 10")
	} else {
		fmt.Println("x is less 10")
	}
}

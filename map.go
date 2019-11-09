package main

import "fmt"

//map 是无序的，每次打印出来的 map 都会不一样，它不能通过 index 获取，而必须通过 key 获取
//map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型
//内置的 len 函数同样适用于 map，返回 map 拥有的 key 的数量
//map 的值可以很方便的修改，通过 numbers["one"]=11 可以很容易的把 key 为 one 的字典值改为 11
//map 和其他基本型别不同，它不是 thread-safe，在多个 go-routine 存取时，必须使用 mutex lock 机制
func main() {

	//	var numbers map[string]int
	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 2

	fmt.Println(numbers)
	delete(numbers, "one")
	fmt.Println(numbers)
}

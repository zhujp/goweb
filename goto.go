package main

import "fmt"

//用 goto 跳转到必须在当前函数内定义的标签,标签名是大小写敏感的
func main() {
	i := 10
Here:
	fmt.Println(i)
	i++
	if i < 20 {
		goto Here
	}

}

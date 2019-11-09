package main

import (
	"fmt"
)

//iota，这个关键字用来声明 enum 的时候采用，它默认开始值是 0，const 中每增加一行加 1： 每增加一行+1

//const (
//	x = iota //0
//	y = iota //1
//	z = iota / 2
//)

//const (
//	x, y, z = iota, iota, iota  //同一行值不变 都是0
//)

//const (
//	x = iota
//	y = iota
//	z //不写默认与前一个相同，即z=iota
//)

const (
	x = iota
	y = "vilay"
	z = iota //每一行会增加1，所以这边是2
)

func main() {
	fmt.Println(x, y, z)
}

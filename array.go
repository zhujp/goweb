package main

import "fmt"

//arr := [3]int{1,2,4} //这种写法只能在函数内部
func main() {

	//	var arr [10]int //默认初始值为0 [0 0 0 0 0 0 0 0 0 0]
	//	arr[5] = 6      //索引从0开始
	//	arr := [3]int{1, 2, 4} //初始化
	//	arr := [10]int{1, 2, 5} //没有填充的默认是0
	//	arr := [...]int{1, 2, 4, 5, 6} //不定义长度，go自动根据数组长度计算
	//二维数组
	//	arr := [3][2]string{[2]string{"vilay", "hello"}, [2]string{"vilay1", "hello2"}} //二维数组
	//	arr := [3][2]string{{"vilay", "hello"}, {"vilay1", "hello2"}}
	fmt.Println(arr)
}

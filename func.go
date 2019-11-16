package main

import "fmt"

//关键字 func 用来声明一个函数 funcName
//函数可以有一个或者多个参数，每个参数后面带有类型，通过 , 分隔
//函数可以返回多个值
//上面返回值声明了两个变量 output1 和 output2，如果你不想声明也可以，直接就两个类型
//如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号，函数如果有返回值，必须申明返回值类型，返回变量可省略
//如果没有返回值，那么就直接省略最后的返回信息
//如果有返回值， 那么必须在函数的外层添加 return 语句
//

func computeValue(x int, y int) int {
	return x + y
}

//申明了返回值变量，在函数里面直接return 就默认直接把值返回,同时在函数内部不要再申明该变量，比如不需要 name := username,直接使用name = username
func getUsernameAges(username string, age int) (name string, old int) {
	name = username
	old = age
	//	return
	return name, age
}

func getUsernameAge(username string, age int) (string, int) {
	name := username
	old := age
	//	return
	return name, old
}

//arg ...int 告诉 Go 这个函数接受不定数量的参数。注意，这些参数的类型全部是 int。在函数体中，变量 arg 是一个 int 的 slice：
func changeParamters(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
}

//普通传值
func getSum(a int) int {
	return a + 1
}

//传指针
func getPSum(a *int) int {
	*a = *a + 1
	return *a
}
func main() {
	//	sum := computeValue(5, 6)
	//	fmt.Println(sum)
	//	username, age := getUsernameAge("vilay", 20)
	//	fmt.Println("username is %s", username)
	//	fmt.Println("age is %d", age)
	//	changeParamters(1, 2, 3, 5)
	a := 5
	//	sum := getSum(a) //函数普通值传递不会改变a本身的值
	//	fmt.Println(sum)
	//	fmt.Println(a)
	sum := getPSum(&a)
	fmt.Println(sum)
	fmt.Println(a)
}

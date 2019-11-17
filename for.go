package main

import "fmt"

func main() {
	//	sum := 0
	// Go 里面没有 , 操作符，那么可以使用平行赋值 i, j = i+1, j-1
	//	for i := 0; i < 10; i++ {
	//		sum += i
	//	}
	//	for sum < 10 {
	//		sum += sum
	//	}
	//	fmt.Println("sum is %d", sum)
	//在循环里面有两个关键操作 break 和 continue, break 操作是跳出当前循环，continue 是跳过本次循环。当嵌套过深的时候， break 可以配合标签使用，即跳转至标签所指定的位置
	usernames := make(map[string]string)
	usernames["one"] = "vilaye"
	usernames["two"] = "vstary"
	usernames["three"] = "vkll"
	//	for k, v := range usernames {
	//		fmt.Println("k is %s", k)
	//		fmt.Println("v is %s", v)
	//	}
	//	for k, _ := range usernames {
	//		fmt.Println("k is %s", k)

	//	}

	for i := range usernames {
		fmt.Println(i)
	}

}

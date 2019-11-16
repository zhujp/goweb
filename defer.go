package main

import "fmt"

//延迟（defer）语句，你可以在函数中添加多个 defer 语句。当函数执行到最后时，这些 defer 语句会按照逆序执行，最后该函数返回。
//特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，
//在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。如下代码所示，我们一般写打开一个资源是这样操作的：

func testDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("value is %d", i)
	}
}

func main() {
	testDefer()
}

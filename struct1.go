package main

import "fmt"

//匿名字段
type Human struct {
	name   string
	age    int
	weight float32
	phone  string
}

type Person struct {
	Human   //匿名字段
	special string
	phone   string
}

//如果 human 里面有一个字段叫做 phone，而 student 也有一个字段叫做 phone，那么该怎么办呢？
//Go 里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过 student.phone 访问的时候，是访问 student 里面的字段，而不是 human 里面的字段。
func main() {

	P := Person{Human{"vilay", 23, 66.3, "1895919"}, "is handsome", "1565"}
	//	fmt.Println(P.special)
	//	fmt.Println(P.Human.name)
	//	fmt.Println(P.name)
	fmt.Println(P.phone)

}

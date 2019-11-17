package main

import "fmt"

type person struct {
	name string
	age  int
}

//定义tag，健值对的形式 uname:name,uage:age
//type person struct {
//	name string `uname:"name"`
//	age  int    `uage:"age"`
//}

func main() {
	//	var P person
	//	P.name = "vilay"
	//	P.age = 20
	//	P := person{"vilays", 30}
	P := new(person) //P 是指针
	//	P := person{name: "vilays", age: 34}
	fmt.Println(P.name)
}

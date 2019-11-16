package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//	var P person
	//	P.name = "vilay"
	//	P.age = 20
	//	P := person{"vilays", 30}
	P := new(person) //P 是指针
	//	P := person{name: "vilays", age: 34}
	fmt.Println(P.name)
}

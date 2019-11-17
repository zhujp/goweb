package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
}

type Employee struct {
	Human   // 匿名字段
	company string
}

// 在 human 上面定义了一个 method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Employee 的 method 重写 Human 的 method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	mark := Student{Human{"vilay", 25, "6666"}, "MIT"}
	sam := Employee{Human{"vilays", 45, "4555"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}

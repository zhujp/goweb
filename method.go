package main

import (
	"fmt"
	"math"
)

//func (r ReceiverType) funcName(parameters) (results)
//虽然 method 的名字一模一样，但是如果接收者不一样，那么 method 就不一样
//method 里面可以访问接收者的字段
//调用 method 通过 . 访问，就像 struct 里面访问字段一样

type Rectangle struct {
	width, height float64
}

type Circle struct {
	redius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.redius * c.redius * math.Pi
}
func main() {

	r := Rectangle{4, 5}
	c := Circle{3}
	fmt.Println(r.area())
	fmt.Println(c.area())

}

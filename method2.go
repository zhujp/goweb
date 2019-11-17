package main

import (
	"fmt"
)

//如果一个 method 的 receiver 是 *T, 你可以在一个 T 类型的实例变量 V 上面调用这个 method，而不需要 &V 去调用这个 method
//如果一个 method 的 receiver 是 T，你可以在一个 T 类型的变量 P 上面调用这个 method，而不需要 P 去调用这个 method

const (
	WHITE = iota //每一行会自增
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte //color byte 的别名

type Box struct {
	width, height, depth float64
	color                Color
}

var Boxlists []Box //box slice

//求体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

//设置颜色
func (b *Box) SetColor(c Color) {
	b.color = c
}

//体积最大的盒子的体积
func (bl Boxlists) BiggestColor() Color {
	volume := 0
	boxColor := Color(WHITE)

	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			volume = bv
			boxColor = b.color
		}
	}

	return boxColor
}

//全部设置一个黑色
func (bl Boxlists) SetItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	colors := []string{"black", "white", "blue", "red", "yellow"}

	return colors[c]
}

func main() {
	boxes := BoxLists{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.SetItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())

}

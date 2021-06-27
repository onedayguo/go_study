package main

import (
	"fmt"
)

type Vertex struct {
	x int
	Y int
}

func (v *Vertex) getX() int {
	fmt.Println("指针接收者：%p", &v)
	return v.x
}

func (v *Vertex) setX(value int) {
	fmt.Println("指针接收者：%p", &v)
	v.x = value
}

func (v Vertex) setXX(value int) {
	fmt.Println("值接收者：%p", &v)
	v.x = value
}

func main1() {

}

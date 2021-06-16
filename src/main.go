package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	x int
	Y int
}

func (this *Vertex) getX() int {
	return this.x
}

func main() {
	var str string = "hello,world"
	fmt.Println(reflect.TypeOf(str[1]))
	fmt.Println('E', 'e')
}

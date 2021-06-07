package main

import "fmt"

type Vertex struct {
	x int
	Y int
}

func (this *Vertex) getX() int {
	return this.x
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.getX())
}

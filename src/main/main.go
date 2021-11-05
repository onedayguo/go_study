package main

import (
	"fmt"
	"strings"
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

func main() {
	b := "现有 %s 机构申请与  绑定，合同已经被法务驳回 %s 次，请跟进"

	//切分字符串
	str2 := strings.Split(b, "%s")

	fmt.Sprintf(string(len(str2)))
}

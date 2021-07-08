package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
	file, err := os.Create("seo_result.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Comma = ','
	headline := []string{"URL", "Title", "Keywords", "Description"}
	writer.Write(headline)

	record1 := []string{"http://baidu.com", "百度", "搜索一下，你就知道", "查询百科信息"}
	record2 := []string{"http://baidu.com", "百度", "搜索一下，你就知道", "查询百科信息"}
	record3 := []string{"http://baidu.com", "百度", "搜索一下，你就知道", "查询百科信息"}
	writer.WriteAll([][]string{record1, record2, record3})

	writer.Flush()
}

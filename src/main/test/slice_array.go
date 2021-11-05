package test

import (
	"fmt"
	"strings"
)

func main1() {

	b := "现有 %s 机构申请与  绑定，合同已经被法务驳回 %s 次，请跟进"

	//切分字符串
	str2 := strings.Split(b, "%s")

	fmt.Sprintf(string(len(str2)))
}

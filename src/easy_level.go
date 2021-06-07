package main

import "fmt"

// 1837. Sum of Digits in Base K
func sumBase(n int, k int) (sum int) {
	for ; n > 0; n = n / k {
		sum += n % k
	}
	return
}

func main() {
	res := sumBase(34, 6)
	fmt.Println(res)
}

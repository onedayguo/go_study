package main

import "fmt"

// 1837. Sum of Digits in Base K
func sumBase(n int, k int) (sum int) {
	for ; n > 0; n = n / k {
		sum += n % k
	}
	return
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//897. Increasing Order Search Tree递归
func increasingBST(root *TreeNode) *TreeNode {
	return recurveLeft(root, nil)
}
func recurveLeft(root *TreeNode, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	var res = recurveLeft(root.Left, root)
	root.Left = nil
	root.Right = recurveLeft(root.Right, tail)
	return res
}

/*
1. Two Sum
*/
func twoSum(nums []int, target int) []int {
	// key:数，value：下标
	var existed = make(map[int]int)
	for i, num := range nums {
		if value, ok := existed[target-num]; ok {
			return []int{value, i}
		} else {
			existed[num] = i
		}
	}
	return nil
}

func main() {
	res := sumBase(34, 6)
	fmt.Println(res)
}

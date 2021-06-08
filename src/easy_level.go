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

func main() {
	res := sumBase(34, 6)
	fmt.Println(res)
}

package leetcode

/**
 * @description: 107. Binary Tree Level Order Traversal II
	Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).
 * @return: 自底向上，从左到右的顺序输出节点
 * @author: kami
 * @关键词：层序遍历,队列，追加在结果的头部
 * @date: 2021/7/4 9:50
*/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int, 1)
	var queue = make([]*TreeNode, 1)
	queue[0] = root
	var first = []int{root.Val}
	res[0] = first
	for {
		var size = len(queue)
		var curNum []int
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				curNum = append(curNum, queue[i].Left.Val)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				curNum = append(curNum, queue[i].Right.Val)
			}
		}
		if len(curNum) < 1 {
			return res
		}
		queue = append(queue[:0], queue[size:]...)
		var temp = [][]int{
			curNum,
		}
		res = append(temp, res...)
	}
}

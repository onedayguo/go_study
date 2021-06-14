package leetcode

import (
	"fmt"
)

// 1. 两数之和
func twoSum1(nums []int, target int) []int {
	hashtable := map[int]int{}
	for i, value := range nums {
		if p, ok := hashtable[target-value]; ok {
			return []int{p, i}
		}
		hashtable[value] = i
	}
	return nil
}

// Definition for singly-linked list.
type ListNode1 struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// 237. 删除链表中的节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 1470. 重新排列数组
func shuffle(nums []int, n int) []int {
	var index int
	r := make([]int, len(nums))
	for i := 0; i < n; i++ {
		r[i*2] = nums[index]
		r[i*2+1] = nums[n+index]
		index++
	}
	return r
}

// LCP 01. 猜数字
func game(guess []int, answer []int) int {
	return 0
}
func deleteNode1(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 1480. 一维数组的动态和
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}

// 1299. 将每个元素替换为右侧最大元素 如果是最后一个元素，用 -1 替换
func replaceElements(arr []int) []int {
	var end = len(arr)
	var max = arr[end-1]
	for i := end - 1; i > -1; i-- {
		var temp = arr[i]
		arr[i] = max
		if temp > max {
			max = temp
		}
	}
	arr[end-1] = -1
	return arr
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	sum := shuffle(nums, 3)
	fmt.Print(sum)
}

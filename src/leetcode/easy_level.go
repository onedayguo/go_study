package leetcode

import (
	"fmt"
)

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

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
2. Add Two Numbers
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curNode *ListNode = &ListNode{0, nil}
	var res = curNode
	var take, l1Value, l2Value int
	for l1 != nil || l2 != nil {
		l1Value = 0
		l2Value = 0
		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		}
		var sum = l1Value + l2Value + take
		var cur = sum % 10
		take = sum / 10
		var next = &ListNode{cur, nil}
		curNode.Next = next
		curNode = next

	}
	if take > 0 {
		var next = &ListNode{1, nil}
		curNode.Next = next
	}
	return res.Next
}

/*
3. Longest Substring Without Repeating Characters
*/
func lengthOfLongestSubstring(s string) int {
	var left, right, max, len int = 0, 0, 0, len(s)
	exist := make(map[byte]bool)
	for left < len && right < len {
		if _, ok := exist[s[right]]; !ok {
			exist[s[right]] = true
			right++
			max = maxNum(max, right-left)
		} else {
			delete(exist, s[left])
			left++
		}
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
4. Median of Two Sorted Arrays
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := []int{}

	for len(nums1) > 0 || len(nums2) > 0 {
		var n int

		if len(nums1) == 0 {
			n = nums2[0]
			nums2 = nums2[1:]
		} else if len(nums2) == 0 {
			n = nums1[0]
			nums1 = nums1[1:]
		} else {
			if nums1[0] > nums2[0] {
				n = nums2[0]
				nums2 = nums2[1:]
			} else {
				n = nums1[0]
				nums1 = nums1[1:]
			}
		}

		arr = append(arr, n)
	}

	l := len(arr)

	if l%2 == 0 {
		return float64(arr[l/2]+arr[l/2-1]) / 2
	} else {
		return float64(arr[l/2])
	}
}

func main() {
	var str string = "hello,world"
	fmt.Println(str[1])
}

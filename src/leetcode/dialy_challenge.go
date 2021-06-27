package leetcode

import "sort"

// Matchsticks to Square，数组能否组成一个正方形
func makesquare(matchsticks []int) bool {
	var sum = 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sort.Ints(matchsticks)

	return dfsSquare(matchsticks, make([]int, 4), len(matchsticks)-1, sum/4)
}

func dfsSquare(nums, sums []int, index, target int) bool {
	if index == -1 {
		if sums[0] == target && sums[1] == target && sums[2] == target {
			return true
		}
		return false
	}
	for i := 0; i < 4; i++ {
		if sums[i]+nums[index] > target {
			continue
		}
		sums[i] += nums[index]
		if dfsSquare(nums, sums, index-1, target) {
			return true
		}
		sums[i] -= nums[index]
	}

	return false
}

/**
 * @description: Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
 * @return: 括号的组合数量
 * @author: kami
 * @关键词：递归
 * @date: 2021/6/27 21:49
 */
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	var sign []byte
	backtrack(&result, sign, 0, 0, n)
	return result
}

func backtrack(result *[]string, sign []byte, open, close, max int) {
	if len(sign) == max*2 {
		*result = append(*result, string(sign))
		return
	}
	if open < max {
		backtrack(result, append(sign, '('), open+1, close, max)
	}
	if close < open {
		backtrack(result, append(sign, ')'), open, close+1, max)
	}
}

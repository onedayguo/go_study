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

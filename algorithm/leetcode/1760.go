package leetcode

import "sort"

func minNum(nums []int, maxOperations int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	begin := 1
	end := max
	ans := 0
	for begin <= end {
		mid := begin + (end-begin)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			count += nums[i] / mid
			if nums[i]%mid == 0 {
				count--
			}
		}
		if count <= maxOperations {
			ans = mid
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return ans
}

func minimumSize(nums []int, maxOperations int) int {
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}
	return sort.Search(max, func(y int) bool {
		if y == 0 {
			return false
		}
		ops := 0
		for _, x := range nums {
			ops += (x - 1) / y
		}
		return ops <= maxOperations
	})
}

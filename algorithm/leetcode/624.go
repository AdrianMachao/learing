package leetcode

import "math"

func maxDistancBetweenArray(arrList [][]int) int {
	minRes := arrList[0][0]
	maxRes := arrList[0][len(arrList[0])-1]
	ans := math.MaxInt
	for i := 1; i < len(arrList); i++ {
		ans = max(ans, max(abs(arrList[i][0]-maxRes), abs(arrList[i][len(arrList[i])]-minRes)))
		minRes = min(minRes, arrList[i][0])
		maxRes = max(maxRes, arrList[i][len(arrList[i])])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(num int) int {
	if num > 0 {
		return num
	}

	return -1 * num
}

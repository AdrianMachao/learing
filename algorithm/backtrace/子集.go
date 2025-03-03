package algorithm

import "sort"

func subsets(nums []int) [][]int {
	var set []int
	var ans [][]int
	var dfs func(int)
	dfs = func(level int) {
		if level == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		set = append(set, nums[level])
		dfs(level + 1)
		set = set[:len(set)-1]
		dfs(level + 1)
	}

	dfs(0)

	return ans
}

func subsetsWithDup1(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}

func subsetsWithDup2(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	dfsWithDup(nums, 0, []int{}, &ans)
	return ans
}

func dfsWithDup(nums []int, pos int, set []int, ans *[][]int) {
	if pos >= len(nums) {
		return
	}

	tmp := make([]int, len(set))
	copy(tmp, set)
	*ans = append(*ans, tmp)

	for i := pos; i < len(nums); i++ {
		// set = append(set, i)
		if i == pos || nums[i] != nums[i-1] {
			dfsWithDup(nums, i+1, append(set, nums[i]), ans)
		}
		// set = set[:len(set)-1]
	}
}

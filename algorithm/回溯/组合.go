package main

func combine(n int, k int) [][]int {
	var ans [][]int
	dfs(n, k, 1, []int{}, &ans)

	return ans
}

func dfs(n, k, pos int, set []int, ans *[][]int) {
	if len(set) == k {
		tmp := make([]int, len(set))
		copy(tmp, set)
		*ans = append(*ans, tmp)
		return
	}

	for i := pos; i <= n; i++ {
		set = append(set, i)
		dfs(n, k, i+1, set, ans)
		set = set[:len(set)-1]
	}
}

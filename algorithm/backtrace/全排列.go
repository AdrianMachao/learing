package algorithm

func permute(nums []int) [][]int {
	var ans [][]int
	var set []int
	dfsPermute(nums, set, map[int]bool{}, &ans)

	return ans
}

func dfsPermute(nums []int, set []int, visited map[int]bool, ans *[][]int) {
	if len(set) == len(nums) {
		tmp := make([]int, len(set))
		copy(tmp, set)
		*ans = append(*ans, tmp)
	}
	for _, v := range nums {
		if visited[v] {
			continue
		}
		visited[v] = true
		dfsPermute(nums, append(set, v), visited, ans)
		visited[v] = false
	}
}

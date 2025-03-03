package algorithm

// 逆时针输出

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	v := 1
	for i := 0; i < n; i++ {
		for j := i; j < n-i; j++ {
			ans[i][j] = v
			v++
		}
		for j := i + 1; j < n-i; j++ {
			ans[j][n-i-1] = v
			v++
		}
		for j := n - i - 2; j >= i; j-- {
			ans[n-i-1][j] = v
			v++
		}
		for j := n - i - 2; j > i; j-- {
			ans[j][i] = v
			v++
		}
	}
	return ans
}

package algorithm

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := range obstacleGrid {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	dp[0][0] = 1
	for i := range obstacleGrid {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i >= 1 {
				dp[i][j] += dp[i-1][j]
			}
			if j >= 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

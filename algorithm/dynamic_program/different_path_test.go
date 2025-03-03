package algorithm

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	obstacleGrid1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	ans1 := uniquePathsWithObstacles(obstacleGrid1)
	fmt.Println("--ans1:", ans1)

	obstacleGrid2 := [][]int{{0, 1}, {0, 0}}
	ans2 := uniquePathsWithObstacles(obstacleGrid2)
	fmt.Println("--ans2:", ans2)

}

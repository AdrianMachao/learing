package algorithm

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	ans := generateMatrix(4)
	for i := range ans {
		for _, v := range ans[i] {
			fmt.Printf("%v ", v)
		}
		fmt.Print("\n")
	}
}

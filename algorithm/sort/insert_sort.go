package algorithm

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		insertNum := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > insertNum {
				nums[j+1] = nums[j]
				nums[j] = insertNum
			} else {
				break
			}
		}
	}
}

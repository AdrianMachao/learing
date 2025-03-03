package algorithm

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, i, len(nums)-1)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		swap(nums, 0, i)
		adjust(nums, 0, i-1)
	}
}

func adjust(nums []int, begin, end int) {
	left, right := begin*2+1, begin*2+2
	maxIndex := begin
	if left <= end && nums[left] > nums[maxIndex] {
		maxIndex = left
	}

	if right <= end && nums[right] > nums[maxIndex] {
		maxIndex = right
	}

	if maxIndex != begin {
		swap(nums, begin, maxIndex)
		adjust(nums, maxIndex, end)
	}
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

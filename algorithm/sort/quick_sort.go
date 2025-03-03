package algorithm

func partition(nums []int, begin, end int) int {
	pivot := nums[begin]
	for begin < end {
		for begin < end && nums[end] >= pivot {
			end--
		}
		nums[begin] = nums[end]
		for begin < end && nums[begin] <= pivot {
			begin++
		}
		nums[end] = nums[begin]
	}
	nums[begin] = pivot

	return begin
}

func quicksort(nums []int, begin, end int) {
	if begin >= end {
		return
	}

	pivot := partition(nums, begin, end)
	quicksort(nums, begin, pivot-1)
	quicksort(nums, pivot+1, end)
}

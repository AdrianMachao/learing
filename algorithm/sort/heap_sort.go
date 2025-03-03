package algorithm

func mergeSort(nums []int, i, j int) {
	if i >= j {
		return
	}

	mid := (i + j) / 2
	mergeSort(nums, i, mid)
	mergeSort(nums, mid+1, j)
	merge(nums, i, mid, j)
}

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, 0, right-left+1)
	i, j := left, mid+1
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}

	for i := left; i <= right; i++ {
		nums[i] = tmp[i-left]
	}
}

package main

func moveZeroes(nums []int) []int {
	k := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
	return nums
}

func moveZeroesNoSwap(nums []int) []int {
	k := 0
	for _, num := range nums {
		if num != 0 {
			nums[k] = num
			k++
		}
	}

	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

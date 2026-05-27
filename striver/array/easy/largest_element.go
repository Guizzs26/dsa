package main

func findLargestElement(nums []int) int {
	largest := nums[0]

	for i := range nums {
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	return largest
}

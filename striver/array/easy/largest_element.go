package main

func findLargestElement(nums []int) int {
	largest := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	return largest
}

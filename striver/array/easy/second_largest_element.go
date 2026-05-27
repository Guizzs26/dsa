package main

import "math"

func findSecondLargestAndSmallestElement(nums []int) (int, int) {
	if len(nums) < 2 {
		return -1, -1
	}

	largest, secondLargest := math.MinInt32, math.MinInt32
	smallest, secondSmallest := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest && num != largest {
			secondLargest = num
		}

		if num < smallest {
			secondSmallest = smallest
			smallest = num
		} else if num < secondSmallest && num != smallest {
			secondSmallest = num
		}
	}

	if secondLargest == math.MinInt32 || secondSmallest == math.MaxInt32 {
		return -1, -1
	}

	return -1, -1
}

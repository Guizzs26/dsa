package leetcode

import (
	"fmt"
	"sort"
)

// Frequency problem (not the correct solution for the n/2 majority problem)
// Time - O(n) + O(n)
// Space - O(n)
func majorityElement(nums []int) int {
	seen := make(map[int]int, len(nums))
	for _, num := range nums {
		seen[num]++
		fmt.Println(seen)
	}

	var majority, maxCount int
	for num, count := range seen {
		if count > maxCount {
			maxCount = count
			majority = num
		}
	}
	return majority
}

// Time -> O(n)
// Space -> O(n)
func majorityElementMap(nums []int) int {
	m := make(map[int]int, 0)
	for _, n := range nums {
		m[n]++

		if m[n] > len(nums)/2 {
			return n
		}
	}
	return -1
}

// Time -> O(n)
// Space -> O(n)
func majorityElementMapOtimized(nums []int) int {
	condition := len(nums) / 2
	frequency := make(map[int]int, len(nums))
	for _, num := range nums {
		frequency[num]++
		if frequency[num] > condition {
			return num
		}
	}
	return -1
}

// Time -> O(n logn)
// Space -> O(logn)
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElementSortSafe(nums []int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)

	return numsCopy[len(numsCopy)/2]
}

func majorityElementSortValidate(nums []int) int {
	sort.Ints(nums)
	candidate := nums[len(nums)/2]
	count := 0

	for _, n := range nums {
		if n == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}
	return -1
}

func majorityElementGoldenSolution(nums []int) int {
	var candidate, count int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

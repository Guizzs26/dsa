package leetcode

import (
	"sort"
)

// Brute force solution
// Time -> O(n2)
// Space -> O(1)
func HasDuplicateBruteForce(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Map solution
// Time -> O(n)
// Space -> O(n)
func HasDuplicateMap(nums []int) bool {
	duplicates := make(map[int]bool)
	for i := range nums {
		if duplicates[nums[i]] {
			return true
		}
		duplicates[nums[i]] = true
	}

	return false
}

func HasDuplicateMap2(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// Quick sort solution (build-in Go function)
// Time -> O(nlogn)
// Space -> O(1)
func HasDuplicateSort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// Map length solution
// Time -> O(n)
// Space -> O(n)
func HasDuplicateLength(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		seen[num] = struct{}{}
	}
	return len(seen) < len(nums)
}

func HasDuplicateGoldenSolution(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

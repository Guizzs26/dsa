package main

import (
	"slices"
)

func hasDuplicate(nums []int) bool {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v == nums[j] {
				return true
			}
		}
	}
	return false
}

func hasDuplicateHashSet(nums []int) bool {
	set := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}

		set[nums[i]] = struct{}{}
	}

	return false
}

func hasDuplicateHashSetIdiomatic(nums []int) bool {
	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}

	return false
}

func hashDuplicateSort(nums []int) bool {
	slices.Sort(nums)

	for i := range nums {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

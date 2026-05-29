package main

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHashMap(nums []int, target int) []int {
	var compl int

	hm := make(map[int]int, len(nums))
	for i, num := range nums {
		compl = target - num

		if _, ok := hm[compl]; ok {
			return []int{hm[compl], i}
		}

		hm[num] = i
	}

	return nil
}

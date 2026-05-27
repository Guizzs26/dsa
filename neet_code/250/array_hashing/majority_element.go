package main

func majorityElement(nums []int) int {
	hm := make(map[int]int, (len(nums)/2)+1)

	for _, num := range nums {
		hm[num]++

		if hm[num] > (len(nums) / 2) {
			return num
		}
	}

	return 0
}

func majorityElementBoyerMoore(nums []int) int {
	candidate, count := 0, 0

	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}

		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

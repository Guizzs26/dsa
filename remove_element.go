package main

func removeElementsBruteForce(nums []int) int {
	i := 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}

func removeElementsHashSetWrong(nums []int) int {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	k := 0
	for num := range set {
		nums[k] = num
		k++
	}

	return k
}

func removeElementsHashSet(nums []int) int {
	hs := make(map[int]struct{}, (len(nums)/2)+1)
	k := 0

	for _, num := range nums {
		if _, ok := hs[num]; !ok {
			hs[num] = struct{}{}
			nums[k] = num
			k++
		}
	}

	return k
}

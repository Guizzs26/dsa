package main

func singleNumber(nums []int) int {
	hs := make(map[int]struct{}, (len(nums)/2)+1)

	for _, num := range nums {
		if _, ok := hs[num]; ok {
			delete(hs, num)
		} else {
			hs[num] = struct{}{}
		}
	}

	for k := range hs {
		return k
	}

	return -1
}

func singleNumberXOR(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}

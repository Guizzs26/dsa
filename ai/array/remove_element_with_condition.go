package main

import "fmt"

func removeElement(nums []int, limit int) []int {
	k := 0

	for i, num := range nums {
		if num < limit {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	return nums[:k]
}

func main() {
	fmt.Println(removeElement([]int{1, 10, 2, 3, 15, 4}, 5))
}

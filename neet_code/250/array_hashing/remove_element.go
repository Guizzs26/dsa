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

/*
Essa solução peca apenas na segunda parte, o motivo é o seguinte:

Range em maps são não-determinísticos e o enunciado do problema exige que a posição relativa dos elementos permaneça
a mesma no final.

Um range no mesmo map poderia retornar diferentes slices:

[1, 2, 3] ou [1, 3, 2] ou [3, 1, 2] ou [3, 2, 1] etc...
*/
func removeElementsHashSetWrong(nums []int) int {
	hs := make(map[int]struct{}, (len(nums)/2)+1)

	for _, num := range nums {
		hs[num] = struct{}{}
	}

	k := 0
	for num := range hs {
		nums[k] = num
		k++
	}

	return k
}

func removeElementsHashSet(nums []int) int {
	hs := make(map[int]struct{}, (len(nums)/2)+1)

	k := 0
	for _, num := range nums {
		if _, exists := hs[num]; !exists {
			hs[num] = struct{}{}
			nums[k] = num
			k++
		}
	}

	return k
}

func removeElements(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

package main

/*

Enunciado: Dado um array de inteiros nums onde todos os elementos aparecem exatamente duas vezes,
exceto por um elemento que aparece apenas uma vez. Encontre e retorne esse elemento solitário.



Exemplo 1: * Input: nums = [2, 2, 1]

Output: 1

Exemplo 2: * Input: nums = [4, 1, 2, 1, 2]

Output: 4

*/

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

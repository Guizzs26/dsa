package main

/*

Enunciado: Dado um array de inteiros nums e um número inteiro target, encontre dois números no array cujo produto (multiplicação)
seja exatamente igual ao target (A x B = target). Retorne os índices desse par.

Exemplo:

Input: nums = [2, 5, 8, 3, 9], target = 24
Output: [2, 3] (8 x 3 = 24)

*/

func productSum(nums []int, target int) []int {
	hm := make(map[int]int, len(nums))

	for i, num := range nums {
		if num == 0 {
			if target == 0 {
				for _, pastIdx := range hm {
					return []int{pastIdx, i}
				}
			}
			hm[num] = i
			continue
		}

		if target%num == 0 {
			compl := target / num
			if pastIdx, ok := hm[compl]; ok {
				return []int{pastIdx, i}
			}
		}
		hm[num] = i
	}

	return nil
}

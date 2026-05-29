package main

/*

Enunciado: Dado un array de inteiros desordenados nums e um número inteiro target, encontre um par de elementos de
forma que a diferença entre eles seja igual ao target (A - B = target). Retorne os índices dos dois elementos.

Você pode assumir que sempre haverá uma solução única.

Exemplo:

Input: nums = [5, 20, 3, 2, 12], target = 9
Output: [4, 2] (Porque 12 - 3 = 9. O índice do 12 é 4 e o do 3 é 2)

*/

func twoSumII(nums []int, target int) []int {
	hs := make(map[int]int, len(nums))

	for i, num := range nums {
		compl1 := num - target
		compl2 := target + num

		if _, ok := hs[compl1]; ok {
			return []int{hs[compl1], i}
		}

		if _, ok := hs[compl2]; ok {
			return []int{hs[compl2], i}
		}

		hs[num] = i
	}

	return nil
}

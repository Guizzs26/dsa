package main

/*

Enunciado: Dados dois arrays de inteiros nums1 e nums2, retorne um novo array contendo os elementos que aparecem em ambos os arrays.
O array de resultado não pode conter elementos duplicados (cada número deve ser único no resultado).

Dica: Use um mapa/set para guardar os números de nums1 e depois use outro mapa/set para filtrar o que bate com nums2.

Exemplo 1: * Input: nums1 = [1, 2, 2, 1], nums2 = [2, 2]

Output: [2]

Exemplo 2: * Input: nums1 = [4, 9, 5], nums2 = [9, 4, 9, 8, 4]

Output: [4, 9] (ou [9, 4], a ordem não importa)

*/

func intersectionOfArrays(nums1, nums2 []int) []int {
	set1 := make(map[int]struct{}, len(nums1))
	set2 := make(map[int]struct{}, len(nums2))

	for _, num := range nums1 {
		if _, ok := set1[num]; !ok {
			set1[num] = struct{}{}
		}
	}

	for _, num := range nums2 {
		if _, ok := set1[num]; ok {
			set2[num] = struct{}{}
		}
	}

	result := make([]int, 0, len(set2))

	for k := range set2 {
		result = append(result, k)
	}

	return result
}

func intersectionOfArraysOptimized(nums1, nums2 []int) []int {
	set := make(map[int]struct{}, len(nums1))

	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	result := make([]int, 0, len(nums2))
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			result = append(result, num)
			delete(set, num)
		}
	}

	return result
}

package main

/*

Enunciado: Dada uma string s, encontre o primeiro caractere não repetido nela e retorne o seu índice.
Se todos os caracteres se repetirem, retorne -1.

Exemplo 1: Input: s = "leetcode"

Output: 0 (A letra 'l' é a primeira que não se repete)

Exemplo 2: Input: s = "loveleetcode"

Output: 2 (As letras 'l' e 'o' se repetem mais à frente. A letra 'v' é a primeira única)

*/

func firstUnique(s string) int {
	freq := make(map[rune]int, len(s))

	for _, char := range s {
		freq[char]++
	}

	for i, char := range s {
		if freq[char] == 1 {
			return i
		}
	}

	return -1
}

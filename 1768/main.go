package main

import (
	"fmt"
)

/*
1768. Merge Strings Alternately
*/
func main() {
	fmt.Println(mergeAlternately("abc", "def"))
}

func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)

	var n int
	if len1 >= len2 {
		n = len1
	} else {
		n = len2
	}

	res := make([]byte, len1+len2)

	var lastInd int
	for i := 0; i < n; i++ {
		if len1 > i {
			res[lastInd] = word1[i]
			lastInd++
		}
		if len2 > i {
			res[lastInd] = word2[i]
			lastInd++
		}
	}

	return string(res)
}

func mergeAlternately2(word1 string, word2 string) string {
	res := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		res += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		res += string(word1[i])
		i++
	}

	for j < len(word2) {
		res += string(word2[j])
		j++
	}

	return res
}

func mergeAlternately3(word1 string, word2 string) string {
	res := make([]rune, len(word1)+len(word2))
	pos1, pos2 := 0, 1
	for i, ch := range word1 {
		res[pos1] = ch
		if i < len(word2) {
			pos1 += 2
		} else {
			pos1++
		}
	}
	for i, ch := range word2 {
		res[pos2] = ch
		if i < len(word1)-1 {
			pos2 += 2
		} else {
			pos2++
		}
	}
	return string(res)
}

package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABCDEF", "ABC"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len1 := len(str1)
	len2 := len(str2)

	var i, j int
	var maxToken []byte

	for i < len1 && j < len2 {
		if str1[i] == str2[i] {
			maxToken = append(maxToken, str1[i])
		}
		i++
		j++
	}

	for len(maxToken) > 1 {
		if len1%len(maxToken) == 0 && len2%len(maxToken) == 0 {
			break
		}
		maxToken = maxToken[:len(maxToken)-1]
	}

	return string(maxToken)
}

package main

import (
	"fmt"
)

/*
Group Anagrams
*/
func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string

	m := make(map[[26]uint][]string, len(strs))

	for _, str := range strs {
		hash := [26]uint{}
		for i := 0; i < len(str); i++ {
			hash[str[i]-'a']++
		}
		m[hash] = append(m[hash], str)
	}

	for _, strings := range m {
		res = append(res, strings)
	}

	return res
}

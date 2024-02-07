package main

import (
	"fmt"
)

func main() {
	res := groupAnagrams([]string{
		"sas", "mat", "tam", "mta", "duh", "ill",
	})
	fmt.Println(res)
	//fmt.Println(getAsciiTotal("duh"))
	//fmt.Println(getAsciiTotal("ill"))
}

func groupAnagrams(ls []string) [][]string {
	group := map[[26]byte][]string{}
	result := make([][]string, 0)
	for _, s := range ls {
		key := getKey(s)
		group[key] = append(group[key], s)
	}
	for _, val := range group {
		result = append(result, val)
	}
	return result
}

func getKey(s string) [26]byte {
	res := [26]byte{}
	for _, c := range s {
		res[c-'a'] += 1
	}
	return res
}

package main

import (
	"fmt"
	"unicode"
)

func longestNiceSubstringV1(s string) string {
	n := len(s)
	longest := ""
	for i := range n {
		for j := i + 1; j <= n; j++ {
			sub := s[i:j]
			if isNice(sub) && len(sub) > len(longest) {
				longest = sub
			}
		}
	}

	return longest
}

func isNice(sub string) bool {
	hashMap := make(map[rune]bool)
	for _, key := range sub {
		hashMap[key] = true
	}

	for key := range hashMap {
		if !hashMap[unicode.ToLower(key)] || !hashMap[unicode.ToUpper(key)] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(longestNiceSubstringV1("YazaAay")) // Output: "aAa"
	fmt.Println(longestNiceSubstringV1("kkJjKaAbBBd"))
	// fmt.Println(longestNiceSubstring("Bb"))      // Output: "Bb"
	// fmt.Println(longestNiceSubstring("c"))       // Output: ""
}

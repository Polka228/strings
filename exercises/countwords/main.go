package main

import (
	"fmt"
	"unicode"
)

// TestCountWords tests the function `countWords`.
func TestCountWords() {
}

// countWords counts the number of words that occurs in a string `s`.
func countWords(s string) (cnt int) {
	var wasLetter bool
	for _, v := range s {
		if (unicode.IsPunct(v) || unicode.IsSpace(v)) && wasLetter {
			wasLetter = false
		} else if !wasLetter {
			cnt++
			wasLetter = true
		}
	}
	return cnt
}

func main() {
	//TestCountWords()
	fmt.Println(countWords("привет, всем"))
}

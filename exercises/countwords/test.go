package main

import (
	"fmt"
)

// TestCountWords tests the function `countWords`.
func TestCountWords() {
	for _, tc := range []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"separators", "..,,.\n \t!-+=*", 0},
		{"single word", "abc", 1},
		{"single word with spaces", "   ccc  ", 1},
		{"two words separated", "aa b", 2},
		{"words plus digits", "1a 1 b1", 3},
		{"mixed unicode", "čeлovek", 1},
		{"words plus punctuation", "cat, dog, fish", 3},
	} {
		if got := countWords(tc.input); got != tc.want {
			fmt.Printf("FAIL %s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func TestWords() {
	pass := true
	for _, tc := range []struct {
		name  string
		input string
		want  []string
	}{

		{"empty", "", nil},
		{"separators", " *-+!,.\n\r\t", nil},
		{"single word", "abc", []string{"abc"}},
		{"single word in spaces", "   abc   ", []string{"abc"}},
		{"two words", "abc def", []string{"abc", "def"}},
		{"two words in spaces", "   abc   def   ", []string{"abc", "def"}},
		{"unicode", "привет, мир!", []string{"привет", "мир"}},
	} {
		if diff := cmp.Diff(tc.want, words(tc.input)); diff != "" {
			fmt.Printf("%s\n\tFAIL: got unexpected diff (-want, +got): %v", tc.name, diff)
			pass = false
		}
	}
	if pass {
		fmt.Println("ALL PASS")
	}
}

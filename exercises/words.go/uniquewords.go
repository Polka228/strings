package main

import (
	"fmt"
	"unicode"

	"github.com/google/go-cmp/cmp"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func countWords(s string) (cnt int) {
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			wasSep = true
		} else if wasSep {
			wasSep = false
			cnt++
		}
	}
	return cnt
}

func TestCountWords() {
	pass := true
	for _, tc := range []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"separators", " *-+!,.\n\r\t", 0},
		{"single word", "abc", 1},
		{"single word in spaces", "   abc   ", 1},
		{"two words", "abc def", 2},
		{"two words in spaces", "   abc   def   ", 2},
		{"unicode", "привет, мир!", 2},
	} {
		if got := countWords(tc.input); got != tc.want {
			fmt.Printf("%s\n\tFAIL: got %v, want %v\n", tc.name, got, tc.want)
			pass = false
		}
	}
	if pass {
		fmt.Println("ALL PASS")
	}
}

func words(s string) []string {
	var w []string
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				w = append(w, string(tmp))
				// tmp = tmp[:0]
				tmp = nil
			}
			wasSep = true
		} else {
			wasSep = false
			tmp = append(tmp, r)
		}
	}
	if len(tmp) > 0 {
		w = append(w, string(tmp))
	}
	return w
}

// uniqueWords returns a slice of unique
func uniqueWords(s string) []string {
	return nil
}

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

func TestUniqueWords() {

}

func main() {
	//TestCountWords()
	//TestWords()
	TestUniqueWords()

}

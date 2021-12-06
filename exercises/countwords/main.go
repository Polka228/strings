func isSep(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsSymbol(r)
}

// countWords counts the number of words that occurs in a string `s`.
func countWords(s string) (cnt int) {
	wasLetter := false
	for _, v := range s {
		if isSep(v) && wasLetter {
			wasLetter = false
		} else if !wasLetter {
			cnt++
			wasLetter = true
		}
		if isSep(v) == true {
			wasLetter = false
		} else if !wasLetter {
			cnt++
			wasLetter = true
		}
	}
	return cnt
}

func main() {
	fmt.Println(isSep('\n'))
	TestCountWords()
	fmt.Println(countWords("привет, всем"))
}
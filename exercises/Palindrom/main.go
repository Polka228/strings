package main

func IsPalidndrom(s string) bool {
	r := []rune(s)
	i,j :=0,len(r)-1
	for ; i < j, j = i + 1,j - 1{
		if r[i] != r[j]{
			return false
		}
	}
	return true
}

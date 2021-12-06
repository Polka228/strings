package main

import "testing"

func countWords(t *testing.T){
	for -,tc := range []struct{
		input string
		want uint
	} {
		{"",0}
		{"abc",1}
		{"this is test",3}
		{"ilike apples",2}
		{"wtahisgoingon",1}
		{"a b c d e f g",7}
		{"1a 2b 4 6",4}
		{"2134",1}
		
	}
} {
	t.Run(tc.input,func(t *testing.T)){
		if got := countWords(tc.input);got != tc.want{
			t.Errorf("got = %v,want = %q",got,tc.want)
		}
	}
}
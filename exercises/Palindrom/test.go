package main

func TestPalindrom(t *testing.T){
	for -,tc := range []struct{
		input string
		want bool
	} {
		{"",true}
		{"abc",false}
		{"aaa",true}
		{"242",true}
		{"a",true}
		{"     huh  ",false}
		{"ātram slidas sadils martā",true}
		{"čč",true}
		{"ūuūu",false}
		{"lļl",true}
	}
} {
	t.Run(tc.input,func(t *testing.T)){
		if got := IsPalindrom(tc.input);got != tc.want{
			t.Errorf("got = %v,want = %q",got,tc.want)
		}
	}
}
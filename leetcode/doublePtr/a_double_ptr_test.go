package doublePtr

import "testing"

// 541. 反转字符串 II
func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "reverseStr", args: args{s: "abcdefg", k: 2}, want: "bacdfeg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStrWithByte(tt.args.s, tt.args.k); got != tt.want {
				//if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 151. 反转字符串中的单词
func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "reverse", args: args{s: "a good   example"}, want: "example good a"},
		{name: "reverse", args: args{s: "  hello world  "}, want: "world hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWordsWithO1(tt.args.s); got != tt.want {
				//if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1768. 交替合并字符串
func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "merge", args: args{word1: "abc", word2: "pqr"}, want: "apbqcr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}

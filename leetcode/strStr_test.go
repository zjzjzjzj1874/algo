package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "hello", args: args{haystack: "hello", needle: "hello"}, want: 0},
		//{name: "hello", args: args{haystack: "butsad", needle: "sad"}, want: 3},
		{name: "hello", args: args{haystack: "BBC ABCDABABCDABCDABDE", needle: "ABCDABD"}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrKMP(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

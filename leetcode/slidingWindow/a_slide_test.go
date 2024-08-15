package slidingWindow

import "testing"

// 1456. 定长子串中元音的最大数目
func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxVowels", args: args{s: "abciiidef", k: 3}, want: 3},
		{name: "maxVowels", args: args{s: "novowels", k: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

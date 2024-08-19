package slidingWindow

import "testing"

// 1493. 删掉一个元素以后全为 1 的最长子数组
func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "longest", args: args{nums: []int{1, 0, 0, 0, 0}}, wantAns: 1},
		{name: "longest", args: args{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, wantAns: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubarray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("longestSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 424. 替换后的最长重复字符
func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "character", args: args{s: "AABABBA", k: 2}, want: 5},
		{name: "character", args: args{s: "AABABA", k: 1}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 2024. 考试的最大困扰度
func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxConsecutiveAnswers", args: args{answerKey: "TTFTTFTT", k: 1}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1759. 统计同质子字符串的数目
func Test_countHomogenous(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "count", args: args{s: "zzzzz"}, wantAns: 15},
		{name: "count", args: args{s: "abbcccaa"}, wantAns: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countHomogenous(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("countHomogenous() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 2414. 最长的字母序连续子字符串的长度
func Test_longestContinuousSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "longestContinuousSubstring", args: args{s: "abcde"}, wantAns: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestContinuousSubstring(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("longestContinuousSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 1423. 可获得的最大点数
func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "maxScore", args: args{cardPoints: []int{1, 2, 3, 4, 5, 6, 1}, k: 3}, wantAns: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxScoreWithOptimize(tt.args.cardPoints, tt.args.k); gotAns != tt.wantAns {
				//if gotAns := maxScore(tt.args.cardPoints, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 2379. 得到 K 个黑块的最少涂色次数
func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "minimumRecolors", args: args{blocks: "WBBWWBBWBW", k: 7}, wantAns: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumRecolors(tt.args.blocks, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("minimumRecolors() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

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

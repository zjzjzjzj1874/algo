package backtrack

import (
	"reflect"
	"testing"
)

// LCR 102. 目标和
func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "findTargetSumWays", args: args{nums: []int{1, 1, 1, 1, 1}, target: 3}, wantAns: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findTargetSumWays(tt.args.nums, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("findTargetSumWays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 085. 括号生成
func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{name: "generate", args: args{n: 3}, wantAns: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := generateParenthesis(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("generateParenthesis() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 087. 复原 IP 地址
func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{name: "restore", args: args{s: "25525511135"}, wantAns: []string{"255.255.11.135", "255.255.111.35"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("restoreIpAddresses() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 086. 分割回文串
func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{name: "partition", args: args{s: "aab"}, wantAns: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{name: "partition", args: args{s: "google"}, wantAns: [][]string{{"g", "o", "o", "g", "l", "e"}, {"g", "oo", "g", "l", "e"}, {"goog", "l", "e"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if gotAns := partition(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
			if gotAns := partitionWithAns(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("partition() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 079. 子集
func Test_subSet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "subsets", args: args{nums: []int{1, 2, 3}}, wantAns: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subsetsWithDFSAns(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				//if gotAns := subsetsWithDFSInput(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				//if gotAns := subsets(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("subsets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 080. 组合
func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "combine", args: args{n: 4, k: 2}, wantAns: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combination() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 081. 组合总和
func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "combination", args: args{
			candidates: []int{2, 3, 6, 7},
			target:     7,
		}, wantAns: [][]int{{7}, {2, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 082. 组合总和 II
func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "combinationSum2", args: args{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
		}, wantAns: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum2() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 083. 全排列
func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "permute", args: args{nums: []int{1, 2, 3}}, wantAns: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := permuteWithDfs(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				//if gotAns := permute(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("permute() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 084. 全排列II
func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "permuteUnique", args: args{nums: []int{1, 1, 2}}, wantAns: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := permuteUnique(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("permuteUnique() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

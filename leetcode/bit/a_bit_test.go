package bit

import (
	"reflect"
	"testing"
)

// 3158. 求出出现两次数字的 XOR 值
func Test_duplicateNumbersXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "xor", args: args{nums: []int{50, 40, 50}}, wantAns: 50},
		{name: "xor", args: args{nums: []int{1, 2, 3}}, wantAns: 0},
		{name: "xor", args: args{nums: []int{1, 2, 1, 3}}, wantAns: 1},
		{name: "xor", args: args{nums: []int{1, 2, 2, 1, 3, 3}}, wantAns: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := duplicateNumbersXORWithBitOp(tt.args.nums); gotAns != tt.wantAns {
				//if gotAns := duplicateNumbersXOR(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("duplicateNumbersXOR() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 005. 最大单词长度乘积
func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxProduct", args: args{words: []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}}, want: 16},
		{name: "maxProduct", args: args{words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductWithLT(tt.args.words); got != tt.want {
				//if got := maxProduct(tt.args.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}

}

// LCR 003. 比特位计数
func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "countBit", args: args{n: 5}, want: []int{0, 1, 1, 2, 1, 2}},
		{name: "countBit", args: args{n: 7}, want: []int{0, 1, 1, 2, 1, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsWithLT(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				//if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 136. 只出现一次的数字
func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "singleNumber", args: args{nums: []int{2, 2, 1}}, want: 1},
		{name: "singleNumber", args: args{nums: []int{4, 1, 2, 2, 1}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 137. 只出现一次的数字 II
func Test_singleNumber137(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "xor", args: args{nums: []int{2, 2, 3, 2}}, want: 3},
		{name: "xor", args: args{nums: []int{0, 1, 0, 1, 0, 1, 99}}, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber137WithBit(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 260. 只出现一次的数字III
func Test_singleNumber260(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "singleNumber", args: args{nums: []int{2, 2, 1, 3}}, want: []int{1, 3}},
		{name: "singleNumber", args: args{nums: []int{4, 1, 2, 2, 1, 5}}, want: []int{4, 5}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber260(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

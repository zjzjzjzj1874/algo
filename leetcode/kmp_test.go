package leetcode

import (
	"reflect"
	"testing"
)

func TestKMP(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "KMP", args: args{str1: "BBC ABCDABABCDABCDABDE", str2: "ABCDABD"}, want: 14},
		{name: "KMP", args: args{str1: "sadbutsad", str2: "sad"}, want: 0},
		{name: "KMP", args: args{str1: "leetcode", str2: "leeto"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		str2 string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{name: "getNextArr", args: args{str2: "asdf"}, want: []int{-1, 0, 0, 0}},
		//{name: "getNextArr", args: args{str2: "asdfasdf"}, want: []int{-1, 0, 0, 0, 0, 1, 2, 3}},
		//{name: "getNextArr", args: args{str2: "asdfasdfasdf"}, want: []int{-1, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7}},
		{name: "getNextArr", args: args{str2: "abac"}, want: []int{0, 0, 1, 0}},
		{name: "getNextArr", args: args{str2: "abab"}, want: []int{0, 0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := GetNextV1(tt.args.str2); !reflect.DeepEqual(got, tt.want) {
			if got := getNext(tt.args.str2); !reflect.DeepEqual(got, tt.want) {
				//if got := GetNextV2(tt.args.str2); !reflect.DeepEqual(got, tt.want) {
				//if got := GetNext(tt.args.str2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

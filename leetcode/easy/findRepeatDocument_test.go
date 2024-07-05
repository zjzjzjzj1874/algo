package easy

import "testing"

func Test_findRepeatDocument(t *testing.T) {
	type args struct {
		documents []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "findRepeat", args: args{documents: []int{2, 5, 3, 0, 5, 0}}, want: 0}, // want：0 、5
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatDocument(tt.args.documents); got != tt.want {
				t.Errorf("findRepeatDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

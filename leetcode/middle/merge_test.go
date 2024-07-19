package middle

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "foo", args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "foo", args: args{intervals: [][]int{{1, 4}, {4, 5}}}, want: [][]int{{1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_meetingHolder(t *testing.T) {
	type args struct {
		meets [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		////[[1,3],[2,6],[8,10],[15,18]] => 2
		////[[1,4],[2,5],[3,6],[4,7]] => 3
		////[[1,6],[2,5],[3,7],[4,7]] => 4
		{name: "minMeet", args: args{meets: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: 2},
		{name: "minMeet", args: args{meets: [][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}}}, want: 3},
		{name: "minMeet", args: args{meets: [][]int{{1, 6}, {2, 5}, {3, 7}, {4, 7}}}, want: 4},
		{name: "minMeet", args: args{meets: [][]int{{1, 6}, {2, 5}, {3, 4}, {4, 7}}}, want: 3},
		{name: "minMeet", args: args{meets: [][]int{{1, 6}, {2, 5}, {2, 5}, {3, 4}, {4, 7}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := meetingRooms(tt.args.meets); got != tt.want {
			if got := meetingHolder(tt.args.meets); got != tt.want {
				t.Errorf("meetingHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

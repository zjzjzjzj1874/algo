package monitonicStack

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "asteroid collision", args: args{asteroids: []int{5, 10, -5}}, want: []int{5, 10}},
		{name: "asteroid collision", args: args{asteroids: []int{8, -8}}, want: []int{}},
		{name: "asteroid collision", args: args{asteroids: []int{-2, 2, 1, -2}}, want: []int{-2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollisionWithLT(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}

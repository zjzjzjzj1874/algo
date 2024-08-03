package middle

import "testing"

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		//points = [[2,2],[-1,-2],[-4,4],[-3,1],[3,-3]], s = "abdca"
		{name: "maxPointsInsideSquare", args: args{points: [][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, s: "abdca"}, wantAns: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxPointsInsideSquare(tt.args.points, tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

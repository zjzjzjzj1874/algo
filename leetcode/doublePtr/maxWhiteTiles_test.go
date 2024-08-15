package doublePtr

import "testing"

// 2271. 毯子覆盖的最多白色砖块数
func Test_maximumWhiteTiles(t *testing.T) {
	type args struct {
		tiles     [][]int
		carpetLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// [1,5],[10,11],[12,18],[20,25],[30,32]
		{name: "maximumWhiteTiles", args: args{tiles: [][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}, carpetLen: 15}, want: 14},
		{name: "maximumWhiteTiles", args: args{tiles: [][]int{{1, 1}}, carpetLen: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWhiteTiles(tt.args.tiles, tt.args.carpetLen); got != tt.want {
				t.Errorf("maximumWhiteTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

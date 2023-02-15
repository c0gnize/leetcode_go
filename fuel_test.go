package leetcode

import "testing"

func Test_minimumFuelCost(t *testing.T) {
	type args struct {
		roads [][]int
		seats int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{[][]int{{0, 1}, {0, 2}, {0, 3}}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumFuelCost(tt.args.roads, tt.args.seats); got != tt.want {
				t.Errorf("minimumFuelCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

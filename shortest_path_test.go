package leetcode

import (
	"reflect"
	"testing"
)

func Test_shortestAlternatingPaths(t *testing.T) {
	type args struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {"1", args{3, [][]int{{0, 1}, {1, 2}}, [][]int{}}, []int{0, 1, -1}},
		{"1", args{3, [][]int{{0, 1}, {0, 2}}, [][]int{{1, 0}}}, []int{0, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestAlternatingPaths(tt.args.n, tt.args.redEdges, tt.args.blueEdges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestAlternatingPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

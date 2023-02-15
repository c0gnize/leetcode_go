package leetcode

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 1, 1, 1}}, 3},
		{"2", args{[]int{3, 4, 3, 2, 5, 4, 3}}, 3},
		{"3", args{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

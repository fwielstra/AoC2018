package main

import "testing"

func Test_calculateFrequency(t *testing.T) {
	type args struct {
		deltas []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{
				[]int{1, -2, 3, 1},
			},
			3,
		},
		{
			"example2",
			args{
				[]int{1, 1, 1},
			},
			3,
		},
		{
			"example3",
			args{
				[]int{1, 1, -2},
			},
			0,
		},
		{
			"example4",
			args{
				[]int{-1, -2, -3},
			},
			-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFrequency(tt.args.deltas); got != tt.want {
				t.Errorf("calculateFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

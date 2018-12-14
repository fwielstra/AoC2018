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

func Test_findFirstFrequency(t *testing.T) {
	type args struct {
		deltas []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"findFirstFrequency example 1",
			args{
				[]int{1, -2, 3, 1},
			},
			2,
		},
		{
			"findFirstFrequency example 2",
			args{
				[]int{1, -1},
			},
			0,
		},
		{
			"findFirstFrequency example 3",
			args{
				[]int{3, 3, 4, -2, -4},
			},
			10,
		},
		{
			"findFirstFrequency example 4",
			args{
				[]int{-6, 3, 8, 5, -6},
			},
			5,
		},
		{
			"findFirstFrequency example 5",
			args{
				[]int{7, 7, -2, -7, -4},
			},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstFrequency(tt.args.deltas); got != tt.want {
				t.Errorf("findFirstFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

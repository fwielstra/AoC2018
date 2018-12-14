package main

import (
	"reflect"
	"testing"
)

func Test_parseClaimString(t *testing.T) {
	type args struct {
		claimString string
	}
	tests := []struct {
		name string
		args args
		want Claim
	}{
		{
			"Example 1",
			args{"#123 @ 3,2: 5x4"},
			Claim{
				id:     "123",
				pos: Coordinate{3, 2},
				width:  5,
				height: 4,
			},
		},
		{
			"Example 2",
			args{"#1 @ 1,3: 4x4"},
			Claim{
				id:     "1",
				pos: Coordinate{1, 3},
				width:  4,
				height: 4,
			},
		},
		{
			"Example 3",
			args{"#2 @ 3,1: 4x4"},
			Claim{
				id:     "2",
				pos: Coordinate{3, 1},
				width:  4,
				height: 4,
			},
		},
		{
			"Example 4",
			args{"#3 @ 5,5: 2x2"},
			Claim{
				id:     "3",
				pos: Coordinate{5, 5},
				width:  2,
				height: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseClaimString(tt.args.claimString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseClaimString() = %v, want %v", got, tt.want)
			}
		})
	}
}

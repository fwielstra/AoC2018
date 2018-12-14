package main

import "testing"

func Test_calculateChecksum(t *testing.T) {
	type args struct {
		boxIds []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{[]string{
				"abcdef",
				"bababc",
				"abbcde",
				"abcccd",
				"aabcdd",
				"abcdee",
				"ababab",
			}},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateChecksum(tt.args.boxIds); got != tt.want {
				t.Errorf("calculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCommonCharacters(t *testing.T) {
	type args struct {
		boxIds []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Example 1",
			args{[]string{
				"abcde",
				"fghij",
				"klmno",
				"pqrst",
				"fguij",
				"axcye",
				"wvxyz",
			}},
			"fgij",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommonCharacters(tt.args.boxIds); got != tt.want {
				t.Errorf("findCommonCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

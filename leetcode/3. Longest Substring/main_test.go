package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"abcabcbb"}, 3},
		{"2", args{"bbbbb"}, 1},
		{"3", args{"pwwkew"}, 3},
		{"4", args{" "}, 1},
		{"5", args{""}, 0},
		{"6", args{"a"}, 1},
		{"7", args{"ab"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

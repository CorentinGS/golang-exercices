package main

import "testing"

func Test_poorPigs(t *testing.T) {
	type args struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{
			name: "test1",
			args: args{
				buckets:       1,
				minutesToDie:  10,
				minutesToTest: 1,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				buckets:       4,
				minutesToDie:  15,
				minutesToTest: 30,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				buckets:       1000,
				minutesToDie:  15,
				minutesToTest: 60,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poorPigs(tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest); got != tt.want {
				t.Errorf("poorPigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeap_BuildHeap(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		h    MaxHeap
		args args
		want MaxHeap
	}{
		{
			name: "Test BuildHeap",
			h:    MaxHeap{},
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: MaxHeap{9, 8, 7, 4, 5, 6, 3, 2, 1},
		},
		{
			name: "Test BuildHeap",
			h:    MaxHeap{},
			args: args{
				arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: MaxHeap{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h = New(tt.args.arr)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("BuildHeap() got = %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestMaxHeap_ExtractPeek(t *testing.T) {
	tests := []struct {
		name    string
		h       MaxHeap
		want    int
		wantErr bool
	}{
		{
			name:    "Test ExtractPeek",
			h:       New([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}),
			want:    9,
			wantErr: false,
		},
		{
			name:    "Test ExtractPeek",
			h:       New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			want:    9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.ExtractPeek()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractPeek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractPeek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

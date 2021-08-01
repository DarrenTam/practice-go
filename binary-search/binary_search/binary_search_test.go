package binary_search

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"it should be able to find the index",
			args{[]int{1, 2, 5, 8, 9, 10, 50, 68, 75, 99}, 9},
			9},
		{"it should not able to find the index",
			args{[]int{1, 2, 5, 8, 9, 10, 50, 68, 75, 99}, 100},
			-1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.value); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

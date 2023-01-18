package code

import "testing"

func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxKelements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxKelements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isItPossible(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isItPossible(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("isItPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCrossingTime(t *testing.T) {
	type args struct {
		n    int
		k    int
		time [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCrossingTime(tt.args.n, tt.args.k, tt.args.time); got != tt.want {
				t.Errorf("findCrossingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxKelements1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{10, 10, 10, 10, 10}, 5}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxKelements(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxKelements() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

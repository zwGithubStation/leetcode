package code

import (
	"reflect"
	"testing"
)

func Test_countDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{7}, 1},
		{"test2", args{121}, 2},
		{"test3", args{1248}, 4},
		{"test3", args{10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.args.num); got != tt.want {
				t.Errorf("countDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distinctPrimeFactors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 4, 3, 7, 10, 6}}, 4},
		{"test2", args{[]int{2, 4, 8, 16}}, 1},
		{"test3", args{[]int{2, 4, 7, 13, 19}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctPrimeFactors(tt.args.nums); got != tt.want {
				t.Errorf("distinctPrimeFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumPartition(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"165462", 60}, 4},
		{"test2", args{"238182", 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPartition(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("minimumPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closestPrimes(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{10, 19}, []int{11, 13}},
		{"test2", args{4, 6}, []int{-1, -1}},
		{"test3", args{19, 31}, []int{29, 31}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestPrimes(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

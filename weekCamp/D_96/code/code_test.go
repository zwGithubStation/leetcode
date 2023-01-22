package code

import "testing"

func Test_getCommon(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
			if got := getCommon(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("getCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOperations(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
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
			if got := minOperations(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScore(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
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
			if got := maxScore(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReachable(t *testing.T) {
	type args struct {
		targetX int
		targetY int
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
			if got := isReachable(tt.args.targetX, tt.args.targetY); got != tt.want {
				t.Errorf("isReachable() = %v, want %v", got, tt.want)
			}
		})
	}
}

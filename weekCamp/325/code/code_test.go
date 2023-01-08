package code

import "testing"

func Test_closetTarget(t *testing.T) {
	type args struct {
		words      []string
		target     string
		startIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1}, 1},
		{"test2", args{[]string{"a", "b", "leetcode"}, "leetcode", 0}, 1},
		{"test3", args{[]string{"i", "eat", "leetcode"}, "ate", 0}, -1},
		{"test4", args{[]string{"pgmiltbptl", "jnkxwutznb", "bmeirwjars",
			"ugzyaufzzp", "pgmiltbptl", "sfhtxkmzwn",
			"pgmiltbptl", "pgmiltbptl", "onvmgvjhxa", "jyzdtwbwqp"}, "pgmiltbptl", 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closetTarget(tt.args.words, tt.args.target, tt.args.startIndex); got != tt.want {
				t.Errorf("closetTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_takeCharacters(t *testing.T) {
	type args struct {
		s string
		k int
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
			if got := takeCharacters(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumTastiness(t *testing.T) {
	type args struct {
		price []int
		k     int
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
			if got := maximumTastiness(tt.args.price, tt.args.k); got != tt.want {
				t.Errorf("maximumTastiness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPartitions(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := countPartitions(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

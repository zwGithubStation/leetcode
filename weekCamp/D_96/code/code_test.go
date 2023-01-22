package code

import "testing"

func Test_categorizeBox(t *testing.T) {
	type args struct {
		length int
		width  int
		height int
		mass   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := categorizeBox(tt.args.length, tt.args.width, tt.args.height, tt.args.mass); got != tt.want {
				t.Errorf("categorizeBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStream_Consec(t *testing.T) {
	type args struct {
		num int
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
			this := &DataStream{}
			if got := this.Consec(tt.args.num); got != tt.want {
				t.Errorf("Consec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xorBeauty(t *testing.T) {
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
			if got := xorBeauty(tt.args.nums); got != tt.want {
				t.Errorf("xorBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPower(t *testing.T) {
	type args struct {
		stations []int
		r        int
		k        int
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
			if got := maxPower(tt.args.stations, tt.args.r, tt.args.k); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

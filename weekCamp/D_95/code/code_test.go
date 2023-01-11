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

package code

func categorizeBox(length, width, height, mass int) string {
	x := length >= 1e4 || width >= 1e4 || height >= 1e4 || length*width*height >= 1e9
	y := mass >= 100
	switch {
	case x && y:
		return "Both"
	case x:
		return "Bulky"
	case y:
		return "Heavy"
	default:
		return "Neither"
	}
}

type DataStream struct {
}

func Constructor(value int, k int) DataStream {
	return DataStream{}
}

func (this *DataStream) Consec(num int) bool {
	return true
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */

func xorBeauty(nums []int) int {
	return 1
}

func maxPower(stations []int, r int, k int) int64 {
	return 1
}

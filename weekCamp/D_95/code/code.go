package code

/*
给你四个整数 length ，width ，height 和 mass ，分别表示一个箱子的三个维度和质量，请你返回一个表示箱子 类别 的字符串。
如果满足以下条件，那么箱子是 "Bulky" 的：
箱子 至少有一个 维度大于等于 10e4 。
或者箱子的 体积 大于等于 10e9 。
如果箱子的质量大于等于 100 ，那么箱子是 "Heavy" 的。
如果箱子同时是 "Bulky" 和 "Heavy" ，那么返回类别为 "Both" 。
如果箱子既不是 "Bulky" ，也不是 "Heavy" ，那么返回类别为 "Neither" 。
如果箱子是 "Bulky" 但不是 "Heavy" ，那么返回类别为 "Bulky" 。
如果箱子是 "Heavy" 但不是 "Bulky" ，那么返回类别为 "Heavy" 。
注意，箱子的体积等于箱子的长度、宽度和高度的乘积。
提示：
1 <= length, width, height <= 10e5
1 <= mass <= 10e3
*/
func categorizeBox(length, width, height, mass int) string {
	x := length >= 10e4 || width >= 10e4 || height >= 10e4 || length*width*height >= 10e9
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
	value, k, cnt int
}

func Constructor(value int, k int) DataStream {

}

func (d *DataStream) Consec(num int) bool {

}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */

func xorBeauty(nums []int) (ans int) {

}

func maxPower(stations []int, r int, k int) int64 {

}

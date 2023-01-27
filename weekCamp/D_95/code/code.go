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

/*
给你一个整数数据流，请你实现一个数据结构，检查数据流中最后 k 个整数是否 等于 给定值 value 。
请你实现 DataStream 类：
DataStream(int value, int k) 用两个整数 value 和 k 初始化一个空的整数数据流。
boolean consec(int num) 将 num 添加到整数数据流。如果后 k 个整数都等于 value ，返回 true ，否则返回 false 。如果少于 k 个整数，条件不满足，所以也返回 false 。
示例 1：
输入：
["DataStream", "consec", "consec", "consec", "consec"]
[[4, 3], [4], [4], [4], [3]]
输出：
[null, false, false, true, false]
解释：
DataStream dataStream = new DataStream(4, 3); // value = 4, k = 3
dataStream.consec(4); // 数据流中只有 1 个整数，所以返回 False 。
dataStream.consec(4); // 数据流中只有 2 个整数
                      // 由于 2 小于 k ，返回 False 。
dataStream.consec(4); // 数据流最后 3 个整数都等于 value， 所以返回 True 。
dataStream.consec(3); // 最后 k 个整数分别是 [4,4,3] 。
                      // 由于 3 不等于 value ，返回 False 。
提示：
1 <= value, num <= 10e9
1 <= k <= 10e5
至多调用 consec 次数为 10e5 次。
*/

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */

type DataStream struct {
	value, k, cnt int
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k, cnt: 0}
}

func (this *DataStream) Consec(num int) bool {
	if this.value == num {
		this.cnt++
	} else {
		this.cnt = 0
	}
	return this.k <= this.cnt
}

/*
给你一个下标从 0 开始的整数数组 nums 。
三个下标 i ，j 和 k 的 有效值 定义为 ((nums[i] | nums[j]) & nums[k]) 。
一个数组的 xor 美丽值 是数组中所有满足 0 <= i, j, k < n  的三元组 (i, j, k) 的 有效值 的异或结果。
请你返回 nums 的 xor 美丽值。
注意：
val1 | val2 是 val1 和 val2 的按位或。
val1 & val2 是 val1 和 val2 的按位与。
示例 1：
输入：nums = [1,4]
输出：5
解释：
三元组和它们对应的有效值如下：
- (0,0,0) 有效值为 ((1 | 1) & 1) = 1
- (0,0,1) 有效值为 ((1 | 1) & 4) = 0
- (0,1,0) 有效值为 ((1 | 4) & 1) = 1
- (0,1,1) 有效值为 ((1 | 4) & 4) = 4
- (1,0,0) 有效值为 ((4 | 1) & 1) = 1
- (1,0,1) 有效值为 ((4 | 1) & 4) = 4
- (1,1,0) 有效值为 ((4 | 4) & 1) = 0
- (1,1,1) 有效值为 ((4 | 4) & 4) = 4
数组的 xor 美丽值为所有有效值的按位异或 1 ^ 0 ^ 1 ^ 4 ^ 1 ^ 4 ^ 0 ^ 4 = 5 。
示例 2：
输入：nums = [15,45,20,2,34,35,5,44,32,30]
输出：34
解释：数组的 xor 美丽值为 34 。
提示：
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e9
*/
func xorBeauty(nums []int) (ans int) {
	return 1
}

func maxPower(stations []int, r int, k int) int64 {
	return 1
}

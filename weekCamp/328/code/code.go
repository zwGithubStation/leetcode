package code

/*
给你一个正整数数组 nums 。
元素和 是 nums 中的所有元素相加求和。
数字和 是 nums 中每一个元素的每一数位（重复数位需多次求和）相加求和。
返回 元素和 与 数字和 的绝对差。
注意：两个整数 x 和 y 的绝对差定义为 |x - y| 。
示例 1：
输入：nums = [1,15,6,3]
输出：9
解释：
nums 的元素和是 1 + 15 + 6 + 3 = 25 。
nums 的数字和是 1 + 1 + 5 + 6 + 3 = 16 。
元素和与数字和的绝对差是 |25 - 16| = 9 。
示例 2：
输入：nums = [1,2,3,4]
输出：0
解释：
nums 的元素和是 1 + 2 + 3 + 4 = 10 。
nums 的数字和是 1 + 2 + 3 + 4 = 10 。
元素和与数字和的绝对差是 |10 - 10| = 0 。
提示：
1 <= nums.length <= 2000
1 <= nums[i] <= 2000
*/

//手写版:各种基本函数操作
func total(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}

func diaTotal(num int) int {
	ans := 0
	for num/10 > 0 {
		ans += num % 10
		num = num / 10
	}
	ans += num
	return ans
}

func digTotals(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += diaTotal(nums[i])
	}
	return ans
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func differenceOfSum(nums []int) int {
	return abs(total(nums), digTotals(nums))
}

//灯塔版：简略极致
//由于元素值一定不小于其数位和，所以答案就是元素和减去数位和。
//代码实现时可以用同一个变量。
func differenceOfSum2(nums []int) (ans int) {
	for _, x := range nums {
		for ans += x; x > 0; x /= 10 {
			ans -= x % 10
		}
	}
	return
}

func rangeAddQueries(n int, queries [][]int) [][]int {
	return [][]int{{1}}
}

func countGood(nums []int, k int) int64 {
	return 1
}

func maxOutput(n int, edges [][]int, price []int) int64 {
	return 1
}

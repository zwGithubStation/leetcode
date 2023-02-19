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

/*
给你一个正整数 n ，表示最初有一个 n x n 、下标从 0 开始的整数矩阵 mat ，矩阵中填满了 0 。
另给你一个二维整数数组 query 。针对每个查询 query[i] = [row1i, col1i, row2i, col2i] ，请你执行下述操作：
找出 左上角 为 (row1i, col1i) 且 右下角 为 (row2i, col2i) 的子矩阵，将子矩阵中的 每个元素 加 1 。也就是给所有满足 row1i <= x <= row2i 和 col1i <= y <= col2i 的 mat[x][y] 加 1 。
返回执行完所有操作后得到的矩阵 mat 。
示例 1：
输入：n = 3, queries = [[1,1,2,2],[0,0,1,1]]
输出：[[1,1,0],[1,2,1],[0,1,1]]
解释：上图所展示的分别是：初始矩阵、执行完第一个操作后的矩阵、执行完第二个操作后的矩阵。
- 第一个操作：将左上角为 (1, 1) 且右下角为 (2, 2) 的子矩阵中的每个元素加 1 。
- 第二个操作：将左上角为 (0, 0) 且右下角为 (1, 1) 的子矩阵中的每个元素加 1 。
示例 2：
输入：n = 2, queries = [[0,0,1,1]]
输出：[[1,1],[1,1]]
解释：上图所展示的分别是：初始矩阵、执行完第一个操作后的矩阵。
- 第一个操作：将矩阵中的每个元素加 1 。
提示：
1 <= n <= 500
1 <= queries.length <= 10e4
0 <= row1i <= row2i < n
0 <= col1i <= col2i < n
*/
func rangeAddQueries(n int, queries [][]int) [][]int {
	return [][]int{{1}}
}

/*
给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。
一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
子数组 是原数组中一段连续 非空 的元素序列。
示例 1：
输入：nums = [1,1,1,1,1], k = 10
输出：1
解释：唯一的好子数组是这个数组本身。
示例 2：
输入：nums = [3,1,4,3,2,2,4], k = 2
输出：4
解释：总共有 4 个不同的好子数组：
- [3,1,4,3,2,2] 有 2 对。
- [3,1,4,3,2,2,4] 有 3 对。
- [1,4,3,2,2,4] 有 2 对。
- [4,3,2,2,4] 有 2 对。
提示：
1 <= nums.length <= 10e5
1 <= nums[i], k <= 10e9
*/
func countGood(nums []int, k int) int64 {
	return 1
}

func maxOutput(n int, edges [][]int, price []int) int64 {
	return 1
}

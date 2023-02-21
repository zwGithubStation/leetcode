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
	var ret [][]int
	for i := 0; i < n; i++ {
		cur := make([]int, n)
		ret = append(ret, cur)
	}
	for i := 0; i < len(queries); i++ {
		for j := queries[i][0]; j <= queries[i][2]; j++ {
			for k := queries[i][1]; k <= queries[i][3]; k++ {
				ret[j][k]++
			}
		}
	}
	return ret
}

//二维差分数组
func rangeAddQueries2(n int, queries [][]int) [][]int {
	// 二维差分模板
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	//差分二维数组的更新原则：
	//左上+x
	//左下的下面元素-x
	//右上的右侧元素-x
	//右下的右下元素+x
	//配合下面的复原二层循环的逻辑一起理解：
	//每一个当前要恢复结果的元素(x, y),要基于四个值进行计算：
	//(x-1, y) 已更新为结果值
	//(x, y-1) 已更新为结果值
	//(x-1, y-1) 已更新为结果值
	//(x, y)上的diff值 目前记录着差分值
	update := func(r1, c1, r2, c2, x int) {
		diff[r1+1][c1+1] += x
		diff[r1+1][c2+2] -= x //超过n的列上的计数在最终计算前缀和时 不起作用(见下面的复原二层循环) 这里将数组从(n+1，n+1)扩大为(n+2, n+2) 只是为了update里面少写判断语句
		diff[r2+2][c1+1] -= x //超过n的行上的计数在最终计算前缀和时 不起作用(见下面的复原二层循环) 这里将数组从(n+1，n+1)扩大为(n+2, n+2) 只是为了update里面少写判断语句
		diff[r2+2][c2+2] += x //超过n的行列上的计数在最终计算前缀和时 不起作用(见下面的复原二层循环) 这里将数组从(n+1，n+1)扩大为(n+2, n+2) 只是为了update里面少写判断语句
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原（原地修改）
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
		}
	}
	// 保留中间 n*n 的部分，即为答案
	diff = diff[1 : n+1]
	for i, row := range diff {
		diff[i] = row[1 : n+1] //注意这里的写法喔， 不能写成row = row[1: n+1]
	}
	return diff
}

//二维前缀和 前置题目：https://leetcode.cn/problems/range-sum-query-2d-immutable/
/*
给定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：
NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
示例 1：
输入:
["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出:
[null, 8, 11, 12]
解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-105 <= matrix[i][j] <= 10e5
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
最多调用 10e4 次 sumRegion 方法
*/

type NumMatrix struct {
	diff [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rowLen := len(matrix)
	if rowLen == 0 {
		return NumMatrix{}
	}
	colLen := len(matrix[0])
	var ret NumMatrix
	ret.diff = make([][]int, rowLen+1)
	for i, _ := range ret.diff {
		ret.diff[i] = make([]int, colLen+1)
	}

	for i := 1; i <= rowLen; i++ {
		for j := 1; j <= colLen; j++ {
			//构造前缀和的方式
			ret.diff[i][j] = matrix[i-1][j-1] + ret.diff[i-1][j] + ret.diff[i][j-1] - ret.diff[i-1][j-1]
		}
	}
	return ret
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	//由前缀和求区域总和的方式
	return this.diff[row2+1][col2+1] - this.diff[row2+1][col1] - this.diff[row1][col2+1] + this.diff[row1][col1]
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
func countGood(nums []int, k int) (ans int64) {
	cnt := make(map[int]int)
	left := 0
	pairs := 0
	for _, x := range nums {
		pairs += cnt[x] //右侧每出现一个新的元素x，pairs的计数就增加cnt[x](前面已经有多少个x，就会多出多少个新的[x，x]对)
		cnt[x]++
		for pairs-cnt[nums[left]]+1 >= k { //左侧尝试去掉冗余元素
			cnt[nums[left]]--
			pairs -= cnt[nums[left]] // 每少一个y 就要减少(cnt[y]-1)个pairs
			left++
		}
		if pairs >= k {
			ans += int64(left + 1) //一个没法从左侧缩短的[left，right]字串，可以提供出left+1个答案
		}
	}
	return ans
}

/*
同向双指针 滑动窗口 前置题目1：
https://leetcode.cn/problems/minimum-size-subarray-sum/
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：
输入：target = 4, nums = [1,4,4]
输出：1
示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
提示：
1 <= target <= 10e9
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e5
*/

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, sum, ans := 0, 0, n+1
	for right, x := range nums { //单调性， 由满足条件逐步到不满足条件
		sum += x
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

/*
同向双指针 滑动窗口 前置题目2：
https://leetcode.cn/problems/subarray-product-less-than-k/
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
示例 1：
输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
示例 2：
输入：nums = [1,2,3], k = 0
输出：0
提示:
1 <= nums.length <= 3 * 10e4
1 <= nums[i] <= 1000
0 <= k <= 10e6
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, prod, ans := 0, 1, 0
	if k <= 1 {
		return 0
	}
	for right, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1 //注意这里如何保持记录的不重复性：每次进行记录的right值是唯一的； [left， right]间的，以righ结尾的子数组数目为right-left+1
	}
	return ans
}

/*
同向双指针 滑动窗口 前置题目3：
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
提示：
0 <= s.length <= 5 * 10e4
s 由英文字母、数字、符号和空格组成
*/

func lengthOfLongestSubstring(s string) int {
	left, ans := 0, 0
	cnt := make([]int, 128)
	for right, x := range s {
		cnt[x]++
		for cnt[x] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
给你一个 n 个节点的无向无根图，节点编号为 0 到 n - 1 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间有一条边。
每个节点都有一个价值。给你一个整数数组 price ，其中 price[i] 是第 i 个节点的价值。
一条路径的 价值和 是这条路径上所有节点的价值之和。
你可以选择树中任意一个节点作为根节点 root 。选择 root 为根的 开销 是以 root 为起点的所有路径中，价值和 最大的一条路径与最小的一条路径的差值。
请你返回所有节点作为根节点的选择中，最大 的 开销 为多少。
示例 1：
输入：n = 6, edges = [[0,1],[1,2],[1,3],[3,4],[3,5]], price = [9,8,7,6,10,5]
输出：24
解释：上图展示了以节点 2 为根的树。左图（红色的节点）是最大价值和路径，右图（蓝色的节点）是最小价值和路径。
- 第一条路径节点为 [2,1,3,4]：价值为 [7,8,6,10] ，价值和为 31 。
- 第二条路径节点为 [2] ，价值为 [7] 。
最大路径和与最小路径和的差值为 24 。24 是所有方案中的最大开销。
示例 2：
输入：n = 3, edges = [[0,1],[1,2]], price = [1,1,1]
输出：2
解释：上图展示了以节点 0 为根的树。左图（红色的节点）是最大价值和路径，右图（蓝色的节点）是最小价值和路径。
- 第一条路径包含节点 [0,1,2]：价值为 [1,1,1] ，价值和为 3 。
- 第二条路径节点为 [0] ，价值为 [1] 。
最大路径和与最小路径和的差值为 2 。2 是所有方案中的最大开销。
提示：
1 <= n <= 10e5
edges.length == n - 1
0 <= ai, bi <= n - 1
edges 表示一棵符合题面要求的树。
price.length == n
1 <= price[i] <= 10e5
*/
func maxOutput(n int, edges [][]int, price []int) int64 {
	return 1
}

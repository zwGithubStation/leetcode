package code

/*
给你一个整数 num 。你知道 Danny Mittal 会偷偷将 0 到 9 中的一个数字 替换 成另一个数字。
请你返回将 num 中 恰好一个 数字进行替换后，得到的最大值和最小值的差位多少。
注意：
当 Danny 将一个数字 d1 替换成另一个数字 d2 时，Danny 需要将 nums 中所有 d1 都替换成 d2 。
Danny 可以将一个数字替换成它自己，也就是说 num 可以不变。
Danny 可以将数字分别替换成两个不同的数字分别得到最大值和最小值。
替换后得到的数字可以包含前导 0 。
Danny Mittal 获得周赛 326 前 10 名，让我们恭喜他。
示例 1：
输入：num = 11891
输出：99009
解释：
为了得到最大值，我们将数字 1 替换成数字 9 ，得到 99899 。
为了得到最小值，我们将数字 1 替换成数字 0 ，得到 890 。
两个数字的差值为 99009 。
示例 2：
输入：num = 90
输出：99
解释：
可以得到的最大值是 99（将 0 替换成 9），最小值是 0（将 9 替换成 0）。
所以我们得到 99 。
提示：
1 <= num <= 10e8
*/

func minMaxDifference(num int) int {
	return 0
}

/*
给你一个下标从 0 开始的整数数组 nums 。
nums 的 最小 得分是满足 0 <= i < j < nums.length 的 |nums[i] - nums[j]| 的最小值。
nums的 最大 得分是满足 0 <= i < j < nums.length 的 |nums[i] - nums[j]| 的最大值。
nums 的分数是 最大 得分与 最小 得分的和。
我们的目标是最小化 nums 的分数。你 最多 可以修改 nums 中 2 个元素的值。
请你返回修改 nums 中 至多两个 元素的值后，可以得到的 最小分数 。
|x| 表示 x 的绝对值。

示例 1：
输入：nums = [1,4,3]
输出：0
解释：将 nums[1] 和 nums[2] 的值改为 1 ，nums 变为 [1,1,1] 。|nums[i] - nums[j]| 的值永远为 0 ，所以我们返回 0 + 0 = 0 。
示例 2：
输入：nums = [1,4,7,8,5]
输出：3
解释：
将 nums[0] 和 nums[1] 的值变为 6 ，nums 变为 [6,6,7,8,5] 。
最小得分是 i = 0 且 j = 1 时得到的 |nums[i] - nums[j]| = |6 - 6| = 0 。
最大得分是 i = 3 且 j = 4 时得到的 |nums[i] - nums[j]| = |8 - 5| = 3 。
最大得分与最小得分之和为 3 。这是最优答案。
提示：
3 <= nums.length <= 10e5
1 <= nums[i] <= 10e9
*/

func minimizeSum(nums []int) int {
	return 0
}

/*
你一个下标从 0 开始的整数数组 nums 。
如果存在一些整数满足 0 <= index1 < index2 < ... < indexk < nums.length ，得到 nums[index1] | nums[index2] | ... | nums[indexk] = x ，那么我们说 x 是 可表达的 。换言之，如果一个整数能由 nums 的某个子序列的或运算得到，那么它就是可表达的。
请你返回 nums 不可表达的 最小非零整数 。
示例 1：
输入：nums = [2,1]
输出：4
解释：1 和 2 已经在数组中，因为 nums[0] | nums[1] = 2 | 1 = 3 ，所以 3 是可表达的。由于 4 是不可表达的，所以我们返回 4 。
示例 2：
输入：nums = [5,3,2]
输出：1
解释：1 是最小不可表达的数字。
提示：
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e9
*/

func minImpossibleOR(nums []int) int {
	return 0
}

/*
给你一个下标从 0 开始的 m x n 二进制 矩阵 grid 。你可以从一个格子 (row, col) 移动到格子 (row + 1, col) 或者 (row, col + 1) ，
前提是前往的格子值为 1 。如果从 (0, 0) 到 (m - 1, n - 1) 没有任何路径，我们称该矩阵是 不连通 的。
你可以翻转 最多一个 格子的值（也可以不翻转）。你 不能翻转 格子 (0, 0) 和 (m - 1, n - 1) 。
如果可以使矩阵不连通，请你返回 true ，否则返回 false 。
注意 ，翻转一个格子的值，可以使它的值从 0 变 1 ，或从 1 变 0 。
示例 1：
输入：grid = [[1,1,1],[1,0,0],[1,1,1]]
输出：true
解释：按照上图所示我们翻转蓝色格子里的值，翻转后从 (0, 0) 到 (2, 2) 没有路径。
示例 2：
输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
输出：false
解释：无法翻转至多一个格子，使 (0, 0) 到 (2, 2) 没有路径。
提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 1000
1 <= m * n <= 10e5
grid[0][0] == grid[m - 1][n - 1] == 1
*/

func isPossibleToCutPath(grid [][]int) bool {
	return false
}

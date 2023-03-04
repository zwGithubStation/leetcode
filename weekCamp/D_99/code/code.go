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
你一个下标从 0 开始的整数数组 nums 。
如果存在一些整数满足 0 <= index1 < index2 < ... < indexk < nums.length ，得到 nums[index1] | nums[index2] | ... | nums[indexk] = x ，那么我们说 x 是 可表达的 。换言之，如果一个整数能由 nums 的某个子序列的或运算得到，那么它就是可表达的。
请你返回 nums 不可表达的 最小非零整数 。
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
给你两个下标从 0 开始的数组 nums1 和 nums2 ，和一个二维数组 queries 表示一些操作。总共有 3 种类型的操作：
操作类型 1 为 queries[i] = [1, l, r] 。你需要将 nums1 从下标 l 到下标 r 的所有 0 反转成 1 或将 1 反转成 0 。l 和 r 下标都从 0 开始。
操作类型 2 为 queries[i] = [2, p, 0] 。对于 0 <= i < n 中的所有下标，令 nums2[i] = nums2[i] + nums1[i] * p 。
操作类型 3 为 queries[i] = [3, 0, 0] 。求 nums2 中所有元素的和。
请你返回一个数组，包含所有第三种操作类型的答案。
示例 1：
输入：nums1 = [1,0,1], nums2 = [0,0,0], queries = [[1,1,1],[2,1,0],[3,0,0]]
输出：[3]
解释：第一个操作后 nums1 变为 [1,1,1] 。第二个操作后，nums2 变成 [1,1,1] ，所以第三个操作的答案为 3 。所以返回 [3] 。
示例 2：
输入：nums1 = [1], nums2 = [5], queries = [[2,0,0],[3,0,0]]
输出：[5]
解释：第一个操作后，nums2 保持不变为 [5] ，所以第二个操作的答案是 5 。所以返回 [5] 。
提示：
1 <= nums1.length,nums2.length <= 10e5
nums1.length = nums2.length
1 <= queries.length <= 10e5
queries[i].length = 3
0 <= l <= r <= nums1.length - 1
0 <= p <= 10e6
0 <= nums1[i] <= 1
0 <= nums2[i] <= 10e9
*/

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	return []int64{}
}

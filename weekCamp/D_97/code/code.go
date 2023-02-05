package code

/*
给你一个正整数数组 nums ，请你返回一个数组 answer ，你需要将 nums 中每个整数进行数位分割后，按照 nums 中出现的 相同顺序 放入答案数组中。
对一个整数进行数位分割，指的是将整数各个数位按原本出现的顺序排列成数组。
比方说，整数 10921 ，分割它的各个数位得到 [1,0,9,2,1] 。
示例 1：
输入：nums = [13,25,83,77]
输出：[1,3,2,5,8,3,7,7]
解释：
- 分割 13 得到 [1,3] 。
- 分割 25 得到 [2,5] 。
- 分割 83 得到 [8,3] 。
- 分割 77 得到 [7,7] 。
answer = [1,3,2,5,8,3,7,7] 。answer 中的数字分割结果按照原数字在数组中的相同顺序排列。
示例 2：
输入：nums = [7,1,3,9]
输出：[7,1,3,9]
解释：nums 中每个整数的分割是它自己。
answer = [7,1,3,9] 。
提示：
1 <= nums.length <= 1000
1 <= nums[i] <= 10e5
*/

func separateDigits(nums []int) []int {
	return []int{}
}

/*
给你一个整数数组 banned 和两个整数 n 和 maxSum 。你需要按照以下规则选择一些整数：
被选择整数的范围是 [1, n] 。
每个整数 至多 选择 一次 。
被选择整数不能在数组 banned 中。
被选择整数的和不超过 maxSum 。
请你返回按照上述规则 最多 可以选择的整数数目。
示例 1：
输入：banned = [1,6,5], n = 5, maxSum = 6
输出：2
解释：你可以选择整数 2 和 4 。
2 和 4 在范围 [1, 5] 内，且它们都不在 banned 中，它们的和是 6 ，没有超过 maxSum 。
示例 2：
输入：banned = [1,2,3,4,5,6,7], n = 8, maxSum = 1
输出：0
解释：按照上述规则无法选择任何整数。
示例 3：
输入：banned = [11], n = 7, maxSum = 50
输出：7
解释：你可以选择整数 1, 2, 3, 4, 5, 6 和 7 。
它们都在范围 [1, 7] 中，且都没出现在 banned 中，它们的和是 28 ，没有超过 maxSum 。
提示：
1 <= banned.length <= 10e4
1 <= banned[i], n <= 10e4
1 <= maxSum <= 10e9
*/

func maxCount(banned []int, n int, maxSum int) int {
	return 1
}

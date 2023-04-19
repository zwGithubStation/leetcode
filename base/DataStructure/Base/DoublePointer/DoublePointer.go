package DoublePointer

//使用双指针的条件：
//
// 单调性(从满足要求到不满足要求 or 从不满足要求到满足要求)

/*
209 长度最小的子数组
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
进阶：
如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
*/

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	left, ans, sum := 0, length+1, 0

	//时间复杂度为0(N)的二重循环！！！！！！！！！！！！！！！！！！！！！！
	for right, x := range nums {
		sum += x //加右
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left] //去左
			left++
		}
	}

	if ans <= length {
		return ans
	}
	return 0
}

/*
713 乘积小于 K 的子数组
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

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	prob, left := 1, 0
	for right, x := range nums {
		prob *= x
		for prob >= k {
			prob /= nums[left]
			left++
		}
		ans += right - left + 1 //记录right固定，left变化的子数字个数 不会重复
	}
	return
}

/*
3. 无重复字符的最长子串
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
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	cnt := [128]int{} // 也可以用 map[byte]int，这里为了方便用的数组
	//下面的博客介绍了for range string 时每个元素的类型(rune=int32；根据实际的字符类型【UTF-8规范中一个字符可以用1-4个字节表示】)
	//https://berryjam.github.io/2018/03/%E4%BB%8Egolang%E5%AD%97%E7%AC%A6%E4%B8%B2string%E9%81%8D%E5%8E%86%E8%AF%B4%E8%B5%B7/
	//https://golangbyexample.com/length-of-string-golang/
	for right, c := range s {
		cnt[c]++
		for cnt[c] > 1 { // 不满足要求
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

/*
1004. 最大连续1的个数 III
https://leetcode.cn/problems/max-consecutive-ones-iii/

给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
示例 1：
输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：
输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
提示：
1 <= nums.length <= 10e5
nums[i] 不是 0 就是 1
0 <= k <= nums.length
*/

func longestOnes(nums []int, k int) (ans int) {
	left, cntZero := 0, 0
	for right, x := range nums {
		cntZero += 1 - x
		for cntZero > k {
			cntZero -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

/*
1234. 替换子串得到平衡字符串
https://leetcode.cn/problems/replace-the-substring-for-balanced-string/

有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
请返回待替换子串的最小可能长度。
如果原字符串自身就是一个平衡字符串，则返回 0。
示例 1：
输入：s = "QWER"
输出：0
解释：s 已经是平衡的了。
示例 2：
输入：s = "QQWE"
输出：1
解释：我们需要把一个 'Q' 替换成 'R'，这样得到的 "RQWE" (或 "QRWE") 是平衡的。
示例 3：
输入：s = "QQQW"
输出：2
解释：我们可以把前面的 "QQ" 替换成 "ER"。
示例 4：
输入：s = "QQQQ"
输出：3
解释：我们可以替换后 3 个 'Q'，使 s = "QWER"。
提示：
1 <= s.length <= 10^5
s.length 是 4 的倍数
s 中只含有 'Q', 'W', 'E', 'R' 四种字符
*/

/*
根据题意，如果在待替换子串之外的任意字符的出现次数超过 m(len/4)，那么无论怎么替换，都无法使这个字符的出现次数等于 m。

反过来说，如果在待替换子串之外的任意字符的出现次数都不超过 m，那么必然可以通过一定地替换，使 s 为平衡字符串，即每个字符的出现次数均为 m。

对于本题，设子串的左右端点为 left 和 right，枚举 right，如果子串外的任意字符的出现次数都不超过 m，则说明从 left 到 right 的这段子串可以是待替换子串，用其长度
right−left+1 更新答案的最小值，并向右移动 left，缩小子串长度。
*/

func balancedString(s string) (ans int) {
	m := len(s) / 4
	cnt, left := ['X']int{}, 0

	for _, x := range s {
		cnt[x]++
	}

	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return
	}

	ans = len(s) //ans预设值要是最大值，因为求最小！！！！！！！
	for right, x := range s {
		cnt[x]--
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-left+1)
			cnt[s[left]]++
			left++
		}
	}
	return
}

/*
1658. 将 x 减到 0 的最小操作数
https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/

给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。
如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
示例 1：
输入：nums = [1,1,4,2,3], x = 5
输出：2
解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
示例 2：
输入：nums = [5,6,7,8,9], x = 4
输出：-1
示例 3：
输入：nums = [3,2,20,1,1,3], x = 10
输出：5
解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。
提示：
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e4
1 <= x <= 10e9
*/

func minOperations(nums []int, x int) int {
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

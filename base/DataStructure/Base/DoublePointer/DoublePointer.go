package DoublePointer

import (
	"sort"
)

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

/*
把问题转换成「从 nums 中移除一个最长的子数组，使得剩余元素的和为 x」。

换句话说，要从 nums 中找最长的子数组，其元素和等于 s−x，这里 s 为 nums 所有元素之和。
*/

func minOperations(nums []int, x int) int {
	curSum, sum, left := 0, 0, 0
	for _, v := range nums {
		sum += v
	}

	if sum < x {
		return -1
	}

	ans := -1 //ans 不能等于0喔！！！
	for right, v := range nums {
		curSum += v
		for curSum > sum-x {
			curSum -= nums[left]
			left++
		}
		if curSum == sum-x {
			ans = max(ans, right-left+1)
		}
	}
	if ans < 0 { //ans是可以为0的喔！！！ 就是sum == x的情形
		return -1
	}
	return len(nums) - ans
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

/*
167. 两数之和 II - 输入有序数组
https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

给你一个下标从 1 开始的整数数组  numbers ，该数组已按 非递减顺序排列   ，请你从数组中找出满足相加之和等于目标数  target 的两个数。
如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
你所设计的解决方案必须只使用常量级的额外空间。
示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。
示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
提示：
2 <= numbers.length <= 3 * 10e4
-1000 <= numbers[i] <= 1000
numbers 按 非递减顺序 排列
-1000 <= target <= 1000
仅存在一个有效答案
*/

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
}

/*
15. 三数之和
https://leetcode.cn/problems/3sum/

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：
输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。
提示：
3 <= nums.length <= 3000
-10e5 <= nums[i] <= 10e5
*/

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	s := len(nums)
	for i, x := range nums[:s-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if x+nums[s-2]+nums[s-1] < 0 {
			continue
		}
		j, k := i+1, s-1
		for j < k {
			sum := x + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return
}

/*
11. 盛最多水的容器
https://leetcode.cn/problems/container-with-most-water/

给定一个长度为 n 的整数数组  height  。有  n  条垂线，第 i 条线的两个端点是  (i, 0)  和  (i, height[i])  。
找出其中的两条线，使得它们与  x  轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。
示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为  49。
示例 2：
输入：height = [1,1]
输出：1
提示：
n == height.length
2 <= n <= 10e5
0 <= height[i] <= 10e4
*/

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] { //移动短板 后续面积可能增大 移动长板 则后续面积一定变小
			left++
		} else {
			right--
		}
	}
	return
}

/*
42. 接雨水
https://leetcode.cn/problems/trapping-rain-water/

给定  n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：
输入：height = [4,2,0,3,2,5]
输出：9
提示：
n == height.length
1 <= n <= 2 * 10e4
0 <= height[i] <= 10e5
*/

func trap(height []int) (ans int) {
	length := len(height)
	preMax := make([]int, length)
	preMax[0] = height[0]
	for i := 1; i < length; i++ {
		preMax[i] = max(height[i], preMax[i-1])
	}

	sufMax := make([]int, length)
	sufMax[length-1] = height[length-1] //位次不要乱了
	for i := length - 1; i >= 0; i-- {
		sufMax[i] = max(height[i], sufMax[i+1])
	}

	for i, v := range height {
		ans += min(sufMax[i], preMax[i]) - v
	}
	return ans
}

// 节省数组空间
func trap2(height []int) (ans int) {
	preMax, sufMax, left, right := 0, 0, 0, len(height)-1

	for left < right {
		preMax = max(preMax, height[left])  //对于当前的left指针， 其左边界木板高度是确定的
		sufMax = max(sufMax, height[right]) //对于当前的right指针，其右边界木板的高度是确定的
		if preMax < sufMax {
			ans += preMax - height[left] //left指针位置的水位最大可能高度确定
			left++
		} else {
			ans += sufMax - height[right] //right指针位置的水位最大可能高度确定
			right--
		}
	}
	return
}

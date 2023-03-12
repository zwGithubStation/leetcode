package code

import "sort"

// 给你一个下标从 0 开始的 环形 字符串数组 words 和一个字符串 target 。环形数组 意味着数组首尾相连。
//
// 形式上， words[i] 的下一个元素是 words[(i + 1) % n] ，而 words[i] 的前一个元素是 words[(i - 1 + n) % n] ，其中 n 是 words 的长度。
// 从 startIndex 开始，你一次可以用 1 步移动到下一个或者前一个单词。
// 返回到达目标字符串 target 所需的最短距离。如果 words 中不存在字符串 target ，返回 -1 。
func closetTarget(words []string, target string, startIndex int) int {
	var length int = len(words)
	if words[startIndex] == target {
		return 0
	}
	for i := 1; i <= length/2; i++ {
		if words[(startIndex+i)%length] == target {
			return i
		}
		if words[(startIndex-i+length)%length] == target {
			return i
		}
	}
	return -1
}

/*
给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。
你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。
示例 1：
输入：s = "aabaaaacaabc", k = 2
输出：8
解释：
从 s 的左侧取三个字符，现在共取到两个字符 'a' 、一个字符 'b' 。
从 s 的右侧取五个字符，现在共取到四个字符 'a' 、两个字符 'b' 和两个字符 'c' 。
共需要 3 + 5 = 8 分钟。
可以证明需要的最少分钟数是 8 。
示例 2：
输入：s = "a", k = 1
输出：-1
解释：无法取到一个字符 'b' 或者 'c'，所以返回 -1 。
提示：
1 <= s.length <= 10e5
s 仅由字母 'a'、'b'、'c' 组成
0 <= k <= s.length

解题思路：
无论每次如何在左右间摇摆选取，最终的结果，肯定是左边一个前缀，右边一个后缀，合起来是最终结果
如此则可以使用枚举的思路，
首先：
前缀 = null, 后缀 =A， 易知后缀的选取方式是从后往前枚举； 当前最短距离 = length(A)
前缀 = B1(只包含首元素), 后缀 =A1，A1应该是A的一个子集，考虑到A必须是从右往左构造出来，A1只能从左往右对A进行缩短！！  当前最短距离 = min(length(A), length(B1)+length(A1))
......
每次前后缀由x到x+1的剪枝策略：
我们记录上一次的counter计数(a多少个，b多少个，c多少个。初始时都是k个);
Bx到Bx+1,显然是让s[x+1]的元素计数++了; 此时我们对Ax进行缩减:
Ax从左往右开始，每个元素但凡计数>k，即可被从Ax中移除，计数-- =====> Ax+1的起始元素的计数必为2，无法再进行删减，前缀Bx+1和后缀Ax+1共同构成当前的最短距离，与上一步的最短距离比较更新ans
如此反复直至A已经被删减为null，即得到最终的结果！
*/
func min(i int, j int) int {
	if i <= j {
		return i
	}
	return j
}

// right初始设置为length-1 还是length，会影响用例的通过情况的喔（改成注释的那段代码，看看哪些用例跑不通！！）
func takeCharacters(s string, k int) int {
	length := len(s)
	right, counter := length, [3]int{} //关系连续的字符的键值存储就可以用数组
	for counter[0] < k || counter[1] < k || counter[2] < k {
		if right == 0 { //可以包含len=0的情况
			return -1
		}
		right-- //先--再算喔
		counter[s[right]-'a']++
	}
	/*
		right, counter := length-1, [3]int{}
		for counter[0] < k || counter[1] < k || counter[2] < k {
			if right == 0 {
				return -1
			}
			counter[s[right]-'a']++
			right--
		}
	*/
	ans := length - right
	for left := 0; left < length && right < length; left++ {
		counter[s[left]-'a']++
		for right < length && counter[s[right]-'a'] > k {
			counter[s[right]-'a']--
			right++
		}
		ans = min(ans, left+1+length-right) //要放在下面的判断语句之前喔
		if right >= length-1 {
			break
		}
	}
	return ans
}

/*给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。
商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。
返回礼盒的 最大 甜蜜度。
*/
//sort.Search(x int, f func(x int) bool) int
//search to find and return the smallest index i in [0, n) at which f(i) is true.
//means that f is false for some prefix of the input range [0, n) and then true for
//the remainder.
//https://pkg.go.dev/sort#Search
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search(price[len(price)-1], func(dist int) bool {
		cnt, x0 := 1, price[0]
		for _, x := range price[1:] {
			if x-x0 >= dist {
				cnt++
				x0 = x
			}
		}
		return cnt < k
	}) - 1
}

// pass暂时
func countPartitions(nums []int, k int) int {
	return 1
}

package code

import "sort"

/*
给你一个正整数 n 。n 中的每一位数字都会按下述规则分配一个符号：

最高有效位 上的数字分配到 正 号。
剩余每位上数字的符号都与其相邻数字相反。
返回所有数字及其对应符号的和。
示例 1：
输入：n = 521
输出：4
解释：(+5) + (-2) + (+1) = 4
示例 2：
输入：n = 111
输出：1
解释：(+1) + (-1) + (+1) = 1
示例 3：
输入：n = 886996
输出：0
解释：(+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0
提示：
1 <= n <= 10e9
*/
func alternateDigitSum(n int) (ans int) {
	sign := 1
	for ; n > 0; n = n / 10 {
		ans += n % 10 * sign
		sign = -sign
	}
	ans = ans * (-sign)
	return
}

/*
班里有 m 位学生，共计划组织 n 场考试。给你一个下标从 0 开始、大小为 m x n 的整数矩阵 score ，
其中每一行对应一位学生，而 score[i][j] 表示第 i 位学生在第 j 场考试取得的分数。矩阵 score 包含的整数互不相同。
另给你一个整数 k 。请你按第 k 场考试分数从高到低完成对这些学生（矩阵中的行）的排序。
返回排序后的矩阵。
示例 1：
输入：score = [[10,6,9,1],[7,5,11,2],[4,8,3,15]], k = 2
输出：[[7,5,11,2],[10,6,9,1],[4,8,3,15]]
解释：在上图中，S 表示学生，E 表示考试。
- 下标为 1 的学生在第 2 场考试取得的分数为 11 ，这是考试的最高分，所以 TA 需要排在第一。
- 下标为 0 的学生在第 2 场考试取得的分数为 9 ，这是考试的第二高分，所以 TA 需要排在第二。
- 下标为 2 的学生在第 2 场考试取得的分数为 3 ，这是考试的最低分，所以 TA 需要排在第三。
示例 2：
输入：score = [[3,4],[5,6]], k = 0
输出：[[5,6],[3,4]]
解释：在上图中，S 表示学生，E 表示考试。
- 下标为 1 的学生在第 0 场考试取得的分数为 5 ，这是考试的最高分，所以 TA 需要排在第一。
- 下标为 0 的学生在第 0 场考试取得的分数为 3 ，这是考试的最低分，所以 TA 需要排在第二。
提示：
m == score.length
n == score[i].length
1 <= m, n <= 250
1 <= score[i][j] <= 10e5
score 由 不同 的整数组成
0 <= k < n

*/
func sortTheStudents(score [][]int, k int) [][]int {
	//写法一：二维数组按照一维数组的逻辑排序 语法都是一样的
	sort.Slice(score, func(x, y int) bool { return score[x][k] > score[y][k] })
	return score
}
func sortTheStudents2(score [][]int, k int) [][]int {
	//写法二：直观方式(仰仗每个元素不重复的性质 需要额外的空间)
	type pair struct {
		x, y int
	}
	pairs := make([]pair, len(score))
	for i, x := range score {
		pairs[i].x, pairs[i].y = i, x[k]
	}
	sort.Slice(pairs, func(x, y int) bool { return pairs[x].y > pairs[y].y })
	ans := [][]int{}
	for _, x := range pairs {
		ans = append(ans, score[x.x])
	}
	return ans
}

/*
给你两个下标从 0 开始的 二元 字符串 s 和 target ，两个字符串的长度均为 n 。你可以对 s 执行下述操作 任意 次：
选择两个 不同 的下标 i 和 j ，其中 0 <= i, j < n 。
同时，将 s[i] 替换为 (s[i] OR s[j]) ，s[j] 替换为 (s[i] XOR s[j]) 。
例如，如果 s = "0110" ，你可以选择 i = 0 和 j = 2，然后同时将 s[0] 替换为 (s[0] OR s[2] = 0 OR 1 = 1)，
并将 s[2] 替换为 (s[0] XOR s[2] = 0 XOR 1 = 1)，最终得到 s = "1110" 。
如果可以使 s 等于 target ，返回 true ，否则，返回 false 。
示例 1：
输入：s = "1010", target = "0110"
输出：true
解释：可以执行下述操作：
- 选择 i = 2 和 j = 0 ，得到 s = "0010".
- 选择 i = 2 和 j = 1 ，得到 s = "0110".
可以使 s 等于 target ，返回 true 。
示例 2：
输入：s = "11", target = "00"
输出：false
解释：执行任意次操作都无法使 s 等于 target 。
提示：
n == s.length == target.length
2 <= n <= 10e5
s 和 target 仅由数字 0 和 1 组成
*/
func makeStringsEqual(s string, target string) bool {
	//pass for now
	return false
}

/*
给你一个整数数组 nums 和一个整数 k 。
将数组拆分成一些非空子数组。拆分的 代价 是每个子数组中的 重要性 之和。
令 trimmed(subarray) 作为子数组的一个特征，其中所有仅出现一次的数字将会被移除。
例如，trimmed([3,1,2,4,3,4]) = [3,4,3,4] 。
子数组的 重要性 定义为 k + trimmed(subarray).length 。
例如，如果一个子数组是 [1,2,3,3,3,4,4] ，trimmed([1,2,3,3,3,4,4]) = [3,3,3,4,4] 。这个子数组的重要性就是 k + 5 。
找出并返回拆分 nums 的所有可行方案中的最小代价。
子数组 是数组的一个连续 非空 元素序列。
示例 1：
输入：nums = [1,2,1,2,1,3,3], k = 2
输出：8
解释：将 nums 拆分成两个子数组：[1,2], [1,2,1,3,3]
[1,2] 的重要性是 2 + (0) = 2 。
[1,2,1,3,3] 的重要性是 2 + (2 + 2) = 6 。
拆分的代价是 2 + 6 = 8 ，可以证明这是所有可行的拆分方案中的最小代价。
示例 2：
输入：nums = [1,2,1,2,1], k = 2
输出：6
解释：将 nums 拆分成两个子数组：[1,2], [1,2,1] 。
[1,2] 的重要性是 2 + (0) = 2 。
[1,2,1] 的重要性是 2 + (2) = 4 。
拆分的代价是 2 + 4 = 6 ，可以证明这是所有可行的拆分方案中的最小代价。
示例 3：
输入：nums = [1,2,1,2,1], k = 5
输出：10
解释：将 nums 拆分成一个子数组：[1,2,1,2,1].
[1,2,1,2,1] 的重要性是 5 + (3 + 2) = 10 。
拆分的代价是 10 ，可以证明这是所有可行的拆分方案中的最小代价。
提示：
1 <= nums.length <= 1000
0 <= nums[i] < nums.length
1 <= k <= 10e9
*/
func minCost(nums []int, k int) int {
	//pass for now
	return 1
}

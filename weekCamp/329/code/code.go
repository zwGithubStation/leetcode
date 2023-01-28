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
func sortTheStudents(score [][]int, k int) [][]int {
    kth := make([]int, len(score))
    mm := make(map[int]int)
    for i, s := range score {
        kth[i] = s[k]
        mm[s[k]] = i
    }
    sort.Ints(kth)
    sco := [][]int{}
    for i := len(kth)-1; i >= 0; i-- {
        sco = append(sco, score[mm[kth[i]]])
    }
    return sco
}
*/
func makeStringsEqual(s string, target string) bool {

}

func minCost(nums []int, k int) int {

}

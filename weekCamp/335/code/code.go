package code

/*
n 个人站成一排，按从 1 到 n 编号。

最初，排在队首的第一个人拿着一个枕头。每秒钟，拿着枕头的人会将枕头传递给队伍中的下一个人。一旦枕头到达队首或队尾，传递方向就会改变，队伍会继续沿相反方向传递枕头。

例如，当枕头到达第 n 个人时，TA 会将枕头传递给第 n - 1 个人，然后传递给第 n - 2 个人，依此类推。
给你两个正整数 n 和 time ，返回 time 秒后拿着枕头的人的编号。



示例 1：

输入：n = 4, time = 5
输出：2
解释：队伍中枕头的传递情况为：1 -> 2 -> 3 -> 4 -> 3 -> 2 。
5 秒后，枕头传递到第 2 个人手中。
示例 2：

输入：n = 3, time = 2
输出：3
解释：队伍中枕头的传递情况为：1 -> 2 -> 3 。
2 秒后，枕头传递到第 3 个人手中。


提示：

2 <= n <= 1000
1 <= time <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/pass-the-pillow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func leftRigthDifference(nums []int) []int {
	return []int{}
}

/*
给你一个下标从 0 开始的字符串 word ，长度为 n ，由从 0 到 9 的数字组成。另给你一个正整数 m 。
word 的 可整除数组 div  是一个长度为 n 的整数数组，并满足：
如果 word[0,...,i] 所表示的 数值 能被 m 整除，div[i] = 1
否则，div[i] = 0
返回 word 的可整除数组。
示例 1：
输入：word = "998244353", m = 3
输出：[1,1,0,0,0,1,1,0,0]
解释：仅有 4 个前缀可以被 3 整除："9"、"99"、"998244" 和 "9982443" 。
示例 2：
输入：word = "1010", m = 10
输出：[0,1,0,1]
解释：仅有 2 个前缀可以被 10 整除："10" 和 "1010" 。
提示：
1 <= n <= 10e5
word.length == n
word 由数字 0 到 9 组成
1 <= m <= 10e9
*/

func divisibilityArray(word string, m int) []int {
	return []int{}
}

/*
给你一个下标从 0 开始的整数数组 nums 。
一开始，所有下标都没有被标记。你可以执行以下操作任意次：
选择两个 互不相同且未标记 的下标 i 和 j ，满足 2 * nums[i] <= nums[j] ，标记下标 i 和 j 。
请你执行上述操作任意次，返回 nums 中最多可以标记的下标数目。
示例 1：
输入：nums = [3,5,2,4]
输出：2
解释：第一次操作中，选择 i = 2 和 j = 1 ，操作可以执行的原因是 2 * nums[2] <= nums[1] ，标记下标 2 和 1 。
没有其他更多可执行的操作，所以答案为 2 。
示例 2：
输入：nums = [9,2,5,4]
输出：4
解释：第一次操作中，选择 i = 3 和 j = 0 ，操作可以执行的原因是 2 * nums[3] <= nums[0] ，标记下标 3 和 0 。
第二次操作中，选择 i = 1 和 j = 2 ，操作可以执行的原因是 2 * nums[1] <= nums[2] ，标记下标 1 和 2 。
没有其他更多可执行的操作，所以答案为 4 。
示例 3：
输入：nums = [7,6,8]
输出：0
解释：没有任何可以执行的操作，所以答案为 0 。
提示：
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e9
*/

func maxNumOfMarkedIndices(nums []int) int {
	return 0
}

/*
给你一个 m x n 的矩阵 grid ，每个元素都为 非负 整数，其中 grid[row][col] 表示可以访问格子 (row, col) 的 最早 时间。也就是说当你访问格子 (row, col) 时，最少已经经过的时间为 grid[row][col] 。
你从 最左上角 出发，出发时刻为 0 ，你必须一直移动到上下左右相邻四个格子中的 任意 一个格子（即不能停留在格子上）。每次移动都需要花费 1 单位时间。
请你返回 最早 到达右下角格子的时间，如果你无法到达右下角的格子，请你返回 -1 。
示例 1：
输入：grid = [[0,1,3,2],[5,1,2,5],[4,3,8,6]]
输出：7
解释：一条可行的路径为：
- 时刻 t = 0 ，我们在格子 (0,0) 。
- 时刻 t = 1 ，我们移动到格子 (0,1) ，可以移动的原因是 grid[0][1] <= 1 。
- 时刻 t = 2 ，我们移动到格子 (1,1) ，可以移动的原因是 grid[1][1] <= 2 。
- 时刻 t = 3 ，我们移动到格子 (1,2) ，可以移动的原因是 grid[1][2] <= 3 。
- 时刻 t = 4 ，我们移动到格子 (1,1) ，可以移动的原因是 grid[1][1] <= 4 。
- 时刻 t = 5 ，我们移动到格子 (1,2) ，可以移动的原因是 grid[1][2] <= 5 。
- 时刻 t = 6 ，我们移动到格子 (1,3) ，可以移动的原因是 grid[1][3] <= 6 。
- 时刻 t = 7 ，我们移动到格子 (2,3) ，可以移动的原因是 grid[2][3] <= 7 。
最终到达时刻为 7 。这是最早可以到达的时间。
示例 2：
输入：grid = [[0,2,4],[3,2,1],[1,0,4]]
输出：-1
解释：没法从左上角按题目规定走到右下角。
提示：
m == grid.length
n == grid[i].length
2 <= m, n <= 1000
4 <= m * n <= 10e5
0 <= grid[i][j] <= 10e5
grid[0][0] == 0
*/

func minimumTime(grid [][]int) int {
	return 0
}

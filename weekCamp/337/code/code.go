package code

/*
给你一个 正 整数 n 。
用 even 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的偶数下标的个数。
用 odd 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的奇数下标的个数。
返回整数数组 answer ，其中 answer = [even, odd] 。
示例 1：
输入：n = 17
输出：[2,0]
解释：17 的二进制形式是 10001 。
下标 0 和 下标 4 对应的值为 1 。
共有 2 个偶数下标，0 个奇数下标。
示例 2：
输入：n = 2
输出：[0,1]
解释：2 的二进制形式是 10 。
下标 1 对应的值为 1 。
共有 0 个偶数下标，1 个奇数下标。
提示：
1 <= n <= 1000
*/

func evenOddBit(n int) []int {
	ans := make([]int, 2)
	for i := 0; n > 0; i ^= 1 {
		ans[i] += n & 1
		n >>= 1
	}
	return ans
}

/*
骑士在一张 n x n 的棋盘上巡视。在有效的巡视方案中，骑士会从棋盘的 左上角 出发，并且访问棋盘上的每个格子 恰好一次 。
给你一个 n x n 的整数矩阵 grid ，由范围 [0, n * n - 1] 内的不同整数组成，其中 grid[row][col] 表示单元格 (row, col) 是骑士访问的第 grid[row][col] 个单元格。骑士的行动是从下标 0 开始的。
如果 grid 表示了骑士的有效巡视方案，返回 true；否则返回 false。
注意，骑士行动时可以垂直移动两个格子且水平移动一个格子，或水平移动两个格子且垂直移动一个格子。下图展示了骑士从某个格子出发可能的八种行动路线。
示例 1：
输入：grid = [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
输出：true
解释：grid 如上图所示，可以证明这是一个有效的巡视方案。
示例 2：
输入：grid = [[0,3,6],[5,8,1],[2,7,4]]
输出：false
解释：grid 如上图所示，考虑到骑士第 7 次行动后的位置，第 8 次行动是无效的。
提示：
n == grid.length == grid[i].length
3 <= n <= 7
0 <= grid[row][col] < n * n
grid 中的所有整数 互不相同
*/

func checkValidGrid(grid [][]int) bool {
	//左上出发规则
	if grid[0][0] != 0 {
		return false
	}
	//1. 带下标排序
	type point struct {
		rId, cId int
	}
	points := make([]point, len(grid[0])*len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid[0]); j++ {
			points[grid[i][j]] = point{i, j}
		}
	}

	//2. legal规则(剪枝？) 顺序判断规则是否成立
	for i := 0; i < len(points)-1; i++ {
		if !(abs(points[i].cId, points[i+1].cId) == 1 && abs(points[i].rId, points[i+1].rId) == 2 ||
			abs(points[i].cId, points[i+1].cId) == 2 && abs(points[i].rId, points[i+1].rId) == 1) {
			return false
		}
	}
	return true
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

/*
给你一个由正整数组成的数组 nums 和一个 正整数 k 。
如果 nums 的子集中，任意两个整数的绝对差均不等于 k ，则认为该子数组是一个 美丽 子集。
返回数组 nums 中 非空 且 美丽 的子集数目。
nums 的子集定义为：可以经由 nums 删除某些元素（也可能不删除）得到的一个数组。只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的子集。
示例 1：
输入：nums = [2,4,6], k = 2
输出：4
解释：数组 nums 中的美丽子集有：[2], [4], [6], [2, 6] 。
可以证明数组 [2,4,6] 中只存在 4 个美丽子集。
示例 2：
输入：nums = [1], k = 1
输出：1
解释：数组 nums 中的美丽数组有：[1] 。
可以证明数组 [1] 中只存在 1 个美丽子集。
提示：
1 <= nums.length <= 20
1 <= nums[i], k <= 1000
*/

func beautifulSubsets(nums []int, k int) int {
	return 0
}

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 value 。
在一步操作中，你可以对 nums 中的任一元素加上或减去 value 。
例如，如果 nums = [1,2,3] 且 value = 2 ，你可以选择 nums[0] 减去 value ，得到 nums = [-1,2,3] 。
数组的 MEX (minimum excluded) 是指其中数组中缺失的最小非负整数。
例如，[-1,2,3] 的 MEX 是 0 ，而 [1,0,3] 的 MEX 是 2 。
返回在执行上述操作 任意次 后，nums 的最大 MEX 。
示例 1：
输入：nums = [1,-10,7,13,6,8], value = 5
输出：4
解释：执行下述操作可以得到这一结果：
- nums[1] 加上 value 两次，nums = [1,0,7,13,6,8]
- nums[2] 减去 value 一次，nums = [1,0,2,13,6,8]
- nums[3] 减去 value 两次，nums = [1,0,2,3,6,8]
nums 的 MEX 是 4 。可以证明 4 是可以取到的最大 MEX 。
示例 2：
输入：nums = [1,-10,7,13,6,8], value = 7
输出：2
解释：执行下述操作可以得到这一结果：
- nums[2] 减去 value 一次，nums = [1,-10,0,13,6,8]
nums 的 MEX 是 2 。可以证明 2 是可以取到的最大 MEX 。
提示：
1 <= nums.length, value <= 10e5
-10e9 <= nums[i] <= 10e9
*/

func findSmallestInteger(nums []int, value int) int {
	return 0
}

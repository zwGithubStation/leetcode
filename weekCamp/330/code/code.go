package code

/*
给你一个正整数 n ，开始时，它放在桌面上。在 109 天内，每天都要执行下述步骤：
对于出现在桌面上的每个数字 x ，找出符合 1 <= i <= n 且满足 x % i == 1 的所有数字 i 。
然后，将这些数字放在桌面上。
返回在 109 天之后，出现在桌面上的 不同 整数的数目。
注意：
一旦数字放在桌面上，则会一直保留直到结束。
% 表示取余运算。例如，14 % 3 等于 2 。
示例 1：
输入：n = 5
输出：4
解释：最开始，5 在桌面上。
第二天，2 和 4 也出现在桌面上，因为 5 % 2 == 1 且 5 % 4 == 1 。
再过一天 3 也出现在桌面上，因为 4 % 3 == 1 。
在十亿天结束时，桌面上的不同数字有 2 、3 、4 、5 。
示例 2：
输入：n = 3
输出：2
解释：
因为 3 % 2 == 1 ，2 也出现在桌面上。
在十亿天结束时，桌面上的不同数字只有两个：2 和 3 。
提示：
1 <= n <= 100
*/
func distinctIntegers(n int) int {

}

/*
现在有一个正凸多边形，其上共有 n 个顶点。顶点按顺时针方向从 0 到 n - 1 依次编号。每个顶点上 正好有一只猴子 。下图中是一个 6 个顶点的凸多边形。
每个猴子同时移动到相邻的顶点。顶点 i 的相邻顶点可以是：
顺时针方向的顶点 (i + 1) % n ，或
逆时针方向的顶点 (i - 1 + n) % n 。
如果移动后至少有两个猴子位于同一顶点，则会发生 碰撞 。
返回猴子至少发生 一次碰撞 的移动方法数。由于答案可能非常大，请返回对 109+7 取余后的结果。
注意，每只猴子只能移动一次。
*/
func monkeyMove(n int) int {

}

func putMarbles(weights []int, k int) int64 {

}

func countQuadruplets(nums []int) int64 {

}

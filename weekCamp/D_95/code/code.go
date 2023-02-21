package code

import (
	"math"
	"sort"
)

/*
给你四个整数 length ，width ，height 和 mass ，分别表示一个箱子的三个维度和质量，请你返回一个表示箱子 类别 的字符串。
如果满足以下条件，那么箱子是 "Bulky" 的：
箱子 至少有一个 维度大于等于 10e4 。
或者箱子的 体积 大于等于 10e9 。
如果箱子的质量大于等于 100 ，那么箱子是 "Heavy" 的。
如果箱子同时是 "Bulky" 和 "Heavy" ，那么返回类别为 "Both" 。
如果箱子既不是 "Bulky" ，也不是 "Heavy" ，那么返回类别为 "Neither" 。
如果箱子是 "Bulky" 但不是 "Heavy" ，那么返回类别为 "Bulky" 。
如果箱子是 "Heavy" 但不是 "Bulky" ，那么返回类别为 "Heavy" 。
注意，箱子的体积等于箱子的长度、宽度和高度的乘积。
提示：
1 <= length, width, height <= 10e5
1 <= mass <= 10e3
*/
func categorizeBox(length, width, height, mass int) string {
	x := length >= 10e4 || width >= 10e4 || height >= 10e4 || length*width*height >= 10e9
	y := mass >= 100
	switch {
	case x && y:
		return "Both"
	case x:
		return "Bulky"
	case y:
		return "Heavy"
	default:
		return "Neither"
	}
}

/*
给你一个整数数据流，请你实现一个数据结构，检查数据流中最后 k 个整数是否 等于 给定值 value 。
请你实现 DataStream 类：
DataStream(int value, int k) 用两个整数 value 和 k 初始化一个空的整数数据流。
boolean consec(int num) 将 num 添加到整数数据流。如果后 k 个整数都等于 value ，返回 true ，否则返回 false 。如果少于 k 个整数，条件不满足，所以也返回 false 。
示例 1：
输入：
["DataStream", "consec", "consec", "consec", "consec"]
[[4, 3], [4], [4], [4], [3]]
输出：
[null, false, false, true, false]
解释：
DataStream dataStream = new DataStream(4, 3); // value = 4, k = 3
dataStream.consec(4); // 数据流中只有 1 个整数，所以返回 False 。
dataStream.consec(4); // 数据流中只有 2 个整数
                      // 由于 2 小于 k ，返回 False 。
dataStream.consec(4); // 数据流最后 3 个整数都等于 value， 所以返回 True 。
dataStream.consec(3); // 最后 k 个整数分别是 [4,4,3] 。
                      // 由于 3 不等于 value ，返回 False 。
提示：
1 <= value, num <= 10e9
1 <= k <= 10e5
至多调用 consec 次数为 10e5 次。
*/

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */

type DataStream struct {
	value, k, cnt int
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k, cnt: 0}
}

func (this *DataStream) Consec(num int) bool {
	if this.value == num {
		this.cnt++
	} else {
		this.cnt = 0
	}
	return this.k <= this.cnt
}

/*
给你一个下标从 0 开始的整数数组 nums 。
三个下标 i ，j 和 k 的 有效值 定义为 ((nums[i] | nums[j]) & nums[k]) 。
一个数组的 xor 美丽值 是数组中所有满足 0 <= i, j, k < n  的三元组 (i, j, k) 的 有效值 的异或结果。
请你返回 nums 的 xor 美丽值。
注意：
val1 | val2 是 val1 和 val2 的按位或。
val1 & val2 是 val1 和 val2 的按位与。
示例 1：
输入：nums = [1,4]
输出：5
解释：
三元组和它们对应的有效值如下：
- (0,0,0) 有效值为 ((1 | 1) & 1) = 1
- (0,0,1) 有效值为 ((1 | 1) & 4) = 0
- (0,1,0) 有效值为 ((1 | 4) & 1) = 1
- (0,1,1) 有效值为 ((1 | 4) & 4) = 4
- (1,0,0) 有效值为 ((4 | 1) & 1) = 1
- (1,0,1) 有效值为 ((4 | 1) & 4) = 4
- (1,1,0) 有效值为 ((4 | 4) & 1) = 0
- (1,1,1) 有效值为 ((4 | 4) & 4) = 4
数组的 xor 美丽值为所有有效值的按位异或 1 ^ 0 ^ 1 ^ 4 ^ 1 ^ 4 ^ 0 ^ 4 = 5 。
示例 2：
输入：nums = [15,45,20,2,34,35,5,44,32,30]
输出：34
解释：数组的 xor 美丽值为 34 。
提示：
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e9
*/
func xorBeauty(nums []int) (ans int) {
	//pass for now
	return 1
}

/*
给你一个下标从 0 开始长度为 n 的整数数组 stations ，其中 stations[i] 表示第 i 座城市的供电站数目。
每个供电站可以在一定 范围 内给所有城市提供电力。换句话说，如果给定的范围是 r ，在城市 i 处的供电站可以给所有满足 |i - j| <= r 且 0 <= i, j <= n - 1 的城市 j 供电。
|x| 表示 x 的 绝对值 。比方说，|7 - 5| = 2 ，|3 - 10| = 7 。
一座城市的 电量 是所有能给它供电的供电站数目。
政府批准了可以额外建造 k 座供电站，你需要决定这些供电站分别应该建在哪里，这些供电站与已经存在的供电站有相同的供电范围。
给你两个整数 r 和 k ，如果以最优策略建造额外的发电站，返回所有城市中，最小供电站数目的最大值是多少。
这 k 座供电站可以建在多个城市。

示例 1：
输入：stations = [1,2,4,5,0], r = 1, k = 2
输出：5
解释：
最优方案之一是把 2 座供电站都建在城市 1 。
每座城市的供电站数目分别为 [1,4,4,5,0] 。
- 城市 0 的供电站数目为 1 + 4 = 5 。
- 城市 1 的供电站数目为 1 + 4 + 4 = 9 。
- 城市 2 的供电站数目为 4 + 4 + 5 = 13 。
- 城市 3 的供电站数目为 5 + 4 = 9 。
- 城市 4 的供电站数目为 5 + 0 = 5 。
供电站数目最少是 5 。
无法得到更优解，所以我们返回 5 。
示例 2：
输入：stations = [4,4,4,4], r = 0, k = 3
输出：4
解释：
无论如何安排，总有一座城市的供电站数目是 4 ，所以最优解是 4 。
提示：
n == stations.length
1 <= n <= 10e5
0 <= stations[i] <= 10e5
0 <= r <= n - 1
0 <= k <= 10e9
*/
/*
介绍1. 前缀和
一个array N：
N = [N1, N2, N3 .... NK, NK+1, Nk+2 .... NG] (共G个元素)
可以有其前缀和array：
M = [M1, M2, M3 .... MK, MK+1, Mk+2 .... MG, MG+1] (共G+1个元素)
其中：
M1 = 0
Mi = Mi-1 + Ni-1(i > 0) 更浅显的含义是，Mi代表N中前i-1个元素的和

前缀和的作用：
计算N的任意连续字串的和，可以通过M得到：
Ni + Ni+1 + Ni+2 + ..... + Nj = Mj+1 - Mi


介绍2. 差分数组
一个array N：
N = [N1, N2, N3 .... NK, NK+1, Nk+2 .... NG] (共G个元素)
建立初始差分array：
M = [0, 0, 0, ... 0, 0, 0, 0... 0, 0, 0, 0] (共G个元素，全部为0)

差分数组可以优化下面的问题：
如果每次对N的操作都是：对N的某个连续子串(e.g. {NK, NK+1, NK+2})进行所有元素值+k
那么暴力的方式是每次操作复杂度0(N)[对应的元素都做修改]
而差分数组则可以每次修改M内最多2个元素的值达到记录该修改操作的效果，并最终通过计算前缀和得到N的最终结果：
具体为
某次N的连续子集{Ni, Ni+1 .... Nj}执行了所有元素+k的操作，只需对M做如下记录：
Mi += k
if j < G-1(Nj不是N的最后一个元素)
{ Mj+1 -= k }
else (Nj不是N的最后一个元素)
{// do nothing}

在经过多次的如上操作后，将M更新为其前缀和(定义见上一个介绍)，然后更新至N中的每个对应元素(Ni = Ni + Mi),即得到N在多次操作后的最终结果
*/
func maxPower(stations []int, r int, k int) int64 {
	//1计算前缀和
	n := len(stations)
	pre := make([]int, n+1)
	for i, x := range stations {
		pre[i+1] = pre[i] + x
	}
	//2计算当前电量值, 并记录当前最小值
	minPowerNow := math.MaxInt64

	for i, _ := range stations {
		stations[i] = pre[min(i+r+1, n)] - pre[max(0, i-r)]
		minPowerNow = min(minPowerNow, stations[i])
	}

	return int64(minPowerNow + sort.Search(k, func(cur_min int) bool {
		cur_min += minPowerNow + 1
		diff := make([]int, n)
		cur_pre, need := 0, 0
		for i, x := range stations {
			cur_pre += diff[i] //计算当前前缀和的过程 得到当前市的变更电站值 为正说明之前的变更使得这里的电站得到了增强
			m := cur_min - x - cur_pre
			if m > 0 {
				need += m //累积式记录
				if need > k {
					return true
				}
				//新增的m个电厂 根据最优化的考量 建在距离该节点r的节点上，相当于在 [i, i+2r]的范围内触发了全部+m的操作
				//因此需要更新差分数组，两个元素分别是i和i+2r+1
				//i的变化直接用于最新的前缀和计算
				cur_pre += m
				//i+2r+1的差分更新确保不溢出
				if i+2*r+1 < n {
					diff[i+2*r+1] -= m
				}
			}
		}
		return false
	}))
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

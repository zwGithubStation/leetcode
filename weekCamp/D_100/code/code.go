package code

import (
	"math"
	"sort"
)

/*
给你一个整数 money ，表示你总共有的钱数（单位为美元）和另一个整数 children ，表示你要将钱分配给多少个儿童。
你需要按照如下规则分配：
所有的钱都必须被分配。
每个儿童至少获得 1 美元。
没有人获得 4 美元。
请你按照上述规则分配金钱，并返回 最多 有多少个儿童获得 恰好 8 美元。如果没有任何分配方案，返回 -1 。
示例 1：
输入：money = 20, children = 3
输出：1
解释：
最多获得 8 美元的儿童数为 1 。一种分配方案为：
- 给第一个儿童分配 8 美元。
- 给第二个儿童分配 9 美元。
- 给第三个儿童分配 3 美元。
没有分配方案能让获得 8 美元的儿童数超过 1 。
示例 2：
输入：money = 16, children = 2
输出：2
解释：每个儿童都可以获得 8 美元。
提示：
1 <= money <= 200
2 <= children <= 30
*/

func distMoney(money int, children int) int {
	var count int
	if money < children { //无法保证每人一块钱
		return -1
	}

	if money < 8 { //无法保证有人能得到8块钱
		return 0
	}

	//在每个人至少有1块钱的基础上 计算得到8块的人数最大值
	for children > 0 && money >= children+7 {
		children--
		money = money - 8
		count++
	}

	//基于剩余的人数/钱数  避免出现有人分到4块 || 钱没有分完的 情况出现
	if (children == 1 && money == 4) || (children == 0 && money != 0) {
		return count - 1
	}

	return count
}

/*
给你一个下标从 0 开始的整数数组 nums 。你需要将 nums 重新排列成一个新的数组 perm 。
定义 nums 的 伟大值 为满足 0 <= i < nums.length 且 perm[i] > nums[i] 的下标数目。
请你返回重新排列 nums 后的 最大 伟大值。
示例 1：
输入：nums = [1,3,5,2,1,3,1]
输出：4
解释：一个最优安排方案为 perm = [2,5,1,3,3,1,1] 。
在下标为 0, 1, 3 和 4 处，都有 perm[i] > nums[i] 。因此我们返回 4 。
示例 2：
输入：nums = [1,2,3,4]
输出：3
解释：最优排列为 [2,3,4,1] 。
在下标为 0, 1 和 2 处，都有 perm[i] > nums[i] 。因此我们返回 3 。
提示：
1 <= nums.length <= 10e5
0 <= nums[i] <= 10e9
*/

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)
	i, j, lens, ret := 0, 1, len(nums), 0
	for i < lens && j < lens {
		if nums[i] < nums[j] {
			i++
			j++
			ret++
		} else {
			j++
		}
	}
	return ret
}

/*
给你一个数组 nums ，它包含若干正整数。
一开始分数 score = 0 ，请你按照下面算法求出最后分数：
从数组中选择最小且没有被标记的整数。如果有相等元素，选择下标最小的一个。
将选中的整数加到 score 中。
标记 被选中元素，如果有相邻元素，则同时标记 与它相邻的两个元素 。
重复此过程直到数组中所有元素都被标记。
请你返回执行上述算法后最后的分数。
示例 1：
输入：nums = [2,1,3,4,5,2]
输出：7
解释：我们按照如下步骤标记元素：
- 1 是最小未标记元素，所以标记它和相邻两个元素：[2,1,3,4,5,2] 。
- 2 是最小未标记元素，所以标记它和左边相邻元素：[2,1,3,4,5,2] 。
- 4 是仅剩唯一未标记的元素，所以我们标记它：[2,1,3,4,5,2] 。
总得分为 1 + 2 + 4 = 7 。
示例 2：
输入：nums = [2,3,5,1,3,2]
输出：5
解释：我们按照如下步骤标记元素：
- 1 是最小未标记元素，所以标记它和相邻两个元素：[2,3,5,1,3,2] 。
- 2 是最小未标记元素，由于有两个 2 ，我们选择最左边的一个 2 ，也就是下标为 0 处的 2 ，以及它右边相邻的元素：[2,3,5,1,3,2] 。
- 2 是仅剩唯一未标记的元素，所以我们标记它：[2,3,5,1,3,2] 。
总得分为 1 + 2 + 2 = 5 。
提示：
1 <= nums.length <= 10e5
1 <= nums[i] <= 10e6
*/

func findScore(nums []int) (ans int64) {
	type pair struct {
		x, i int
	}
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, i}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].i < a[j].i)
	})

	target := make([]bool, len(nums))

	for _, x := range a {
		if target[x.i] == false {
			target[x.i] = true
			if x.i > 0 {
				target[x.i-1] = true
			}
			if x.i < len(nums)-1 {
				target[x.i+1] = true
			}
			ans += int64(x.x)
		}
	}
	return ans
}

/*
给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * ne2 分钟内修好 n 辆车。
同时给你一个整数 cars ，表示总共需要修理的汽车数目。
请你返回修理所有汽车 最少 需要多少时间。
注意：所有机械工可以同时修理汽车。
示例 1：
输入：ranks = [4,2,3,1], cars = 10
输出：16
解释：
- 第一位机械工修 2 辆车，需要 4 * 2 * 2 = 16 分钟。
- 第二位机械工修 2 辆车，需要 2 * 2 * 2 = 8 分钟。
- 第三位机械工修 2 辆车，需要 3 * 2 * 2 = 12 分钟。
- 第四位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
16 分钟是修理完所有车需要的最少时间。
示例 2：
输入：ranks = [5,1,8], cars = 6
输出：16
解释：
- 第一位机械工修 1 辆车，需要 5 * 1 * 1 = 5 分钟。
- 第二位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
- 第三位机械工修 1 辆车，需要 8 * 1 * 1 = 8 分钟。
16 分钟时修理完所有车需要的最少时间。
提示：
1 <= ranks.length <= 10e5
1 <= ranks[i] <= 100
1 <= cars <= 10e6
*/

func repairCars(ranks []int, cars int) int64 {
	minRank := 100
	for _, x := range ranks {
		if x < minRank {
			minRank = x
		}
	}
	return int64(sort.Search(minRank*cars*cars, func(t int) bool {
		fixCount := 0
		for _, x := range ranks {
			fixCount += int(math.Sqrt(float64(t / x)))
		}
		return fixCount >= cars
	}))
}

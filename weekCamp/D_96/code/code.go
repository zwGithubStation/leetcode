package code

import (
	"container/heap"
	"sort"
)

/*
给你两个整数数组 nums1 和 nums2 ，它们已经按非降序排序，请你返回两个数组的 最小公共整数 。如果两个数组 nums1 和 nums2 没有公共整数，请你返回 -1 。
如果一个整数在两个数组中都 至少出现一次 ，那么这个整数是数组 nums1 和 nums2 公共 的。
示例 1：
输入：nums1 = [1,2,3], nums2 = [2,4]
输出：2
解释：两个数组的最小公共元素是 2 ，所以我们返回 2 。
示例 2：
输入：nums1 = [1,2,3,6], nums2 = [2,3,4,5]
输出：2
解释：两个数组中的公共元素是 2 和 3 ，2 是较小值，所以返回 2 。
*/
func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return -1
}

/*
给你两个整数数组 nums1 和 nums2 ，两个数组长度都是 n ，再给你一个整数 k 。你可以对数组 nums1 进行以下操作：
选择两个下标 i 和 j ，将 nums1[i] 增加 k ，将 nums1[j] 减少 k 。换言之，nums1[i] = nums1[i] + k 且 nums1[j] = nums1[j] - k 。
如果对于所有满足 0 <= i < n 都有 num1[i] == nums2[i] ，那么我们称 nums1 等于 nums2 。
请你返回使 nums1 等于 nums2 的 最少 操作数。如果没办法让它们相等，请你返回 -1 。
示例 1：
输入：nums1 = [4,3,1,4], nums2 = [1,3,7,1], k = 3
输出：2
解释：我们可以通过 2 个操作将 nums1 变成 nums2 。
第 1 个操作：i = 2 ，j = 0 。操作后得到 nums1 = [1,3,4,4] 。
第 2 个操作：i = 2 ，j = 3 。操作后得到 nums1 = [1,3,7,1] 。
无法用更少操作使两个数组相等。
示例 2：
输入：nums1 = [3,8,5,2], nums2 = [2,4,1,6], k = 1
输出：-1
解释：无法使两个数组相等。
提示：
n == nums1.length == nums2.length
2 <= n <= 10^5
0 <= nums1[i], nums2[j] <= 10^9
0 <= k <= 10^5
*/

func minOperations(nums1, nums2 []int, k int) (ans int64) {
	sum := 0
	for i, x := range nums1 {
		x -= nums2[i]
		if k > 0 {
			if x%k != 0 {
				return -1
			}
			if x > 0 {
				sum += x / k
			}
		}
	}
}

/*
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，两者长度都是 n ，再给你一个正整数 k 。你必须从 nums1 中选一个长度为 k 的 子序列 对应的下标。
对于选择的下标 i0 ，i1 ，...， ik - 1 ，你的 分数 定义如下：
nums1 中下标对应元素求和，乘以 nums2 中下标对应元素的 最小值 。
用公示表示： (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]) 。
请你返回 最大 可能的分数。
一个数组的 子序列 下标是集合 {0, 1, ..., n-1} 中删除若干元素得到的剩余集合，也可以不删除任何元素。
示例 1：
输入：nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
输出：12
解释：
四个可能的子序列分数为：
- 选择下标 0 ，1 和 2 ，得到分数 (1+3+3) * min(2,1,3) = 7 。
- 选择下标 0 ，1 和 3 ，得到分数 (1+3+2) * min(2,1,4) = 6 。
- 选择下标 0 ，2 和 3 ，得到分数 (1+3+2) * min(2,3,4) = 12 。
- 选择下标 1 ，2 和 3 ，得到分数 (3+3+2) * min(1,3,4) = 8 。
所以最大分数为 12 。
示例 2：
输入：nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
输出：30
解释：
选择下标 2 最优：nums1[2] * nums2[2] = 3 * 10 = 30 是最大可能分数。
提示：
n == nums1.length == nums2.length
1 <= n <= 105
0 <= nums1[i], nums2[j] <= 105
1 <= k <= n
*/
func maxScore(nums1, nums2 []int, k int) int64 {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].y > a[j].y })

	h := hp{nums2[:k]} // 复用内存
	for i, p := range a[:k] {
		sum += p.x
		h.IntSlice[i] = p.x
	}
	ans := sum * a[k-1].y
	heap.Init(&h)
	for _, p := range a[k:] {
		if p.x > h.IntSlice[0] {
			sum += p.x - h.replace(p.x)
			ans = max(ans, sum*p.y)
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}
func (h hp) replace(v int) int  { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/*
给你一个无穷大的网格图。一开始你在 (1, 1) ，你需要通过有限步移动到达点 (targetX, targetY) 。
每一步 ，你可以从点 (x, y) 移动到以下点之一：
(x, y - x)
(x - y, y)
(2 * x, y)
(x, 2 * y)
给你两个整数 targetX 和 targetY ，分别表示你最后需要到达点的 X 和 Y 坐标。如果你可以从 (1, 1) 出发到达这个点，请你返回true ，否则返回 false 。
示例 1：
输入：targetX = 6, targetY = 9
输出：false
解释：没法从 (1,1) 出发到达 (6,9) ，所以返回 false 。
示例 2：
输入：targetX = 4, targetY = 7
输出：true
解释：你可以按照以下路径到达：(1,1) -> (1,2) -> (1,4) -> (1,8) -> (1,7) -> (2,7) -> (4,7) 。
提示：
1 <= targetX, targetY <= 109
*/
func isReachable(targetX, targetY int) bool {
	g := gcd(targetX, targetY)
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

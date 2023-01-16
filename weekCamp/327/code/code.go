package code

import (
	"container/heap"
	"sort"
)

func maximumCount(nums []int) int {
	negCnt, posCnt := 0, 0
	for _, x := range nums {
		if x < 0 {
			negCnt++
		}
		if x > 0 {
			posCnt++
		}
	}
	if posCnt >= negCnt {
		return posCnt
	}
	return negCnt
}

//使用二分查找优化上面函数
//熟练使用下标
func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
func maximumCountBySearch(nums []int) int {
	negCnt := sort.SearchInts(nums, 0)
	posCnt := len(nums) - sort.SearchInts(nums, 1)
	return max(negCnt, posCnt)
}

func maxKelements(nums []int, k int) (ans int64) {
	h := hp{nums}
	heap.Init(&h)
	for ; k > 0; k-- {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

func isItPossible(word1, word2 string) bool {
	c1 := map[rune]int{}
	for _, c := range word1 {
		c1[c]++
	}
	c2 := map[rune]int{}
	for _, c := range word2 {
		c2[c]++
	}
	for x, c := range c1 {
		for y, d := range c2 {
			if y == x { // 无变化
				if len(c1) == len(c2) {
					return true
				}
			} else if len(c1)-b2i(c == 1)+b2i(c1[y] == 0) ==
				len(c2)-b2i(d == 1)+b2i(c2[x] == 0) { // 基于长度计算变化量
				return true
			}
		}
	}
	return false
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func findCrossingTime(n int, k int, time [][]int) int {
	return 1
}

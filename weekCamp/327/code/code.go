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

func isItPossible(word1 string, word2 string) bool {
	return true
}

func findCrossingTime(n int, k int, time [][]int) int {
	return 1
}

package code

import "sort"

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

func maxKelements(nums []int, k int) int64 {
	return 1
}

func isItPossible(word1 string, word2 string) bool {
	return true
}

func findCrossingTime(n int, k int, time [][]int) int {
	return 1
}

type ansd = []int

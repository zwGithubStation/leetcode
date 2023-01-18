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

//结构体中只有一个类型，没有参数名 是hp的anonymous fields
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

func findCrossingTime(n, k int, time [][]int) (cur int) {
	sort.SliceStable(time, func(i, j int) bool {
		a, b := time[i], time[j]
		return a[0]+a[2] < b[0]+b[2]
	})
	waitL, waitR := make(hp1, k), hp1{}
	for i := range waitL {
		waitL[i].i = k - 1 - i // 下标越大效率越低
	}
	workL, workR := hp2{}, hp2{}
	for n > 0 {
		for len(workL) > 0 && workL[0].t <= cur {
			heap.Push(&waitL, heap.Pop(&workL)) // 左边完成放箱
		}
		for len(workR) > 0 && workR[0].t <= cur {
			heap.Push(&waitR, heap.Pop(&workR)) // 右边完成搬箱
		}
		if len(waitR) > 0 { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
			p := heap.Pop(&waitR).(pair)
			cur += time[p.i][2]
			heap.Push(&workL, pair{p.i, cur + time[p.i][3]}) // 放箱，记录完成时间
		} else if len(waitL) > 0 { // 左边过桥
			p := heap.Pop(&waitL).(pair)
			cur += time[p.i][0]
			heap.Push(&workR, pair{p.i, cur + time[p.i][1]}) // 搬箱，记录完成时间
			n--
		} else if len(workL) == 0 { // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
			cur = workR[0].t
		} else if len(workR) == 0 {
			cur = workL[0].t
		} else {
			cur = min(workL[0].t, workR[0].t)
		}
	}
	for len(workR) > 0 {
		p := heap.Pop(&workR).(pair) // 右边完成搬箱
		// 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
		cur = max(p.t, cur) + time[p.i][2]
	}
	return cur // 最后一个过桥的时间
}

type pair struct{ i, t int }
type hp1 []pair

func (h hp1) Len() int            { return len(h) }
func (h hp1) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp1) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp1) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp1) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp2 []pair

func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

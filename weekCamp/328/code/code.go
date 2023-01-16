package code

func differenceOfSum(nums []int) (ans int) {
	for _, x := range nums {
		for ans += x; x > 0; x /= 10 {
			ans -= x % 10
		}
	}
	return
}

func rangeAddQueries(n int, queries [][]int) [][]int {
	m := n

	// 二维差分模板
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, m+1)
	}
	update := func(r1, c1, r2, c2, x int) {
		r2++
		c2++
		diff[r1][c1] += x
		diff[r1][c2] -= x
		diff[r2][c1] -= x
		diff[r2][c2] += x
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原
	ans := make([][]int, n+1)
	ans[0] = make([]int, m+1)
	for i, row := range diff[:n] {
		ans[i+1] = make([]int, m+1)
		for j, x := range row[:m] {
			ans[i+1][j+1] = ans[i+1][j] + ans[i][j+1] - ans[i][j] + x
		}
	}
	ans = ans[1:]
	for i, row := range ans {
		ans[i] = row[1:]
	}
	return ans
}

func countGood(nums []int, k int) (ans int64) {
	cnt := map[int]int{}
	left, pairs := 0, 0
	for _, x := range nums {
		pairs += cnt[x]
		cnt[x]++ // 移入右端点
		for pairs-cnt[nums[left]]+1 >= k {
			cnt[nums[left]]-- // 移出左端点
			pairs -= cnt[nums[left]]
			left++
		}
		if pairs >= k {
			ans += int64(left + 1)
		}
	}
	return
}

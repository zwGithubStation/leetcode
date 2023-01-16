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

func maxOutput(n int, edges [][]int, price []int) int64 {
	ans := 0
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}
	// 返回带叶子的最大路径和，不带叶子的最大路径和
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		p := price[x]
		maxS1, maxS2 := p, 0
		for _, y := range g[x] {
			if y != fa {
				s1, s2 := dfs(y, x)
				// 前面最大带叶子的路径和 + 当前不带叶子的路径和
				// 前面最大不带叶子的路径和 + 当前带叶子的路径和
				ans = max(ans, max(maxS1+s2, maxS2+s1))
				maxS1 = max(maxS1, s1+p)
				maxS2 = max(maxS2, s2+p) // 这里加上 p 是因为 x 必然不是叶子
			}
		}
		return maxS1, maxS2
	}
	dfs(0, -1)
	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

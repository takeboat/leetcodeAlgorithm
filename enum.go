package main

import (
	"fmt"
	"math"
	"slices"
)

func findMaxK(nums []int) int {
	ans := -1
	has := map[int]bool{}
	for _, x := range nums {
		if abs(x) > ans && has[-x] {
			ans = abs(x)
		}
		has[x] = true
	}
	return ans
}

func maxProfit(prices []int) int {
	var ans int
	minPrice := math.MaxInt32
	for _, price := range prices {
		minPrice = min(minPrice, price)
		ans = max(ans, price-minPrice)
	}
	return ans
}

func maximumDifference(nums []int) int {
	ans := -1
	small := math.MaxInt32
	for _, num := range nums {
		small = min(small, num)
		ans = max(ans, num-small)
	}
	return ans
}

func maximumSum(nums []int) int {
	ans := -1
	mx := [82]int{}
	for i := range mx {
		mx[i] = math.MinInt32
	}
	// 当前的值+ 左侧维护的最大值
	for i := 0; i < len(nums); i++ {
		s := 0
		for x := nums[i]; x > 0; x /= 10 {
			s += x % 10
		}
		ans = max(ans, mx[s]+nums[i])
		mx[s] = max(mx[s], nums[i])
	}
	return ans
}

func numEquivDominoPairs(dominoes [][]int) int {
	var ans int
	dMap := make(map[int]int)
	for _, d := range dominoes {
		x, y := min(d[0], d[1]), max(d[0], d[1])
		ans += dMap[x*10+y]
		dMap[x*10+y]++
	}
	return ans
}

func maxOperations(nums []int, k int) int {
	var ans int
	nMap := make(map[int]int)
	for _, x := range nums {
		// 先判断
		if nMap[k-x] > 0 {
			nMap[k-x]--
			ans++
		} else {
			nMap[x]++
		}
	}
	return ans
}

func containsNearbyDuplicate(nums []int, k int) bool {
	nMap := make(map[int]int)
	for i, x := range nums {
		if j, ok := nMap[x]; ok && i-j <= k {
			return true
		}
		nMap[x] = i
	}
	return false
}

func pairSums(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	nMap := make(map[int]int)
	for _, num := range nums {
		if j, ok := nMap[target-num]; ok && j > 0 {
			ans = append(ans, []int{num, target - num})
			nMap[target-num]--
		} else {
			nMap[num]++
		}
	}
	return ans
}

func maxSum(nums []int) int {
	ans := math.MinInt32
	d := make(map[int]int)
	for _, num := range nums {
		mx := maxDigit(num)
		if v, ok := d[mx]; ok {
			ans = max(ans, num+v)
			d[mx] = max(v, num)
		} else {
			d[mx] = num
		}
	}
	if ans == math.MinInt32 {
		return -1
	}
	return ans
}

func maxDigit(num int) int {
	var mx int
	for ; num > 0; num /= 10 {
		mx = max(mx, num%10)
	}
	return mx
}

func countTrapezoids(points [][]int) int {
	const mod = 1e9 + 7
	var ans int
	cnt := make(map[int]int, len(points))
	for _, p := range points {
		cnt[p[1]]++
	}
	s := 0
	for _, c := range cnt {
		// 每个点可以组成 n*(n-1)/2 条线段
		k := c * (c - 1) / 2
		ans += s * k
		s += k
	}
	return ans % mod
}

func countBadPairs(nums []int) int64 {
	n := len(nums)
	ans := n * (n - 1) / 2
	cnt := make(map[int]int)
	for j, num := range nums {
		ans -= cnt[num-j]
		cnt[num-j]++
	}
	return int64(ans)
}

func countPairs(words []string) int64 {
	var ans int64
	cnt := make(map[string]int)
	for _, word := range words {
		t := []byte(word)
		base := t[0]
		for i := range t {
			t[i] = (t[i] - base + 26) % 26
		}
		word = string(t)
		ans += int64(cnt[word])
		cnt[word]++
	}
	return ans
}

func minMirrorPairDistance(nums []int) int {
	ans := math.MaxInt32
	pos := map[int]int{}
	for i, num := range nums {
		if j, ok := pos[num]; ok {
			ans = min(ans, abs(i-j))
		}
		pos[reverse(num)] = i
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// reverse 翻转数字,高位0抹去 120 => 21
func reverse(num int) int {
	// 123
	var res int
	for ; num > 0; num /= 10 {
		res = res*10 + num%10
	}
	return res
}

// distince = j - i
// score = values[i] + values[j] - distince
// score = (values[i] + i) + (values[j] - j)
// let |left = values[i] + i
// 核心 合并同类项
func maxScoreSightseeingPair(values []int) int {
	var ans int
	left := 0
	for j, value := range values {
		ans = max(ans, left+value-j)
		left = max(left, value+j)
	}
	return ans
}

func countNicePairs(nums []int) int {
	const mod = 1e9 + 7
	var ans int
	cnt := map[int]int{}
	for _, num := range nums {
		val := num - reverse(num)
		if j, ok := cnt[val]; ok {
			ans += j
		}
		cnt[val]++
	}
	return ans % mod
}

// pos (j-i > 2) 索引
func maximumProduct(nums []int, m int) int64 {
	ans := math.MinInt
	leftMin := math.MaxInt
	leftMax := math.MinInt
	for i := m - 1; i < len(nums); i++ {
		num := nums[i-m+1]
		leftMin = min(leftMin, num)
		leftMax = max(leftMax, num)
		x := nums[i]
		ans = max(ans, x*leftMin, x*leftMax)
	}
	return int64(ans)
}

// abs(i-j) >= indexDifference
// abs(nums[i]-nums[j]) >= valueDifference
// 和常规最大的不同在于 维护的左侧不是 curIndex 而是 curIndex-indexDifference
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	if valueDifference == 0 && indexDifference == 0 {
		return []int{0, 0}
	}
	// 这里应该存储index
	leftMin, leftMax := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		num := nums[j-indexDifference]
		if num > nums[leftMax] {
			leftMax = j - indexDifference
		}
		if num < nums[leftMin] {
			leftMin = j - indexDifference
		}
		if nums[j]-nums[leftMin] >= valueDifference {
			return []int{leftMin, j}
		}
		if nums[leftMax]-nums[j] >= valueDifference {
			return []int{leftMax, j}
		}
	}
	return []int{-1, -1}
}

func countBeautifulPairs(nums []int) int {
	var ans int
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 {
			x /= 10
		}
		cnt[x]++
	}
	return ans
}

func gcd(a, b int) int {
	// 辗转相除法
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func similarPairs(words []string) int {
	var ans int
	cnt := map[int]int{}
	for _, word := range words {
		// 计算掩码
		var mask int
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if cnt[mask] > 0 {
			ans += cnt[mask]
		}
		cnt[mask]++
	}
	return ans
}

// 统计余数
func canArrange(arr []int, k int) bool {
	cnt := make([]int, k)
	for _, x := range arr {
		r := (x%k + k) % k
		cnt[r]++
	}
	// 0 应该是成对的
	if cnt[0]%2 != 0 {
		return false
	}
	for i := 1; i <= k/2; i++ {
		cmp := k - i  // 补数
		if i == cmp { // 如果k是奇数的话需要看当下这个是不是偶数
			if cnt[i]%2 != 0 {
				return false
			}
		} else {
			if cnt[i] != cnt[cmp] {
				return false
			}
		}
	}
	return true
}

func specialTriplets(nums []int) int {
	const mod = 1e9 + 7
	var ans int
	// 统计`k` x频率
	suf := map[int]int{}
	for _, x := range nums {
		suf[x]++
	}
	pre := map[int]int{}
	for _, x := range nums {
		suf[x]--
		ans += pre[x*2] * suf[x*2]
		pre[x]++
	}
	return ans % mod
}

func countPalindromicSubsequence(s string) int {
	n := len(s)
	var ans int
	suf := [26]int{}
	for i := 1; i < n; i++ {
		suf[s[i]-'a']++
	}
	pre := [26]int{}
	vis := [26][26]bool{} // [mid][side] bool
	for i := 1; i < n-1; i++ {
		mid := s[i] - 'a'
		suf[mid]--
		pre[s[i-1]-'a']++
		for j := 0; j < 26; j++ {
			if suf[j] > 0 && pre[j] > 0 && !vis[mid][j] {
				vis[mid][j] = true
				ans++
			}
		}
	}
	return ans
}

func countPalindromicSubsequence1(s string) int {
	n := len(s)
	var ans int
	suf := [26]int{}
	for i := 1; i < n; i++ {
		suf[s[i]-'a']++
	}
	pre := [26]int{}
	vis := [26][26]bool{} // [mid][side] bool
	for i := 1; i < n-1; i++ {
		mid := s[i] - 'a'
		suf[mid]--
		pre[s[i-1]-'a']++
		for j := 0; j < 26; j++ {
			if suf[j] > 0 && pre[j] > 0 && !vis[mid][j] {
				vis[mid][j] = true
				ans++
			}
		}
	}
	return ans
}

// numberOfRightTriangles 运用乘法原理
// 乘法原理:
// 如果完成一件事情需要拆分为n个步骤,做第一步有M1种方法,做第二步有M2种方法,
// 做第n步有Mn种方法,并且每个步骤缺一不可,那么玩车给你这件事情的方法数就是M1*M2*...*Mn
func numberOfRightTriangles(grid [][]int) int64 {
	var ans int64
	n := len(grid[0]) // 列数
	colSum := make([]int, n)
	for _, row := range grid {
		for i, x := range row {
			colSum[i] += x
		}
	}
	for _, row := range grid {
		var rowSum int
		for _, x := range row {
			rowSum += x
		}
		for j, x := range row {
			if x == 1 {
				ans += int64((rowSum - 1) * (colSum[j] - 1))
			}
		}
	}
	return ans
}

func numberOfBoomerangs(points [][]int) int {
	var ans int
	cnt := map[int]int{}
	// 枚举每个点作为回旋镖的中心点i
	for _, p1 := range points {
		// 清空哈希表,只保持当前中心点i的距离统计
		clear(cnt)
		for _, p2 := range points {
			// 计算点p1与p2之间欧式距离的平方(避免浮点数)
			d := (p2[0]-p1[0])*(p2[0]-p1[0]) + (p2[1]-p1[1])*(p2[1]-p1[1])
			// 对于每一个新的p2,假设之前已有m=cnt[d]个点与p1的距离也为d
			// 那么新增的回旋镖的数量为 A(m+1,2) - A(m,2) = m*2
			// 也就是说, 新的p2可以与之前的m个点构成2*m个有序三元组 (i, 旧点, p2), (i, p2, 旧点)
			ans += cnt[d] * 2
			// 将当前p2 纳入距离为d的点集合, m 增加1
			cnt[d]++
		}
	}
	return ans
}

func find132pattern(nums []int) bool {
	// i < j < k
	// nums[i] < nums[k]
	// nums[j] > nums[i]
	// 条件转换为 nums[j] > nums[i] && nums[i] < (nums[j]>nums[k]中的最大值)
	n := len(nums)
	if n < 3 {
		return false
	}
	stack := []int{}
	candidates := make([]int, n)
	for i := range candidates {
		candidates[i] = math.MinInt
	}
	stack = append(stack, nums[n-1])
	for j := n - 2; j >= 1; j-- {
		for len(stack) > 0 && nums[j] > stack[len(stack)-1] {
			candidates[j] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[j])
	}
	leftMin := nums[0]
	for j := 1; j <= n-2; j++ {
		if nums[j] > leftMin && leftMin < candidates[j] {
			return true
		}
		leftMin = min(leftMin, nums[j])
	}
	return false
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	ans := make([]int, n)
	type edge struct{ to, w int }
	g := make([][]edge, n) // 将数组转换为邻接表形式的
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}
	for i, gi := range g {
		// 叶子节点跳过
		if len(gi) == 1 {
			continue
		}
		var cnt int
		var dfs func(int, int, int)
		// 不可以回头
		dfs = func(cur, father, sum int) {
			if sum%signalSpeed == 0 {
				cnt++
			}
			for _, e := range g[cur] {
				if e.to != father {
					dfs(e.to, cur, sum+e.w)
				}
			}
		}
		sum := 0
		// 对于当前中心点 i，我们依次看它的每一条邻边（每一个分支），
		// 用 DFS 算出这个分支里有多少个节点到 i 的路径和能被 signalSpeed 整除（cnt）
		for _, e := range gi {
			cnt = 0
			dfs(e.to, i, e.w)
			ans[i] += cnt * sum
			sum += cnt
		}
	}
	return ans
}

func sortMatrix(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	// 性质 对角线的(i-j)是固定的
	// k = i-j+n
	// i = k+j-n
	// j = i-k+n ;i=m-1; j= m+n-k-1
	// j = i-k+n,;i=0;j=n-k
	for k := 1; k < m+n; k++ {
		minJ := max(0, n-k)       // i=0时候j的取值
		maxJ := min(n-1, m-1+n-k) // i=m-1时 j的取值
		diagnoal := make([]int, 0)
		for j := minJ; j <= maxJ; j++ {
			diagnoal = append(diagnoal, grid[k+j-n][j])
		}
		if minJ > 0 {
			// 上对角线 递增排序
			slices.Sort(diagnoal)
		} else {
			// 下对角线 递减排序
			slices.SortFunc(diagnoal, func(a, b int) int { return b - a })
		}
		for j := minJ; j <= maxJ; j++ {
			grid[k+j-n][j] = diagnoal[j-minJ]
		}
	}
	return grid
}

// 对角线遍历
func gridDiagonal(grid [][]int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
	// 正对角线
	// i-j+n = k; +n的目的是保证k值为正
	// i = k+j-n
	// j = i+n-k
	diagnoal := make([][]int, m+n)
	for k := 1; k < m+n; k++ {
		minJ := max(0, n-k)
		maxJ := min(n-1, m-1+n-k)
		for j := minJ; j <= maxJ; j++ {
			diagnoal[k] = append(diagnoal[k], grid[k+j-n][j])
		}
	}
	// 反对角线
	// i+j = k
	// i = k-j
	// j = k-i
	diagnoal11 := make([][]int, m+n)
	for k := 0; k < m+n-1; k++ {
		minI := max(0, k-(n-1))
		maxI := min(m-1, k)
		for i := minI; i <= maxI; i++ {
			diagnoal11[k] = append(diagnoal11[k], grid[i][k-i])
		}
	}
	fmt.Println(diagnoal[1:])
	fmt.Println(diagnoal11[:len(diagnoal11)-1])
}

// 正对角线排序
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	// k = i-j+n
	for k := 1; k < m+n; k++ {
		minJ := max(0, n-k)
		maxJ := min(n-1, m-1+n-k)
		diagnoal := make([]int, 0)
		for j := minJ; j <= maxJ; j++ {
			diagnoal = append(diagnoal, mat[k+j-n][j])
		}
		slices.Sort(diagnoal)
		for j := minJ; j <= maxJ; j++ {
			mat[k+j-n][j] = diagnoal[j-minJ]
		}
	}
	return mat
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	set := map[int]struct{}{}
	for i := range m {
		ans[i] = make([]int, n)
		for j := range n {
			clear(set)
			for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
				set[grid[x][y]] = struct{}{}
			}
			topLeft := len(set)
			clear(set)
			for x, y := i+1, j+1; x <= m-1 && y <= n-1; x, y = x+1, y+1 {
				set[grid[x][y]] = struct{}{}
			}
			bottomRight := len(set)
			ans[i][j] = abs(topLeft - bottomRight)
		}
	}
	return ans
}

func differenceOfDistinctValues1(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
	}
	set := map[int]struct{}{}
	for k := 1; k < m+n; k++ {
		minJ := max(0, n-k)
		maxJ := min(n-1, m-1+n-k)
		diagnoal := make([]int, 0)
		for j := minJ; j <= maxJ; j++ {
			diagnoal = append(diagnoal, grid[k+j-n][j])
		}
		// 获取对角线数组 计算topLeft && bottomRight
		top := make([]int, len(diagnoal))
		bottom := make([]int, len(diagnoal))
		clear(set)
		for i := 0; i < len(diagnoal); i++ {
			top[i] = len(set)
			set[diagnoal[i]] = struct{}{}
		}
		clear(set)
		for i := len(diagnoal) - 1; i >= 0; i-- {
			bottom[i] = len(set)
			set[diagnoal[i]] = struct{}{}
		}
		for j := minJ; j <= maxJ; j++ {
			ans[k+j-n][j] = abs(top[j-minJ] - bottom[j-minJ])
		}
	}
	return ans
}

func differenceOfDistinctValues2(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
	}
	set := map[int]struct{}{}
	for k := 1; k < m+n; k++ {
		minJ := max(0, n-k)
		maxJ := min(n-1, m-1+n-k)
		clear(set)
		for j := minJ; j <= maxJ; j++ {
			i := k + j - n
			ans[i][j] = len(set)
			set[grid[i][j]] = struct{}{}
		}
		clear(set)
		for j := maxJ; j >= minJ; j-- {
			i := k + j - n
			ans[i][j] = abs(len(set) - ans[i][j])
			set[grid[i][j]] = struct{}{}
		}
	}
	return ans
}

func minAbsoluteDifference(nums []int) int {
	ans := math.MaxInt32
	last1, last2 := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if last2 != -1 {
				ans = min(ans, abs(i-last2))
			}
			last1 = i
		} else if nums[i] == 2 {
			if last1 != -1 {
				ans = min(ans, abs(i-last1))
			}
			last2 = i
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

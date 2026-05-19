package main

import (
	"math"
	"slices"
	"sort"
)

func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	p := make([]int, n+1) // prefixSum
	mn, mx := 0, 0
	for i := range nums {
		p[i+1] = p[i] + nums[i]
		mn = min(p[i+1], mn)
		mx = max(p[i+1], mx)
	}
	return mx - mn
}

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	// s -> t 的 distance
	n := len(nextCost)
	nextPreSum := make([]int, 2*n+1)
	prevPreSum := make([]int, 2*n+1)
	for i := range 2 * n {
		nextPreSum[i+1] = nextPreSum[i] + nextCost[i%n]
		prevPreSum[i+1] = prevPreSum[i] + previousCost[i%n]
	}
	var ans int64
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		cur, target := int(s[i]-'a'), int(t[i]-'a')
		nextCnt := (target - cur + n) % n
		prevCnt := (cur - target + n) % n
		nextCost := nextPreSum[cur+nextCnt] - nextPreSum[cur]
		prevCost := prevPreSum[cur+1+n] - prevPreSum[cur+1+n-prevCnt]
		ans += int64(min(prevCost, nextCost))
	}
	return ans
}

// [low,high] 中的正奇数个数，等于 [1,high] 中的正奇数个数，减去 [1,low−1] 中的正奇数个数。（这个想法类似 前缀和)
// [1,n] 中正奇数的个数 n+1/2
func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}

func shiftingLetters(S string, shifts []int) string {
	s := []byte(S)
	n := len(shifts)
	var total int
	for i := n - 1; i >= 0; i-- {
		total += shifts[i] // 累加total值
		cur := int(s[i] - 'a')
		s[i] = byte('a' + (cur+total)%26)
	}
	return string(s)
}

func numOfSubarrays(arr []int) int {
	const mod = 1e9 + 7
	var ans int
	odd, even := 0, 1
	pre := 0
	// pre[k] 前k数组是多少个奇数
	for _, num := range arr {
		pre += num % 2
		if pre%2 == 1 {
			ans = (ans + even) % mod
			odd++
		} else {
			ans = (ans + odd) % mod
			even++
		}
	}
	return ans % mod
}

func subarraysDivByK(nums []int, k int) int {
	var ans int
	cnt := map[int]int{}
	cnt[0] = 1
	pre := 0
	for _, num := range nums {
		pre += num
		mod := (pre%k + k) % k
		if c, ok := cnt[mod]; ok {
			ans += c
		}
		cnt[mod]++
	}
	return ans
}

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	cnt := map[int]int{}
	cnt[0] = -1
	pre := 0
	for i, num := range nums {
		pre += num
		mod := (pre%k + k) % k
		if idx, ok := cnt[mod]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			cnt[mod] = i
		}
	}
	return false
}

// 用位运算来代替那些复杂操作
// 前缀和搭配哈希表计数
func beautifulSubarrays(nums []int) int64 {
	var ans int64
	s := 0
	cnt := map[int]int{}
	cnt[0] = 1
	for _, num := range nums {
		s ^= num
		ans += int64(cnt[s])
		cnt[s]++
	}
	return ans
}

// 将0看做-1 计算nums的前缀和得到sum数组
// 遍历sum数组，找到sum[i]=sum[j] 中max(i-j)
func findMaxLength(nums []int) int {
	var ans int
	pos := map[int]int{0: -1}
	s := 0
	for i, num := range nums {
		s += num*2 - 1
		if j, ok := pos[s]; ok {
			ans = max(ans, i-j)
		} else {
			pos[s] = i
		}
	}
	return ans
}

// 将0看做-1 计算nums的前缀和得到sum数组
// 遍历sum数组，找到sum[i]=sum[j]中max(i-j)
// 将ans = s[j+1:i+1]
func findLongestSubarray(s []string) []string {
	var ans []string
	pre := 0
	pos := map[int]int{0: -1}
	for i := 0; i < len(s); i++ {
		r := rune(s[i][0])
		if r >= '0' && r <= '9' {
			pre += 1
		} else {
			pre -= 1
		}
		if j, ok := pos[pre]; ok {
			if len(ans) < i-j {
				ans = s[j+1 : i+1]
			}
		} else {
			pos[pre] = i
		}
	}
	return ans
}

// 将两个状态合二为一 (处理方法还是很厉害)
func maxBalancedSubarray(nums []int) int {
	var ans int
	type pair struct{ xor, diff int }
	pos := map[pair]int{{}: -1}
	pre := pair{}
	for i, num := range nums {
		pre.xor ^= num
		pre.diff += num%2*2 - 1
		if j, ok := pos[pre]; ok {
			ans = max(ans, i-j)
		} else {
			pos[pre] = i
		}
	}
	return ans
}

func maximumSubarraySum(nums []int, k int) int64 {
	ans := math.MinInt64
	n := len(nums)
	pre := make([]int, n+1)
	for i, num := range nums {
		pre[i+1] = pre[i] + num
	}
	// 计算前缀和
	pos := map[int]int{}
	for j := 0; j < n; j++ {
		x := nums[j]
		v1 := x - k
		v2 := x + k
		if i, ok := pos[v1]; ok {
			ans = max(ans, pre[j+1]-pre[i])
		}
		if i, ok := pos[v2]; ok {
			ans = max(ans, pre[j+1]-pre[i])
		}
		// 这里是关键, 如果没有存在这个值那么添加到map中
		// 如果当前值的前缀和 < 之前出现的前缀和小 那么就更新
		if old, ok := pos[x]; !ok || pre[old] > pre[j] {
			pos[x] = j
		}
	}
	if ans == math.MinInt64 {
		return 0
	}
	return int64(ans)
}

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	ans := math.MaxInt
	dp := make([]int, n+1) // dp[i] 表示前i个元素中满足条件的最短子数组长度
	for i := range dp {
		dp[i] = math.MaxInt
	}
	pos := make(map[int]int) // 记录前缀和出现的位置
	pos[0] = 0               // pre = 0 的是前0个元素和
	p := 0
	for i := 0; i < n; i++ {
		p += arr[i]
		if j, ok := pos[p-target]; ok {
			curLen := i - j + 1
			if dp[j] != math.MaxInt {
				ans = min(ans, dp[j]+curLen)
			}
			dp[i+1] = min(dp[i], curLen)
		} else {
			dp[i+1] = dp[i]
		}
		pos[p] = i + 1
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func countStableSubarrays(capacity []int) int64 {
	type pair struct {
		cap int
		pre int
	}
	var ans int64
	cnt := map[pair]int{}
	sum := capacity[0]
	// c[l] = c[r]
	// c[l] + sum[l+1] = sum[r]
	for r := 1; r < len(capacity); r++ {
		ans += int64(cnt[pair{capacity[r], sum}])
		cnt[pair{capacity[r-1], sum + capacity[r-1]}]++
		sum += capacity[r]
	}
	return ans
}

func maxSubarraySum(nums []int, k int) int64 {
	// 数组长度可以被k整除
	n := len(nums)
	// j 可以取 -1，表示子数组从 nums[0] 开始，此时 pre[-1] = 0。
	// j = -1 时，子数组长度为 i - (-1) = i + 1。
	// 为了让长度被 k 整除，需要 i + 1 是 k 的倍数，即 i % k == k - 1。
	// 因此，余数为 k-1 的最小前缀和初始值必须是 0（代表空前缀和）。
	// 其他余数没有天然的“空前缀和”，应设成 +∞，避免被误用。
	// 认为 pre[-1] = 0
	minPre := make([]int, k) // 记录同余数相同下最小的前缀和
	// 初始状态下 k-1对应的同余的前缀和是0 其余设置为max表示未知状态
	for i := range k - 1 { // k-1 的目的是
		minPre[i] = math.MaxInt / 2
	}
	ans := math.MinInt
	pre := 0
	// 累加前缀和,
	// 维护一个idx 同余的最小前缀和
	// 可以计算 curPrefixSum - min(idxPrefixSum), 维护最大值
	for i := 0; i < n; i++ {
		pre += nums[i]
		j := i % k // 余数相同 就有(i-j)%k==0
		ans = max(ans, pre-minPre[j])
		minPre[j] = min(minPre[j], pre)
	}
	return int64(ans)
}

// leetcode 2602
// 贴一个题解 画图解析很直观, 前缀和就是一种提效的工具
// https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solutions/2191417/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/
func minOperations1(nums []int, queries []int) []int64 {
	n := len(nums)
	slices.Sort(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	} // 计算前缀和
	ans := make([]int64, len(queries))
	for i := range queries {
		x := queries[i]
		// 找到q的位置 返回>=x的最小值
		j := sort.SearchInts(nums, x)
		leftPart := x*j - pre[j]
		rightPart := pre[n] - pre[j] - (n-j)*x
		ans[i] = int64(leftPart) + int64(rightPart)
	}
	return ans
}

// leetcode 1685
func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, x := range nums {
		total += x
	}
	ans := make([]int, n)
	leftSum := 0
	for i, x := range nums {
		// 这里将模拟运算抽象成了数学公式来做
		leftPart := i*x - leftSum
		rightSum := total - leftSum - x
		rightPart := rightSum - (n-1-i)*x
		ans[i] = leftPart + rightPart
		leftSum += x
	}
	return ans
}

func distance(nums []int) []int64 {
	n := len(nums)
	// 按照相同元素分组
	groups := make(map[int][]int)
	for i, x := range nums {
		groups[x] = append(groups[x], i) // value: index 添加后是天然递增的
	}
	ans := make([]int64, n)
	for _, g := range groups {
		// 计算分组的前缀和
		m := len(g) // index arr
		s := make([]int, m+1)
		for i, x := range g { //
			s[i+1] = s[i] + x
		} // 计算前缀和
		for i, x := range g {
			left := x*i - s[i]
			right := s[m] - s[i] - (m-i)*x
			ans[x] = int64(left + right)
		}
	}
	return ans
}

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	ans := make([]bool, n)
	return ans
}

package main

import "math"

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

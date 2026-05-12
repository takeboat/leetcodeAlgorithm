package main

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

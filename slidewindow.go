package main

import (
	"math"
	"slices"
	"sort"
)

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	// n+1 段的空闲时间
	// 空闲时间可以为0
	free := make([]int, n+1)
	free[0] = startTime[0]
	for i := 1; i < n; i++ {
		free[i] = startTime[i] - endTime[i-1]
	}
	free[n] = eventTime - endTime[n-1]
	var ans int
	// 滑动窗口的模板
	s := 0
	for index, f := range free {
		s += f
		if index < k {
			continue
		}
		ans = max(ans, s)
		s -= free[index-k]
	}
	return ans
}

type position struct {
	x int
	y int
}

var dir = []position{'U': {0, 1}, 'D': {0, -1}, 'L': {1, 0}, 'R': {-1, 0}}

// 查看偏移量
func distinctPoints(s string, k int) int {
	if len(s) < k {
		return 1
	}
	set := map[position]struct{}{}
	p := position{}
	for i, c := range s {
		p.x += dir[c].x
		p.y += dir[c].y
		left := i - k + 1
		if left < 0 {
			continue
		}
		set[p] = struct{}{}
		p.x -= dir[s[left]].x
		p.y -= dir[s[left]].y
	}
	return len(set)
}

// 循环数组
func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	size := abs(k)
	sum := 0
	start := 0
	if k > 0 {
		start = 1
	} else {
		start = n - size
	}
	for i := range size {
		sum += code[(start+i)%n]
	}
	ans[0] = sum
	for i := 1; i < n; i++ {
		out := code[(start+i-1)%n]
		in := code[(start+i+size-1)%n]
		sum = sum - out + in
		ans[i] = sum
	}
	return ans
}

// 计算不变换的总收益
// 计算变换过程中的最大收益
// (newStrategy - oldStrategy) * price
func amaxProfit(prices []int, strategy []int, k int) int64 {
	totalProfit := 0
	sum, maxSum := 0, 0
	for i := range k / 2 {
		totalProfit += prices[i] * strategy[i]
		sum -= prices[i] * strategy[i]
	}
	for i := k / 2; i < k; i++ {
		totalProfit += prices[i] * strategy[i]
		sum += prices[i] * (1 - strategy[i])
	}
	maxSum = max(maxSum, sum)
	for i := k; i < len(prices); i++ {
		totalProfit += prices[i] * strategy[i]
		in := prices[i] * (1 - strategy[i])
		out := prices[i-k/2] - prices[i-k]*strategy[i-k]
		sum += in - out
		maxSum = max(maxSum, sum)
	}
	return int64(totalProfit + maxSum)
}

// getSubarrayBeauty 计算子数组的美丽值
func getSubarrayBeauty(nums []int, k int, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] {
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	// 统计第n小的数据
	for i, num := range nums[k-1:] {
		cnt[num+bias]++
		left := x
		// 计数排序的思想来 计算前x个的最小
		for v, c := range cnt[:bias] {
			left -= c
			if left <= 0 {
				ans[i] = v - bias
				break
			}
		}
		cnt[nums[i]+bias]--
	}
	return ans
}

func minOperations(nums []int, x int) int {
	target := 0
	for _, num := range nums {
		target += num
	}
	target -= x
	if target < 0 {
		return -1
	}
	ans, left, sum := -1, 0, 0
	for right, num := range nums {
		sum += num
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			ans = max(ans, right-left+1)
		}
	}
	if ans < 0 {
		return -1
	}
	return len(nums) - ans
}

func maxFrequency(nums []int, k int) int {
	ans := 1
	slices.Sort(nums) // 瓶颈在排序
	var left int
	for right := 1; right < len(nums); right++ {
		size := right - left + 1
		k -= (nums[right] - nums[right-1]) * (size - 1)
		for k < 0 {
			k += nums[right] - nums[left]
			left++
			// 同步更新size
			size = right - left + 1
		}
		ans = max(ans, size)
	}
	return ans
}

// longestEqualSubarray 最长的等值子数组
func longestEqualSubarray(nums []int, k int) int {
	var ans int
	posLists := make([][]int, len(nums)+1)
	for i, x := range nums {
		posLists[x] = append(posLists[x], i)
	}
	for _, pos := range posLists {
		if len(pos) <= ans {
			continue
		}
		left := 0
		for right := range pos {
			for pos[right]-pos[left]-(right-left) > k {
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return ans
}

// maximumWhiteTiles
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	// 难点在于 比较难以想象最优解是枚举右端点
	var ans int
	left, cover := 0, 0
	for _, tile := range tiles {
		cover += tile[1] - tile[0] + 1
		carpetLeft := tile[1] - carpetLen + 1
		// 不再覆盖left的内容
		for carpetLeft > tiles[left][1] {
			cover -= tiles[left][1] - tiles[left][0] + 1
			left++
		}
		uncover := max(carpetLeft-tiles[left][0], 0)
		ans = max(ans, cover-uncover)
	}
	return ans
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	ans, sum := 0, 0
	left := 0
	for right, fruit := range fruits {
		posRight := fruit[0]
		sum += fruit[1]
		// cost = 窗口大小 + min(startPos离左右窗口的距离)
		for left <= right &&
			((posRight-fruits[left][0])+min(abs(posRight-startPos), abs(fruits[left][0]-startPos))) > k {
			sum -= fruits[left][1]
			left++
		}
		ans = max(ans, sum)
	}
	return ans
}

func minLength(nums []int, k int) int {
	var sum, left int
	ans := math.MaxInt32
	n := len(nums)
	numCnt := make([]int, 100001)
	for right := 0; right < n; right++ {
		numCnt[nums[right]]++         // 计数+1
		if numCnt[nums[right]] <= 1 { // 首次添加进来
			sum += nums[right]
		}
		for sum >= k && left <= right {
			ans = min(ans, right-left+1)
			numCnt[nums[left]]--
			if numCnt[nums[left]] == 0 {
				sum -= nums[left]
			}
			left++
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func shortestBeautifulSubstring(s string, k int) string {
	var ans string
	var left int
	var cnt int
	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			cnt++
		}
		for cnt >= k && left <= right {
			if cnt == k {
				str := s[left : right+1]
				// 主要是字典序判断
				if ans == "" || len(str) < len(ans) || (len(str) == len(ans) && str < ans) {
					ans = str
				}
			}
			if s[left] == '1' {
				cnt--
			}
			left++
		}
	}
	return ans
}

func minSizeSubarray(nums []int, target int) int {
	sum := 0 // 数组和
	for i := range nums {
		sum += nums[i]
	}
	mul := target / sum
	target = target % sum
	n := len(nums)
	doubled := make([]int, 2*n)
	copy(doubled, nums)
	copy(doubled[n:], nums)
	left, s1 := 0, 0
	ans := math.MaxInt32
	for right := 0; right < 2*n; right++ {
		s1 += doubled[right]
		for s1 > target && left <= right {
			s1 -= doubled[left]
			left++
		}
		if s1 == target {
			length := right - left + 1
			if length <= n { // 长度不能超过原数组长度
				ans = min(ans, length)
			}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans + mul*len(nums)
}

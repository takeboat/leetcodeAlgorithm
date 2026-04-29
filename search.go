package main

import (
	"container/heap"
	"math"
	"math/bits"
	"slices"
	"sort"
	"strings"
)

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, slices.Max(nums)+1
	check := func(m int) bool {
		sum := 0
		for _, num := range nums {
			sum += (num + m - 1) / num
			if sum > threshold {
				return false
			}
		}
		return true
	}
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func minimumTime(time []int, totalTrips int) int64 {
	check := func(mid int) bool {
		trips := 0
		for _, t := range time {
			trips += mid / t
			if trips >= totalTrips {
				return true
			}
		}
		return false
	}
	left, right := 1, slices.Min(time)*totalTrips+1
	for left < right {
		mid := (right-left)/2 + left
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

func shipWithinDays(weights []int, days int) int {
	check := func(weight int) bool {
		d := 1
		sum := 0
		for _, w := range weights {
			if sum+w > weight {
				d++
				sum = 0
			}
			sum += w
		}
		return d <= days
	}
	left := slices.Max(weights)
	right := 1
	for _, weight := range weights {
		right += weight
	}
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// dist 是行驶距离
func minSpeedOnTime(dist []int, hour float64) int {
	check := func(speed int) bool {
		totalTime := 0.0
		n := len(dist)
		for i := 0; i < n-1; i++ {
			totalTime += math.Ceil(float64(dist[i]) / float64(speed))
		}
		totalTime += float64(dist[n-1]) / float64(speed)
		return totalTime <= hour
	}
	if hour <= float64(len(dist)-1) {
		return -1
	}
	left, right := 1, 10000000
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func separateSquares(squares [][]int) float64 {
	left := 0
	right := 0
	for _, square := range squares {
		right = max(right, square[1]+square[2])
	}
	const m = 100_000
	check := func(y int) bool {
		up, down := 0, 0
		for _, square := range squares {
			if intersects(y, square[1], square[2]) {
				down += (y - square[1]) * square[2] * m
				up += (square[1] + square[2] - y) * square[2] * m
			} else if y > square[1] {
				down += square[2] * square[2] * m
			} else {
				up += square[2] * square[2] * m
			}
		}
		return up == down
	}
	for left < right*m {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return float64(right / m)
}

func intersects(y int, y1 int, edgeLength int) bool {
	return y1+edgeLength > y
}

// h篇论文被引用了h次
func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}
	// [) 左闭右开区间
	left, right := 1, n+1
	check := func(index int) bool {
		count := 0
		for _, citation := range citations {
			if citation >= index {
				count++
			}
		}
		return count >= index
	}
	// 查找
	// 查找最大值 需要在满足条件的情况下 放大
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func maximumCandies(candies []int, k int64) int {
	totalCandies := 0
	maxCandies := 0
	for _, candy := range candies {
		totalCandies += candy
		maxCandies = max(maxCandies, candy)
	}
	// low 至少分得的糖果数量
	check := func(low int) bool {
		cnt := 0
		for _, candy := range candies {
			cnt += candy / low
		}
		return cnt >= int(k)
	}
	left, right := 1, maxCandies+1
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxW := 0
	for _, w := range workerTimes {
		if w > maxW {
			maxW = w
		}
	}
	// 右界用 int64
	// n*(n+1)/2
	left, right := int64(1), int64(maxW)*int64(mountainHeight)*int64(mountainHeight+1)/2

	check := func(t int64) bool {
		sum := int64(0)
		for _, wt := range workerTimes {
			limit := t / int64(wt)
			// 整数 sqrt(8*limit + 1)
			s := int64(math.Sqrt(float64(8*limit + 1)))
			for s*s > 8*limit+1 {
				s--
			} // 保证 s² ≤ 8*limit+1
			k := (s - 1) / 2
			sum += k
			if sum >= int64(mountainHeight) {
				return true
			}
		}
		return false
	}

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 活跃字符串的最小时间
// 二分答案法
func minTime(s string, order []int, k int) int {
	f := func(n int) int {
		return n * (n + 1) / 2
	}
	total := f(len(s))
	if total < k {
		return -1
	}
	// 计算子字符串的数量
	check := func(t int) bool {
		cnt := 0
		r := []rune(s)
		for i := 0; i <= t; i++ {
			r[order[i]] = '*'
		}
		for part := range strings.SplitSeq(string(r), "*") {
			cnt += f(len(part))
		}
		return total-cnt >= k
	}
	left, right := 0, len(s)
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func repairCars(ranks []int, cars int) int64 {
	// r * n * n 时间内修好 n 量汽车
	check := func(t int) bool {
		cnt := 0
		for _, rank := range ranks {
			cnt += int(math.Sqrt(float64(t / rank)))
			if cnt >= cars {
				return true
			}
		}
		return false
	}
	left, right := 1, slices.Max(ranks)*cars*cars+1
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(right)
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

/*
type SnapshotArray struct {
	data   map[int][]pair // key代表的是 length, value 对应的是index的值
	snapID int
}
type pair struct {
	snapID int
	value  int
}

func Constructor(length int) SnapshotArray {
	data := make(map[int][]pair)
	for i := range length {
		data[i] = []pair{{snapID: 0, value: 0}}
	}
	return SnapshotArray{data: data, snapID: 0}
}

func (sna *SnapshotArray) Set(index int, val int) {
	// fetch the pair slice
	history := sna.data[index]
	last := history[len(history)-1]
	// 如果相同
	if sna.snapID == last.snapID {
		history[len(history)-1] = pair{snapID: sna.snapID, value: val}
	} else {
		history = append(history, pair{snapID: sna.snapID, value: val})
	}
	// set the pair slice
	sna.data[index] = history
}

func (sna *SnapshotArray) Snap() int {
	sna.snapID++
	return sna.snapID - 1
}

func (sna *SnapshotArray) Get(index int, snap int) int {
	// 二分查找
	history := sna.data[index]
	j := sort.Search(len(history), func(i int) bool {
		// sort.Search 会向右收缩
		return history[i].snapID <= snap+1
	}) - 1
	// 全都是大于snap
	// 但是这里的情况可能不会出现
	// 加入边界情况 防止出现错误
	if j >= 0 {
		return history[j].value
	}
	return 0
}

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if !f(h) {
			j = h
		} else {
			i = h + 1
		}
	}
	return i
}
*/

/*
type pair struct {
	value     string
	timestamp int
}
type TimeMap struct {
	m map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]pair)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.m[key] = append(tm.m[key], pair{value: value, timestamp: timestamp})
	return
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	h := tm.m[key]
	l, r := 0, len(h)
	for l < r {
		mid := int(uint32(l+r) >> 1)
		if h[mid].timestamp > timestamp {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		return ""
	}
	return h[l-1].value
}
*/

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

/*
	type Router struct {
		limit           int
		queue           []packet
		packeSet        map[packet]struct{}
		destToTimestamp map[int][]int
	}

	type packet struct {
		source      int
		destination int
		timestamp   int
	}

	func Constructor(memoryLimit int) Router {
		return Router{
			limit:           memoryLimit,
			queue:           make([]packet, 0),
			packeSet:        make(map[packet]struct{}),
			destToTimestamp: make(map[int][]int),
		}
	}

	func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
		pkt := packet{source: source, destination: destination, timestamp: timestamp}
		// 有相同的不添加
		if _, ok := r.packeSet[pkt]; ok {
			return false
		}
		r.packeSet[pkt] = struct{}{}
		if len(r.queue) == r.limit {
			r.ForwardPacket()
		}
		r.queue = append(r.queue, pkt)
		r.destToTimestamp[destination] = append(r.destToTimestamp[destination], timestamp) // 头部添加
		return true
	}

	func (r *Router) ForwardPacket() []int {
		if len(r.queue) == 0 {
			return nil
		}
		pkt := r.queue[0]
		r.queue = r.queue[1:] // 出队
		r.destToTimestamp[pkt.destination] = r.destToTimestamp[pkt.destination][1:]
		delete(r.packeSet, pkt)
		return []int{pkt.source, pkt.destination, pkt.timestamp}
	}

	func (r *Router) GetCount(destination int, startTime int, endTime int) int {
		timestamps := r.destToTimestamp[destination]
		return sort.SearchInts(timestamps, endTime+1) - sort.SearchInts(timestamps, startTime)
	}
*/

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		l, r := bits.OnesCount32(uint32(arr[i])), bits.OnesCount32(uint32(arr[j]))
		if l == r {
			return arr[i] < arr[j]
		}
		return l < r
	})
	return arr
}

func maximumLength(s string) int {
	groups := [26][]int{}
	cnt := 0
	// 统计连续字符的数量
	for i := range s {
		cnt++
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], cnt)
			cnt = 0
		}
	}
	// 示例
	// 若有 s: aaabbabccc
	// a: 3, 1,
	// b: 2, 1
	// c: 3
	ans := 0
	for _, a := range groups {
		if len(a) == 0 {
			continue
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })
		a = append(a, 0, 0) // 假设有两个空串
		ans = max(ans, a[0]-2, min(a[0]-1, a[1]), a[2])
	}
	if ans == 0 {
		return -1
	}
	return ans
}

// 至少出现出现k次的最长特殊子字符串
func maximumLengthHelp(s string, k int) int {
	groups := [26][]int{}
	cnt := 0
	// 统计连续字符的数量
	for i := range s {
		cnt++
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], cnt)
			cnt = 0
		}
	}
	ans := 0
	for _, a := range groups {
		if len(a) == 0 {
			continue
		}
		// 二分答案区间 [)
		left := 1
		right := slices.Max(a) + 1

		// check函数主要是为了校验 mid长度的子字符串在s中能否出现至少k次
		check := func(mid int) bool {
			cnt := 0
			for _, v := range a {
				if v >= mid {
					cnt += v - mid + 1
				}
				if cnt >= k {
					return true
				}
			}
			return false
		}

		for left < right {
			mid := int(uint(left+right) >> 1)
			ans = max(ans, mid)
			if check(mid) {
				left = mid + 1 // 尝试更长
			} else {
				right = mid
			}
		}
	}
	if ans == 0 {
		return -1
	}
	return ans
}

func numSteps(s string) int {
	ans := len(s) - 1
	carry := 0
	for i := len(s) - 1; i > 0; i-- {
		sum := int(s[i]-'0') + carry
		ans += sum % 2
		carry = (sum + sum%2) / 2
	}
	// 如果carry == 1 还需要执行一次 /2 操作 操作数要+1
	return ans + carry
}

func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	// 检查是否能形成 k 对（即 2k 个标记下标）
	check := func(k int) bool {
		// 尝试将前 k 个小的和后 k 个大的配对
		for i := 0; i < k; i++ {
			// 如果存在任何一对不满足 nums[i]*2 <= nums[n-k+i]
			// 说明 k 对不可行
			if nums[i]*2 > nums[n-k+i] {
				return false
			}
		}
		return true
	}
	// 二分查找最大的可行对数 k
	// 对数范围 [0, n/2]
	left, right := 0, n/2+1
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			// mid 对可行，尝试更多
			left = mid + 1
		} else {
			// mid 对不可行，尝试更少
			right = mid
		}
	}
	return (left - 1) * 2
}

func maxNumOfMarkedIndices1(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	check := func(k int) bool {
		for i := 0; i < k; i++ {
			if nums[i]*2 > nums[n-k+i] {
				return false
			}
		}
		return true
	}
	left, right := 0, n/2+1
	for left+1 < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left * 2
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	mod := 1_000_000_007
	n := len(nums1)
	sortNums1 := make([]int, n)
	copy(sortNums1, nums1)
	slices.Sort(sortNums1)
	total := 0
	maximprovement := 0
	for i := range nums2 {
		originalDiff := abs(nums1[i] - nums2[i])
		total += originalDiff
		l, r := 0, n
		// 找到最接近 nums2[i] 的 sortNums1[l]
		for l < r {
			mid := int(uint(l+r) >> 1)
			if sortNums1[mid] < nums2[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		newDiff := math.MaxInt32
		// 第一个>=nums2[i]的sortNums1[l]
		if l < n {
			newDiff = min(newDiff, abs(sortNums1[l]-nums2[i]))
		}
		// 最后一个<nums2[i]的sortNums1[l-1]
		if l > 0 {
			newDiff = min(newDiff, abs(sortNums1[l-1]-nums2[i]))
		}
		maximprovement = max(maximprovement, originalDiff-newDiff)
	}
	return (total - maximprovement) % mod
}

// 好
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	check := func(mid int) bool {
		return x-arr[mid] > arr[mid+k]-x
	}
	for left < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}

type TopVotedCandidate struct {
	tops [][]int
	// element [person,time]
	// array was sorted by time
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	m := map[int]int{} // 记录每个人的票数
	tops := [][]int{}
	maxVote := 0
	for i := range times {
		// 如果map中不存在 那么先插入
		if _, ok := m[persons[i]]; !ok {
			m[persons[i]] = 0
		}
		m[persons[i]]++
		if m[persons[i]] >= maxVote {
			// 如果有人的票数大于等于maxVote 那么需要更新
			// 题目中表明如果票数相同的情况下 时间最近的人胜出
			maxVote = m[persons[i]]
			// 将tops数组中添加一个元素 表示在times[i] 的时刻 persons[i] 获得了最多的票
			tops = append(tops, []int{persons[i], times[i]})
		}
	}
	return TopVotedCandidate{tops: tops}
}

func (tvs *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(tvs.tops), func(i int) bool {
		return tvs.tops[i][1] > t
	})
	return tvs.tops[i-1][0]
}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	left, right := slices.Min(bloomDay), slices.Max(bloomDay)+1
	check := func(mid int) bool {
		total := 0
		cnt := 0
		for i := range bloomDay {
			if bloomDay[i] <= mid {
				cnt++
				if cnt == k {
					total++
					cnt = 0
				}
			} else {
				cnt = 0
			}
		}
		return total >= m
	}
	for left < right {
		mid := int(uint(left+right) >> 1)
		// 从单调性来看 如果check(mid) 为 true 那么 mid 向右都是满足条件的
		// mid 的左边需要二分来找是否满足
		// mid的左边是不满足的 mid的右边都是满足的
		// 那么最后 结束条件是 left = right
		// mid = left = right 那么mid就是满足条件最小的值
		if check(mid) { // 找最小的
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func lowerbound(nums []int) int {
	left, right := 0, len(nums)
	check := func(mid int) bool {
		// just show some example
		return mid > 0
	}
	for left < right {
		mid := int(uint(left+right) >> 1)
		// 这里的单调性看 如果check(mid) 为true 那么mid 的左侧都是满足条件的
		// mid 的右侧是不满足的
		// 那么最后 结束条件是 left = right
		// mid = left = right
		if check(mid) { // 找最大的
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 最后满足条件的位置+1
// if check(mid) {
// 	left = mid + 1
// } else {
// 	right = mid
// }
//
// if check(mid) {
// 	right = mid
// } else {
// 	left = mid + 1
// }

// 剧情触发
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	// c r h 文明 资源 人口
	prefixSum := make([][3]int, len(increase))
	prefixSum[0] = [3]int{increase[0][0], increase[0][1], increase[0][2]}
	for i := 1; i < len(increase); i++ {
		c := prefixSum[i-1][0] + increase[i][0]
		r := prefixSum[i-1][1] + increase[i][1]
		h := prefixSum[i-1][2] + increase[i][2]
		prefixSum[i] = [3]int{c, r, h}
	}
	ans := make([]int, 0)
	for _, requirement := range requirements {
		c, r, h := requirement[0], requirement[1], requirement[2]
		// 如果是这样的话 第 0 天就触发剧情了
		if c == 0 && r == 0 && h == 0 {
			ans = append(ans, 0)
			continue
		}
		left, right := 0, len(increase)
		for left < right {
			mid := int(uint(left+right) >> 1)
			if prefixSum[mid][0] >= c && prefixSum[mid][1] >= r && prefixSum[mid][2] >= h {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == len(increase) {
			ans = append(ans, -1)
		} else {
			ans = append(ans, left+1)
		}
	}
	return ans
}

func minimumK(nums []int) int {
	left, right := 1, 100_001
	check := func(k int) bool {
		cnt := 0
		for _, num := range nums {
			if num%k != 0 {
				cnt += 1
			}
			cnt += num / k
		}
		return cnt <= k*k
	}
	for left < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func numSpecial(mat [][]int) int {
	ans := 0
	isSpecial := func(i, j int) bool {
		for k := range mat[i] {
			if k == j {
				continue
			}
			if mat[i][k] == 1 {
				return false
			}
		}
		for k := range mat {
			if k == i {
				continue
			}
			if mat[k][j] == 1 {
				return false
			}
		}
		return true
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && isSpecial(i, j) {
				ans++
			}
		}
	}
	return ans
}

func numSpecial1(mat [][]int) int {
	rowsOneCnt := make([]int, len(mat))
	columnsOneCnt := make([]int, len(mat[0]))
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				rowsOneCnt[i]++
				columnsOneCnt[j]++
			}
		}
	}
	ans := 0
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && rowsOneCnt[i] == 1 && columnsOneCnt[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	ans := make([][]int, len(queries))
	elements := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		elements = append(elements, node.Val)
		dfs(node.Right)
	}
	for i, query := range queries {
		mini, maxi := -1, -1
		left, right := 0, len(elements)
		for left < right {
			mid := int(uint(left+right) >> 1)
			if elements[mid] >= query {
				right = mid
			} else {
				left = mid + 1
			}
		}
		// left := sort.Search(len(elements), func(i int) bool {
		// 	return elements[i] < query
		// })
		if left < len(elements) {
			maxi = elements[left]
			if maxi == query {
				mini = query
			} else if left > 0 {
				mini = elements[left-1]
			}
		} else if left > 0 {
			mini = elements[left-1]
		}
		ans[i] = []int{mini, maxi}
	}
	return ans
}

func maximumRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)+1
	check := func(k int) bool {
		removed := make([]bool, len(s))
		for i := range k {
			removed[removable[i]] = true
		}
		j := 0
		for i := 0; i < len(s) && j < len(p); i++ {
			if !removed[i] && s[i] == p[j] {
				j++
			}
		}
		return j == len(p)
	}
	for left < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func checkOnesSegment(s string) bool {
	if len(s) < 2 {
		return true
	}
	for i := 1; i < len(s); i++ {
		// 01
		if s[i] == '1' && s[i-1] == '0' {
			return false
		}
	}
	return true
}

func maxValue(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}
	left, right := 1, maxSum+1
	// k 代表 nums[index] 从k-1中递减,最多cnt项
	calc := func(k, cnt int) int {
		if cnt <= 0 {
			return 0
		}
		if k > cnt {
			return cnt * (2*k - cnt - 1) / 2 // 等差数列求和公式
		}
		return (k-1)*k/2 + cnt - k + 1 // 求到1的求和 nums是正整数
	}
	check := func(k int) bool {
		var sum int
		left := index // 左侧元素个数
		right := n - index - 1
		sum += k + calc(k, left) + calc(k, right)
		return sum <= maxSum
	}
	for left < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// left 满足条件的最小值 => left-1 不满足条件的最大值
	return left - 1
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func furthestBuilding1(heights []int, bricks int, ladders int) int {
	maxHeap := &MaxIntHeap{}
	heap.Init(maxHeap)
	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff <= 0 {
			continue
		}
		bricks -= diff
		heap.Push(maxHeap, diff)
		if bricks < 0 {
			if ladders == 0 {
				return i
			}
			maxDiff := heap.Pop(maxHeap).(int)
			bricks += maxDiff
			ladders--
		}
	}
	return len(heights) - 1
}

type MaxIntHeap []int                    // 大根堆
func (h MaxIntHeap) Len() int            { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinIntHeap []int                    // 小根堆
func (h MinIntHeap) Len() int            { return len(h) }
func (h MinIntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type MaxIntHeap struct {
	sort.IntSlice
}

func (h MaxIntHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *MaxIntHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MaxIntHeap) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 小根堆
type MinIntHeap struct {
	sort.IntSlice
}

// sort.IntSlice 这里其实不用写Less
func (h MinIntHeap) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *MinIntHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MinIntHeap) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 堆实现
// 堆的二叉树结构
/*
type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]int, 0),
	}
}

func (h *MinHeap) up(i int) {
	// 一直比较
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MinHeap) down(i int) {
	n := len(h.data)
	left := 2*i + 1
	for left < n {
		// 比较孩子中小的和父节点比较
		// 如果小的孩子比父节点大 那么就交换
		small := left
		if right := left + 1; right < n && h.data[right] < h.data[left] {
			small = right
		}
		// 如果父节点小于等于左右孩子中小的孩子节点
		if h.data[i] <= h.data[small] {
			break
		}
		// 否则交换
		h.data[i], h.data[small] = h.data[small], h.data[i]
		// 更新子节点坐标
		i = small
		left = 2*i + 1
	}
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

// return == -1 => empty heap
func (h *MinHeap) Pop() int {
	n := len(h.data)
	if n == 0 {
		return -1
	}
	res := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.down(0)
	return res
}
*/

func pickGifts(gifts []int, k int) int64 {
	h := &MaxIntHeap{gifts}
	heap.Init(h)
	for i := 0; i < k; i++ {
		gift := heap.Pop(h).(int)
		heap.Push(h, int(math.Sqrt(float64(gift))))
	}
	var sum int64
	for _, gift := range h.IntSlice {
		sum += int64(gift)
	}
	return sum
}

func lastStoneWeight(stones []int) int {
	h := &MaxIntHeap{stones}
	heap.Init(h) // 建堆结构
	for h.Len() > 1 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x != y {
			heap.Push(h, x-y)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.IntSlice[0]
}

// 最小无限集合
type SmallestInfiniteSet struct {
	h    *MinIntHeap  // 最小堆
	set  map[int]bool // 记录数字是否添加到堆中,防止重复添加
	next int          // 记录下一个的数字
}

func NewSmallestInfiniteSet() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		h:    &MinIntHeap{IntSlice: make([]int, 0)},
		set:  make(map[int]bool),
		next: 1,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.h.Len() > 0 {
		small := heap.Pop(s.h).(int)
		delete(s.set, small)
		return small
	}
	small := s.next
	s.next++
	return small
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.next && !s.set[num] {
		s.set[num] = true
		heap.Push(s.h, num)
	}
}

func maxKelements(nums []int, k int) int64 {
	var ans int
	h := &MaxIntHeap{nums}
	heap.Init(h)
	for range k {
		m := heap.Pop(h).(int)
		ans += m
		heap.Push(h, int((m+2)/3))
		fmt.Printf("%d, %d\n", m, int((m+2)/3))
	}
	return int64(ans)
}

func minOperations3(nums []int, k int) int {
	h := &MinIntHeap{nums}
	heap.Init(h)
	var ans int
	for h.Len() >= 2 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x >= k {
			break
		}
		heap.Push(h, x*2+y)
		ans++
	}
	return ans
}

func minStoneSum(piles []int, k int) int {
	h := &MaxIntHeap{piles}
	heap.Init(h)
	for range k {
		x := heap.Pop(h).(int)
		heap.Push(h, x-x/2)
	}
	var sum int
	for _, x := range h.IntSlice {
		sum += x
	}
	return sum
}

type KthLargest struct {
	h *MinIntHeap
	k int
}

// 第k大元素 => 维护最大的k个元素的最小堆
// add => 如果加入的值小于堆顶值 直接丢弃, 如果加入的值大于堆顶值那么加入之后再pop
func NewKthLargest(k int, nums []int) KthLargest {
	h := &MinIntHeap{sort.IntSlice{}}
	kl := KthLargest{
		h: h,
		k: k,
	}
	for i := range nums {
		kl.Add(nums[i])
	}
	return kl
}

func (l *KthLargest) Add(val int) int {
	heap.Push(l.h, val)
	if l.h.Len() > l.k {
		heap.Pop(l.h)
	}
	return l.h.IntSlice[0]
}

func resultsArray(queries [][]int, k int) []int {
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	h := &MaxIntHeap{sort.IntSlice{}}
	// 使用最大堆维护前k个元素
	for i, q := range queries {
		x, y := q[0], q[1]
		heap.Push(h, abs(x)+abs(y))
		if h.Len() > k {
			heap.Pop(h)
		}
		if h.Len() == k {
			ans[i] = h.IntSlice[0]
		}
	}
	return ans
}

type SeatManager struct {
	h        *MinIntHeap  // 使用小根堆来存储unreserve的座位号 要保持 s.next > unreserve number
	reserved map[int]bool // reserved seat
	next     int          // next seat
}

func NewSeatManager(n int) SeatManager {
	return SeatManager{
		h:        &MinIntHeap{sort.IntSlice{}},
		reserved: make(map[int]bool),
		next:     1,
	}
}

func (s *SeatManager) Reserve() int {
	if s.h.Len() > 0 {
		res := heap.Pop(s.h).(int)
		s.reserved[res] = true
		return res
	}
	res := s.next
	s.next++
	return res
}

func (s *SeatManager) Unreserve(seatNumber int) {
	if seatNumber < s.next {
		s.reserved[seatNumber] = false
		heap.Push(s.h, seatNumber)
	}
}

func maximumProduct1(nums []int, k int) int {
	h := &MinIntHeap{nums}
	ans := 1
	heap.Init(h)
	const mod = 1e9 + 7
	for range k {
		top := heap.Pop(h).(int)
		heap.Push(h, top+1)
	}
	for _, x := range nums {
		ans *= x
		ans %= mod
	}
	return ans % mod
}

func smallestChair(times [][]int, targetFriend int) int {
}

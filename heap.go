package main

import (
	"container/heap"
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

type MinIntHeap struct {
	sort.IntSlice
} // 小根堆
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
res := h.data[0]
	}
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

type SmallestInfiniteSet struct {
	h   *MinIntHeap
	set map[int]bool
	cur int
}

func NewSmallestInfiniteSet() SmallestInfiniteSet {
	return SmallestInfiniteSet{}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
}

func (s *SmallestInfiniteSet) AddBack(num int) {
}

package main

import (
	"math"
	"sort"
)

func leftGreater(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	st := []int{-1} // 哨兵值
	for i, x := range nums {
		// 要插入一个值首先弹出栈中所有 <= 他的值
		for len(st) > 1 && nums[st[len(st)-1]] <= x {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	return left
}

func rightGreater(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	st := []int{n} // 哨兵值
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(st) > 1 && nums[st[len(st)-1]] <= x {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}
	return right
}

// 每日温度
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	st := make([]int, 0) // 从栈底到栈顶单调递减
	for i := 0; i < n; i++ {
		// 如果有了高温天气
		for len(st) > 0 && temperatures[st[len(st)-1]] < temperatures[i] {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			ans[top] = i - top
		}
		st = append(st, i)
	}
	return ans
}

func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0} // 哨兵值
	// 从右向左维持一个递增序列
	// 进行一次循环的时候
	// 找到右侧第一个小于当前值的价格
	for i := n - 1; i >= 0; i-- {
		for st[len(st)-1] > prices[i] {
			st = st[:len(st)-1]
		}
		ans[i] = prices[i] - st[len(st)-1]
		st = append(st, prices[i])
	}
	return ans
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n, m := len(nums1), len(nums2)
	idx := make(map[int]int, n)
	for i, x := range nums1 {
		idx[x] = i
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	for i := m - 1; i >= 0; i-- {
		x := nums2[i]
		// 单调递减
		for len(st) > 0 && st[len(st)-1] < x {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			if j, ok := idx[x]; ok {
				ans[j] = st[len(st)-1]
			}
		}
		st = append(st, x)
	}
	return ans
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range n {
		ans[i] = -1
	}
	st := []int{}
	for i := 2*n - 1; i >= 0; i-- {
		x := nums[i%n]
		// 保持严格递减
		for len(st) > 0 && st[len(st)-1] <= x {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i%n] = st[len(st)-1]
		}
		st = append(st, x)
	}
	return ans
}

// 1,2,3,4,3,1,2,3,4,3
// 1,2,1,1,2,1

// 上一个大于当前元素的值的距离
type pair struct {
	price int
	day   int
}
type StockSpanner struct {
	st  []pair
	cur int // cur day
}

func Constructor3() StockSpanner {
	return StockSpanner{
		st:  []pair{{price: math.MaxInt, day: -1}},
		cur: -1,
	}
}

func (ss *StockSpanner) Next(price int) int {
	ss.cur++
	for ss.st[len(ss.st)-1].price <= price {
		ss.st = ss.st[:len(ss.st)-1]
	}
	ans := ss.cur - ss.st[len(ss.st)-1].day
	ss.st = append(ss.st, pair{price: price, day: ss.cur})
	return ans
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	type car struct {
		pos int
		spd int
	}
	cars := make([]car, 0, n)
	for i := range n {
		cars = append(cars, car{position[i], speed[i]})
	}
	// 按照距离排序
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})
	fleets := 1
	curTime := float64(target-cars[n-1].pos) / float64(cars[n-1].spd)
	for i := n - 2; i >= 0; i-- {
		time := float64(target-cars[i].pos) / float64(cars[i].spd)
		if time > curTime {
			fleets++
			curTime = time
		}
	}
	return fleets
}

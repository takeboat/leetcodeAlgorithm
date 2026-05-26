package main

import (
	"container/list"
	"slices"
)

type RecentCounter struct {
	requests []int
}

func ConstructorRecentCounter() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.requests = append(rc.requests, t)
	for rc.requests[0] < t-3000 {
		rc.requests = rc.requests[1:]
	}
	return len(rc.requests)
}

type RideSharingSystem struct {
	driver       []int
	rider        []int
	waitingRider map[int]bool
}

func NewRideSharingSystem() RideSharingSystem {
	return RideSharingSystem{
		driver:       make([]int, 0),
		rider:        make([]int, 0),
		waitingRider: make(map[int]bool),
	}
}

func (rhs *RideSharingSystem) AddRider(riderId int) {
	rhs.rider = append(rhs.rider, riderId)
	rhs.waitingRider[riderId] = true
}

func (rhs *RideSharingSystem) AddDriver(driverId int) {
	rhs.driver = append(rhs.driver, driverId)
}

func (rhs *RideSharingSystem) MatchDriverWithRider() []int {
	for len(rhs.rider) > 0 && !rhs.waitingRider[rhs.rider[0]] {
		rhs.rider = rhs.rider[1:]
	}
	if len(rhs.rider) == 0 || len(rhs.driver) == 0 {
		return []int{-1, -1}
	}
	rider := rhs.rider[0]
	driver := rhs.driver[0]
	rhs.rider = rhs.rider[1:]
	rhs.driver = rhs.driver[1:]
	return []int{driver, rider}
}

func (rhs *RideSharingSystem) CancelRider(riderId int) {
	delete(rhs.waitingRider, riderId)
}

/*
// 这里换了个概念就是用户等待队列 变为 用户取消队列

	type RideSharingSystem struct {
		driver       []int
		rider        []int
		cancled  map[int]bool
	}

	func (rhs *RideSharingSystem) AddRider(riderId int) {
		rhs.rider = append(rhs.rider, riderId)
		delete(rhs.cancled, riderId)
	}

	func (rhs *RideSharingSystem) AddDriver(driverId int) {
		rhs.driver = append(rhs.driver, driverId)
	}

	func (rhs *RideSharingSystem) MatchDriverWithRider() []int {
		for len(rhs.rider) > 0 && rhs.cancled[rhs.rider[0]] {
			rhs.rider = rhs.rider[1:]
		}
		if len(rhs.rider) == 0 || len(rhs.driver) == 0 {
			return []int{-1, -1}
		}
		rider := rhs.rider[0]
		driver := rhs.driver[0]
		rhs.rider = rhs.rider[1:]
		rhs.driver = rhs.driver[1:]
		return []int{driver, rider}
	}

	func (rhs *RideSharingSystem) CancelRider(riderId int) {
		rhs.cancled[riderId] = true
	}
*/

// o(n^2) 暴力解法
func deckRevealedIncreasing(deck []int) []int {
	slices.Sort(deck)
	// 先排序
	// 按照反操作回归顺序
	n := len(deck)
	queue := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		x := deck[i] // 递减获取值
		if len(queue) < 1 {
			queue = append(queue, x)
			continue
		}
		end := queue[len(queue)-1]
		queue = append([]int{x, end}, queue[:len(queue)-1]...)
	}
	return queue
}

func predictPartyVictory(s string) string {
	winner := []string{"Radiant", "Dire"} // 天辉 夜魔
	dire := 0
	radiant := 0
	// 先统计数量
	for _, x := range s {
		if x == 'D' {
			dire++
		} else {
			radiant++
		}
	}
	disableD := 0
	disableR := 0
	senate := []byte(s)
	// 要循环往复 直至一方的数量为0, 已经被制裁的不可以再上
	for dire > 0 && radiant > 0 {
		for i := range senate {
			if senate[i] == 'B' {
				continue
			}
			if senate[i] == 'D' {
				if disableD > 0 { // 制裁buff
					senate[i] = 'B' // 制裁标记
					disableD--
				} else {
					radiant--
					disableR++
				}
			} else {
				if disableR > 0 {
					senate[i] = 'B'
					disableR--
				} else {
					dire-- // 减少一个dire数量 然后加入一个制裁buff 标记下一个已经制裁
					disableD++
				}
			}
		}
	}
	if radiant > 0 {
		return winner[0]
	}
	return winner[1]
}

type FrontMiddleBackQueue struct {
	left  *list.List
	right *list.List
}

// 维护两个队列 保证len(left) >= len(right) && len(left) <= len(right)+1
func NewFrontMiddleBackQueue() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left.PushFront(val)
	if q.left.Len() == q.right.Len()+2 {
		back := q.left.Back().Value.(int)
		q.left.Remove(q.left.Back())
		q.right.PushFront(back)
	}
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	if q.left.Len() == q.right.Len()+1 {
		q.right.PushFront(q.left.Back().Value.(int))
		q.left.Remove(q.left.Back())
	}
	q.left.PushBack(val)
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.right.PushBack(val)
	if q.left.Len()+1 == q.right.Len() {
		front := q.right.Front().Value.(int)
		q.right.Remove(q.right.Front())
		q.left.PushBack(front)
	}
}

func (q *FrontMiddleBackQueue) PopFront() int {
	if q.left.Len() == 0 {
		return -1
	}
	front := q.left.Front().Value.(int)
	q.left.Remove(q.left.Front())
	if q.left.Len()+1 == q.right.Len() {
		f := q.right.Front().Value.(int)
		q.right.Remove(q.right.Front())
		q.left.PushBack(f)
	}
	return front
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if q.left.Len() == 0 {
		return -1
	}
	mid := q.left.Back().Value.(int)
	q.left.Remove(q.left.Back())
	if q.left.Len()+1 == q.right.Len() {
		f := q.right.Front().Value.(int)
		q.right.Remove(q.right.Front())
		q.left.PushBack(f)
	}
	return mid
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if q.left.Len() == 0 {
		return -1
	}
	if q.right.Len() == 0 {
		back := q.left.Back().Value.(int)
		q.left.Remove(q.left.Back())
		return back
	} else {
		back := q.right.Back().Value.(int)
		q.right.Remove(q.right.Back())
		if q.right.Len()+2 == q.left.Len() {
			b := q.left.Back().Value.(int)
			q.left.Remove(q.left.Back())
			q.right.PushFront(b)
		}
		return back
	}
}

func finalString(s string) string {
	queue := list.New()
	tail := true
	for _, ch := range []byte(s) {
		if ch == 'i' {
			tail = !tail
		} else if tail {
			queue.PushBack(ch)
		} else {
			queue.PushFront(ch)
		}
	}
	ans := make([]byte, 0, len(s))
	for i := queue.Front(); i != nil; i = i.Next() {
		ans = append(ans, i.Value.(byte))
	}
	if !tail {
		slices.Reverse(ans)
	}
	return string(ans)
}

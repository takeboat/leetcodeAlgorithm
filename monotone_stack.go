package main

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
	st := make([]int, 0) // 存储value就好了
	// 递增栈(从栈底到栈顶)
	for i := 0; i < n; i++ {
		for len(st) > 0 && st[len(st)-1] >= prices[i] {
			// pop up
			st = st[:len(st)-1]
		}
		// ans[i] =
	}
	return ans
}

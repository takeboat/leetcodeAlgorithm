package main

import (
	"slices"
)

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	// 结构化数据 方便查找
	learned := make([][]bool, m)
	for i := 0; i < m; i++ {
		learned[i] = make([]bool, n+1)
		for _, j := range languages[i] {
			learned[i][j] = true
		}
	}
	todolist := [][2]int{}
	// 找到不能够沟通的两人
	// 遍历朋友关系 查询这两者之间是否学会同种语言，有同种语言的可以沟通；没有的不可以沟通,加入到todolist中
next:
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1
		for _, x := range languages[u] {
			if learned[v][x] {
				continue next
			}
		}
		todolist = append(todolist, [2]int{u, v})
	}
	ans := m // 最坏情况 每个人都需要学一种新的语言才可以沟通

	// 需要教会的语言k
	for k := 1; k <= n; k++ {
		set := map[int]struct{}{}
		// 遍历不能沟通的两人: 找到这两人是否都会语言k
		// 如果其中一人不会语言k,则将其加入到set中
		for _, f := range todolist {
			u, v := f[0], f[1]
			// u需要学习语言k
			if !learned[u][k] {
				set[u] = struct{}{}
			}
			// v需要学习语言k
			if !learned[v][k] {
				set[v] = struct{}{}
			}
		}
		ans = min(ans, len(set))
	}
	return ans
}

func minimumTeachings1(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	// 结构化数据 方便查找
	learned := make([][]bool, m)
	for i := 0; i < m; i++ {
		learned[i] = make([]bool, n+1)
		for _, j := range languages[i] {
			learned[i][j] = true
		}
	}
	// 找到不能够沟通的两人
	total := 0
	vis := make([]bool, m)  // 记录这个人是否被访问过其语言列表
	cnt := make([]int, n+1) // 记录会该语言的人数
	add := func(u int) {
		if vis[u] {
			return
		}
		vis[u] = true
		total++
		for _, x := range languages[u] {
			cnt[x]++
		}
	}
next:
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1
		for _, x := range languages[u] {
			if learned[v][x] {
				continue next
			}
		}
		add(u)
		add(v)
	}
	return total - slices.Max(cnt)
}

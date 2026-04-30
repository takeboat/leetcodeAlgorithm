package main

import (
	"strings"
)

func buildArray(target []int, _ int) []string {
	mx := target[len(target)-1]
	ans := make([]string, 0)
	push := func() { ans = append(ans, "Push") }
	pop := func() { ans = append(ans, "Pop") }
	j := 0
	for i := 1; i <= mx && j < len(target); i++ {
		push()
		if i == target[j] {
			j++
		} else {
			pop()
		}
	}
	return ans
}

func backspaceCompare(s string, t string) bool {
	build := func(s string) string {
		stack := []rune{}
		for _, r := range s {
			if r != '#' {
				stack = append(stack, r)
			} else if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
		return string(stack)
	}
	return build(s) == build(t)
}

func removeStars(s string) string {
	stack := []rune{}
	for _, r := range s {
		if r != '*' {
			stack = append(stack, r)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

// use `stack` to impl BrowserHistory
type BrowserHistory struct {
	his []string
	cur int
}

func Constructor1(homepage string) BrowserHistory {
	return BrowserHistory{his: []string{homepage}, cur: 0}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.cur++
	bh.his = bh.his[:bh.cur]
	bh.his = append(bh.his, url)
}

func (bh *BrowserHistory) Back(steps int) string {
	bh.cur = max(0, bh.cur-steps)
	return bh.his[bh.cur]
}

func (bh *BrowserHistory) Forward(steps int) string {
	bh.cur = min(len(bh.his)-1, bh.cur+steps)
	return bh.his[bh.cur]
}

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	stack := []int{}
	for i := 0; i < n; i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(stack) == 0
}

// 创建了26个栈 为每个字母创建一个栈
func calculateScore(o string) (ans int64) {
	stack := [26][]int{}
	for i, c := range o {
		c -= 'a'
		if st := stack[25-c]; len(st) > 0 {
			ans += int64(i - st[len(st)-1])
			stack[25-c] = st[:len(st)-1]
		} else {
			stack[c] = append(stack[c], i)
		}
	}
	return
}

// 26栈玩法
func clearStars(s string) string {
	stack := [26][]int{}
	deleted := make([]bool, len(s))
	for i, c := range s {
		if c == '*' {
			for j := range 26 {
				st := stack[j]
				if len(st) > 0 {
					deleted[st[len(st)-1]] = true
					stack[j] = st[:len(st)-1]
					break
				}
			}
		} else {
			c -= 'a'
			stack[c] = append(stack[c], i)
		}
	}
	var sb strings.Builder
	for i, c := range s {
		if c != '*' && !deleted[i] {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func clearStars1(S string) string {
	s := []byte(S)
	stack := [26][]int{}
	for i, c := range s {
		if c == '*' {
			for j := range 26 {
				st := stack[j]
				if len(st) > 0 {
					pos := st[len(st)-1]
					s[pos] = '*'
					stack[j] = st[:len(st)-1]
					break
				}
			}
		} else {
			c -= 'a'
			stack[c] = append(stack[c], i)
		}
	}
	var sb strings.Builder
	for _, c := range s {
		if c != '*' {
			sb.WriteByte(c)
		}
	}
	return sb.String()
}

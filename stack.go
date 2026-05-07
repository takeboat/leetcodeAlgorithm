package main

import (
	"math"
	"strconv"
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

type CustomStack struct {
	stack   []int
	maxSize int
}

func Constructor2(maxSize int) CustomStack {
	return CustomStack{
		stack:   make([]int, 0),
		maxSize: maxSize,
	}
}

func (cs *CustomStack) Push(x int) {
	if len(cs.stack) < cs.maxSize {
		cs.stack = append(cs.stack, x)
	}
}

func (cs *CustomStack) Pop() int {
	if len(cs.stack) == 0 {
		return -1
	}
	x := cs.stack[len(cs.stack)-1]
	cs.stack = cs.stack[:len(cs.stack)-1]
	return x
}

func (cs *CustomStack) Increment(k int, val int) {
	for i := 0; i < len(cs.stack) && i < k; i++ {
		cs.stack[i] += val
	}
}

func exclusiveTime(n int, logs []string) []int {
	time := make([]int, n)
	type item struct {
		id int
		t  int
	}
	start := []item{}
	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		event := parts[1]
		t, _ := strconv.Atoi(parts[2])
		if event == "start" {
			if len(start) > 0 {
				prev := start[len(start)-1]
				time[prev.id] += t - prev.t
			}
			start = append(start, item{id, t})
		} else {
			// 弹出一个元素
			prev := start[len(start)-1]
			start = start[:len(start)-1]
			time[prev.id] += t - prev.t + 1
			// 如果start中还有元素
			if len(start) > 0 {
				// 开始时间重新赋值
				start[len(start)-1].t = t + 1
			}
		}
	}
	return time
}

// minLength
func minLength1(S string) int {
	s := []byte(S)
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if (s[i] == 'B' && top == 'A') || (s[i] == 'D' && top == 'C') {
				stack = stack[:len(stack)-1] // pop
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return len(stack)
}

// makeGood
func makeGood(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if top^32 == s[i] {
				stack = stack[:len(stack)-1] // pop
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}

func resultingString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			diff := abs(int(top) - int(s[i]))
			if diff == 1 || diff == 25 {
				stack = stack[:len(stack)-1] // pop
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}

func minDeletion(nums []int) int {
	var ans int
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		n := len(stack)
		// 如果栈大小为偶数，可以随意加入元素
		// 如果栈大小为奇数，那么加入的元素不能和栈顶相同
		if n > 0 && n%2 == 1 && stack[n-1] == nums[i] {
			stack = stack[:n-1]
		}
		stack = append(stack, nums[i])
	}
	if len(stack)%2 != 0 {
		ans++
	}
	return ans + len(nums) - len(stack)
}

func removeDuplicates(s string, k int) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		// 这里每次判断前k个字符是否相同 那么这个算法时间复杂度就是(o(kn))
		// 下边使用了计数栈来优化
		if len(stack) >= k && sameCharacter(string(stack[len(stack)-k:])) {
			stack = stack[:len(stack)-k]
		}
	}
	return string(stack)
}

func sameCharacter(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			return false
		}
	}
	return true
}

// 使用了计数栈来优化 判断相同字符的操作
func removeDuplicates1(s string, k int) string {
	type pair struct {
		ch  byte
		cnt int
	}
	stack := []pair{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1].ch == s[i] {
			stack[len(stack)-1].cnt++
			// 如果计数已经到达k，那么就需要pop
			if stack[len(stack)-1].cnt == k {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, pair{s[i], 1})
		}
	}
	// 拼接字符串
	var sb strings.Builder
	for _, p := range stack {
		for i := 0; i < p.cnt; i++ {
			sb.WriteByte(p.ch)
		}
	}
	return sb.String()
}

func maxDepth(s string) int {
	var ans int
	depth := 0
	for i := range s {
		if s[i] == '(' {
			depth++
			ans = max(ans, depth)
		} else if s[i] == ')' && depth > 0 {
			depth--
		}
	}
	return ans
}

func reverseParentheses(s string) string {
	stack := []string{""}
	// 每一层将结果返回给上一层
	for _, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, "")
		case ')':
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			reversed := reverseStr(top)
			stack[len(stack)-1] += reversed
		default:
			stack[len(stack)-1] += string(ch)
		}
	}
	return stack[0]
}

func reverseStr(s string) string {
	rune := []rune(s)
	for i, j := 0, len(rune); i < j; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}
	return string(rune)
}

func mergeAdjacent(nums []int) []int64 {
	stack := make([]int64, 0)
	for i := range nums {
		stack = append(stack, int64(nums[i]))
		for len(stack) > 1 && stack[len(stack)-1] == stack[len(stack)-2] {
			v := 2 * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, v)
		}
	}
	return stack
}

func robotWithString(s string) string {
	n := len(s)
	sufMin := make([]byte, n+1)
	sufMin[n] = math.MaxInt8
	// 记录后缀最小值
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}
	var sb strings.Builder
	stack := make([]byte, 0)
	// 遍历字符串s 当遇到了小于等于后缀最小值的字符时就开始弹出
	for i, ch := range s {
		stack = append(stack, byte(ch))
		for len(stack) > 0 && stack[len(stack)-1] <= sufMin[i+1] {
			sb.WriteByte(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return sb.String()
}

func scoreOfParentheses(s string) int {
	stack := []int{0} // 记录最终累加结果
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top == 0 {
				// 如果top==0那么其中就是表示为`()`
				stack[len(stack)-1] += 1
			} else {
				// 其中有嵌套
				stack[len(stack)-1] += top * 2
			}
		}
	}
	return stack[0]
}

func removeSubstring(s string, k int) string {
	type pair struct {
		ch  byte
		cnt int
	}
	var sb strings.Builder
	stack := make([]pair, 0)
	for i := 0; i < len(s); i++ {
		// 先添加
		if len(stack) > 0 && stack[len(stack)-1].ch == s[i] {
			stack[len(stack)-1].cnt++
		} else {
			stack = append(stack, pair{s[i], 1})
		}
		// 再消除
		n := len(stack)
		if n > 1 && stack[n-1].ch == ')' && stack[n-2].ch == '(' && stack[n-1].cnt >= k && stack[n-2].cnt >= k {
			stack[n-2].cnt -= k
			if stack[n-2].cnt == 0 {
				// pop 两个
				stack = stack[:n-2]
			} else {
				stack = stack[:n-1]
			}
		}
	}
	for i := range stack {
		sb.WriteString(strings.Repeat(string(stack[i].ch), stack[i].cnt))
	}
	return sb.String()
}

// 简单题不要想复杂
// 需要将最外层的括号去掉
// 标记层数即可
func removeOuterParentheses(s string) string {
	var sb strings.Builder
	d := 0
	// 层级标记 最外层不加入
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if d > 0 {
				sb.WriteByte(s[i])
			}
			d++
		} else {
			d--
			if d > 0 {
				sb.WriteByte(s[i])
			}
		}
	}
	return sb.String()
}

// 平衡字符串的性质
// 性质: 一定有一个 [ 匹配一个]
func minSwaps(s string) int {
	var ans int
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			c++
		} else if c > 0 {
			c-- // 匹配到了
		} else {
			// 没有匹配到对应的左括号
			ans++ // 交换
			c++   // 左括号有了匹配后面的
		}
	}
	return ans
}

func checkValidString(s string) bool {
	// 栈消除
	stack := make([]int, 0)
	star := make([]int, 0)
	for i, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, i)
		case '*':
			star = append(star, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else if len(star) > 0 {
				star = star[:len(star)-1]
			} else {
				return false
			}
		}
	}
	// 这里(* 的匹配不能够匹配数量,数量是结果 不是原因, 需要看( * 的位置分布
	// 最后匹配 ( 和 *
	i, j := len(stack)-1, len(star)-1
	for i >= 0 && j >= 0 {
		if stack[i] > star[j] {
			return false
		}
		i--
		j--
	}
	return i == -1
}

func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))
	depth := 0
	for i := range seq {
		if seq[i] == '(' {
			depth++
			ans[i] = depth % 2
		} else {
			ans[i] = depth % 2
			depth--
		}
	}
	return ans
}

// 最小
// 修改了平衡的概念
// `(` 对应了 `))`
func minInsertions(s string) int {
	ans := 0
	openCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			openCnt++
		} else {
			if (i+1) < len(s) && s[i+1] == ')' {
				i++
			} else {
				ans++
			}
			if openCnt > 0 {
				openCnt--
			} else {
				ans++
			}
		}
	}
	return ans + openCnt*2
}

func maximumGain(s string, x int, y int) int {
	mxstr, mnstr := "ab", "ba"
	mxScore, mnScore := x, y
	if x < y {
		mxstr, mnstr = mnstr, mxstr
		mxScore, mnScore = mnScore, mxScore
	}
	// 先按照分数最大的字符串进行匹配
	stack := make([]byte, 0)
	build := func(s string, substr string, score int) int {
		var ans int
		for i := 0; i < len(s); i++ {
			stack = append(stack, s[i])
			n := len(stack)
			if n >= 2 && string(stack[n-2:]) == substr {
				ans += score
				stack = stack[:n-2]
			}
		}
		return ans
	}
	s1 := build(s, mxstr, mxScore)
	newStr := string(stack)
	clear(stack)
	s2 := build(newStr, mnstr, mnScore)
	return s1 + s2
}

func maximumGain1(s string, x int, y int) int {
	mxstr, mnstr := "ab", "ba"
	mxScore, mnScore := x, y
	if x < y {
		mxstr, mnstr = mnstr, mxstr
		mxScore, mnScore = mnScore, mxScore
	}
	// 先按照分数最大的字符串进行匹配
	stack := make([]byte, 0)
	build := func(s string, substr string, score int) int {
		var ans int
		for i := 0; i < len(s); i++ {
			stack = append(stack, s[i])
			n := len(stack)
			if n >= 2 && string(stack[n-2:]) == substr {
				ans += score
				stack = stack[:n-2]
			}
		}
		return ans
	}
	s1 := build(s, mxstr, mxScore)
	newStr := string(stack)
	clear(stack)
	s2 := build(newStr, mnstr, mnScore)
	return s1 + s2
}

func minAddToMakeValid(s string) int {
	stack := make([]byte, 0)
	// 存储不能够匹配的括号
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack)
}

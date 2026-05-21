package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func smallestNumber(n int) int {
	// 统计n的二进制位数是多少
	// 返回 1<<m - 1 就可以了
	return 1<<bits.Len(uint(n)) - 1
}

// 非负整数 x 二进制的所有位都相同，意味着 x=0 或者 x 的二进制全为 1，即 x=2^k−1，其中 k 是非负整数。 问题转换为统计[1-n] 中有多少个2^k
func countMonobit(n int) int {
	return bits.Len(uint(n + 1))
}

func minChanges(n int, k int) int {
	if n&k != k {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	// 除以2的次数 bits.Len(uint(num))-1
	// 减去1的次数 bits.OnesCount(uint(num))
	return bits.Len(uint(num)) - 1 + bits.OnesCount(uint(num))
}

func findComplement(num int) int {
	// 补数
	return num ^ (1<<bits.Len(uint(num)) - 1)
}

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	return n ^ (1<<bits.Len(uint(n)) - 1)
}

func binaryGap(n int) int {
	var ans int
	num := bits.OnesCount(uint(n))
	if num <= 1 {
		return ans
	}
	last := -1
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				ans = max(ans, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return ans
}

func findKOr(nums []int, k int) int {
	var ans int
	// 如何快速统计nums中每个位的次数
	// 统计每个位的1的个数
	// 如果个数>=k 那么这个位为1 ans |= 1<<i
	for i := 0; i < 31; i++ {
		cnt := 0
		for _, num := range nums {
			if num>>i&1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			ans |= 1 << i
		}
	}
	return ans
}

// 数字的二进制表示中是否是01交替出现的
func hasAlternatingBits(n int) bool {
	x := n ^ (n >> 1)
	return x&(x+1) == 0
}

func hammingWeight(n int) int {
	return bits.OnesCount(uint(n))
}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	ans := make([]int, n)
	var maskA, maskB uint
	for i := range n {
		maskA |= 1 << A[i]
		maskB |= 1 << B[i]
		ans[i] = bits.OnesCount(maskA & maskB) // 计算交集 这个取值速度为O(1) 打表的
	}
	return ans
}

func insertBits(N int, M int, i int, j int) int {
	mask := ^((1 << (j + 1)) - 1) | ((1 << i) - 1)
	return N&mask | (M << i)
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// ans[i] = bits.OnesCount(uint(i))  // 调用api还是爽
		ans[i] = i&1 + ans[i>>1]
	}
	return ans
}

// evenOddBit 返回n的二进制形式中奇数下标和偶数下标中1的个数
func evenOddBit(n int) []int {
	even, odd := 0, 0
	evenBits := n & 0x55555555 // 5 => 0101
	oddBits := n & 0xAAAAAAAA  // A => 1010
	even = bits.OnesCount(uint(evenBits))
	odd = bits.OnesCount(uint(oddBits))
	ans := []int{even, odd}
	return ans
}

func findFinalValue(nums []int, original int) int {
	ans := original
	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}
	for has[ans] {
		ans *= 2
	}
	return ans
}

// 使用位运算优化上边的
func findFinalValue1(nums []int, original int) int {
	mask := 0
	for _, x := range nums {
		k := x / original
		if x%original == 0 && (k&(k-1)) == 0 {
			mask |= k
		}
	}
	// mask 这里需要找到最低位的0 因为这里orginnal的值是连续增长的，中间出现断开就不可以
	// mask 这里取反后就变成了找最低位置的1
	mask = ^mask
	// -mask 主要是运用取反 这里pc机器做的是补码操作 (各位取反然后+1)
	return original * (mask & -mask)
}

func validStrings(n int) []string {
	bytes := make([]byte, n)
	ans := make([]string, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n { // 长度够了
			ans = append(ans, string(bytes))
			return
		}
		bytes[i] = '1'
		dfs(i + 1)
		if i == 0 || bytes[i-1] == '1' {
			bytes[i] = '0'
			dfs(i + 1)
		}
	}
	dfs(0)
	return ans
}

func validStrings1(n int) []string {
	// 使用位运算
	mask := 1<<n - 1 //
	ans := make([]string, 0)
	// 遍历可能出现的所有元素 然后校验
	// 这里是逆向思维 判断一个数的二进制形式中是否包含相邻0 等价于 将其取反判断该数的二进制形式中是否包含相邻1
	// i & (i>>1) != 0 的情况下说明 i 的二进制形式中包含相邻的1
	for i := 0; i < (1 << n); i++ {
		// 满足条件
		if i&(i>>1) == 0 {
			// 满足条件但是需要取反
			ans = append(ans, fmt.Sprintf("%0*b", n, i^mask))
		}
	}
	return ans
}

func minimumFlips(num int) int {
	n := uint(num)
	reverse := bits.Reverse(n) >> uint(bits.LeadingZeros(n))
	return bits.OnesCount(uint(reverse ^ n))
}

func sortByReflection(nums []int) []int {
	// 如果提前计算 那么需要计算数组的交换也需要做
	reverse := func(n int) int {
		// 翻转之后右移前导0的位数可以获取有效翻转位
		return int(bits.Reverse(uint(n)) >> bits.LeadingZeros(uint(n)))
	}
	// sort.Slice(nums, func(i, j int) bool {
	// 	ri, rj := reverse(nums[i]), reverse(nums[j])
	// 	if ri != rj {
	// 		return ri < rj
	// 	}
	// 	return nums[i] < nums[j]
	// })
	slices.SortFunc(nums, func(i, j int) int {
		ri, rj := reverse(i), reverse(j)
		if ri != rj {
			return ri - rj
		}
		return i - j
	})
	return nums
}

// TODO: 系统学习BFS之后才开始这个
// 结合位运算 查看这种解法
func minSplitMerge(nums1 []int, nums2 []int) int {
	return 1
}

func decode(encoded []int, first int) []int {
	n := len(encoded)
	ans := make([]int, n+1)
	ans[0] = first
	for i := range encoded {
		ans[i+1] = ans[i] ^ encoded[i]
	}
	return ans
}

func findArray(pref []int) []int {
	n := len(pref)
	ans := make([]int, n)
	ans[0] = pref[0]
	for i := 1; i < n; i++ {
		ans[i] = pref[i-1] ^ pref[i]
	}
	return ans
}

func longestSubsequence(nums []int) int {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	if xor != 0 {
		return len(nums)
	}
	// 异或和为0
	// 1.nums全是0 2.nums中元素出现的次数>2
	for _, x := range nums {
		if x != 0 {
			return len(nums) - 1
		}
	}
	return 0
}

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, x := range derived {
		xor ^= x
	}
	return xor == 0
}

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	mask := 1<<maximumBit - 1
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		nums[i] ^= nums[i-1]
	}
	for i := range nums {
		ans[n-i-1] = nums[i] ^ mask
	}
	return ans
}

// leetcode 1525
func minOperations2(nums []int, k int) int {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	return bits.OnesCount(uint(k ^ xor))
}

func countTriplets(arr []int) int {
	var ans int
	n := len(arr)
	// 计算前缀异或和
	prefix := make([]int, n+1)
	for i := range arr {
		prefix[i+1] = prefix[i] ^ arr[i]
	}
	// 对于i,j,k 来说有以下
	// preifx[i] ^ prefix[j] == prefix[k+1] ^ prefix[j]
	// 简化得到 preifx[i] == prefix[k+1]
	// 枚举k 如果有前缀和相同的坐标为i 那么会产生 (k-i)个三元组
	cnt := map[int]int{}   // 记录相同前缀出现的次数  例如在当前k+1的位置中的前缀和在之前出现过3次 那么 cnt[] = 3
	total := map[int]int{} // 记录相同前缀和的 索引和 例如 在i=1, i=3, i=5 中前缀和和当前k+1位置的前缀和一样那么total=9
	// 知道这些值就可以算出 增量的三元组 k-i1+k-i2+k-i3 => k*cnt-sum(idx)
	for k := range n {
		if c, ok := cnt[prefix[k+1]]; ok {
			ans += c*k - total[prefix[k+1]] // 这里主要是累加(k-i) 出现的所有的和
		}
		cnt[prefix[k]]++
		total[prefix[k]] += k
	}
	return ans
}

func minimizeXor(num1 int, num2 int) int {
	// 看num2和num1的差值
	// 如果相等 返回num1
	// 如果 num2 > num1 那么将返回num1 | diff
	// 如果 num1 > num2
	diff := bits.OnesCount(uint(num2)) - bits.OnesCount(uint(num1))
	if diff == 0 {
		return num1
	}
	if diff > 0 {
		// 找到num1最低 diff位置 将最低位的0变成1
		for i := 0; i < 32 && diff > 0; i++ {
			if num1>>i&1 == 0 {
				num1 |= 1 << i
				diff--
			}
		}
		return num1
	}
	if diff < 0 {
		// 将最低位的1变成0
		diff = -diff
		cnt := bits.Len(uint(num1))
		for i := 0; i < cnt && diff > 0; i++ {
			if num1>>i&1 == 1 {
				num1 ^= 1 << i
				diff--
			}
		}
	}
	return num1
}

func minimizeXor1(num1 int, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))
	for ; c1 > c2; c2++ {
		// 将最低位的1变成0
		num1 &= num1 - 1
	}
	for ; c2 > c1; c1++ {
		// 将最低位的0变成1
		num1 |= num1 + 1
	}
	return num1
}

func xorBeauty(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}

// // maximumXOR 计算经过任意次「将某个数的某个 1 变成 0」操作后，整个数组能得到的最大异或和。
//
// 核心观察：
// 1. 操作只能把 二进制位中的 1 改为 0，不能把 0 改为 1，所以 0 的位永远无法变成 1。
// 2. 异或和等价于每个二进制位上 1 的个数的奇偶性：
//   - 某位上有奇数个 1 → 异或和该位为 1
//   - 某位上有偶数个 1 → 异或和该位为 0
//     3. 对于任意一个在 nums 中出现过 1 的二进制位，我们总能通过`操作`减少一些数的该位 1（改为 0），
//     来自由控制最后该位 1 的个数是奇数还是偶数，从而让异或结果的该位为 1。
//     例如：如果当前该位有偶数个 1，只需将其中任意一个 1 改为 0，就变成了奇数个 1。
//     4. 如果某一位在所有数中都没有出现过 1，那最终异或结果该位只能是 0。
//
// 因此，最大异或和就是所有数的 按位或（OR），即只要某个位在任意一个数中出现过 1，
// 它就能在最终结果中为 1。
func maximumXOR(nums []int) int {
	res := 0
	for _, num := range nums {
		res |= num
	}
	return res
}

// 检查最后一个元素是不是全是1
func hasTrailingZeros(nums []int) bool {
	cnt := 0
	for _, num := range nums {
		if num&1 == 0 {
			cnt++
		}
	}
	return cnt >= 2
}

func minFlips(a int, b int, c int) int {
	or := a | b
	and := a & b
	diff := or ^ c                                // 哪些位置需要变动
	double := bits.OnesCount(uint(diff & and))    // 哪些位置需要更改两次
	single := bits.OnesCount(uint(diff)) - double // 剩下的需要变动1次
	return double*2 + single
}

// 最大元素的连续长度
func longestSubarray(nums []int) int {
	ans := 1
	cnt := 1
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == mx {
			cnt++
			ans = max(ans, cnt)
		} else if nums[i] > mx {
			cnt = 1
			ans = 1
			mx = nums[i]
		} else {
			cnt = 0
		}
	}
	return ans
}

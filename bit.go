package main

import (
	"math/bits"
)

func smallestNumber(n int) int {
	// 统计n的二进制位数是多少
	// 返回 1<<m - 1 就可以了
	return 1<<bits.Len(uint(n)) - 1
}

// 非负整数 x 二进制的所有位都相同，意味着 x=0 或者 x 的二进制全为 1，即 x=2^k−1，其中 k 是非负整数。
// 问题转换为统计[1-n] 中有多少个2^k
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
	return nil
}

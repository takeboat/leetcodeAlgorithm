package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	pre := make([]int, len(nums)+1)
	for i, num := range nums {
		pre[i+1] = pre[i] ^ num
	}
	fmt.Println(pre)
	nums1 := []int{3, 4, 5}
	ans := 0
	for _, num := range nums1 {
		ans ^= num
	}
	fmt.Println("[3, 4, 5]: ", ans)
	fmt.Println("pre[5]^pre[2]", pre[5]^pre[2])
	fmt.Println("this is", pre[5]^pre[2] == ans)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := PrefixSum(nums)
	fmt.Println(nums)
	fmt.Println(ans)
	// [1 2 3 4 5 6 7 8 9 10]
	// [0 1 3 6 10 15 21 28 36 45 55]
	// 计算索引 2-5 之间的和 3+4+5+6 = 18
	// ans[6]-ans[2] = 21-3 = 18
	// so 计算nums 索引在 i--j 之间的和
	// 就是 ans[j+1] - ans[i]
}

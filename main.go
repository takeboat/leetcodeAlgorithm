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
}

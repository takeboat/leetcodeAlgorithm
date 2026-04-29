package main

import "fmt"

func main() {
	grid := [][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}
	ans := differenceOfDistinctValues1(grid)
	fmt.Println(ans)
}

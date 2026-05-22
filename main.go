package main

import "fmt"

func main() {
	rc := ConstructorRecentCounter()
	fmt.Println(rc.Ping(642))
	fmt.Println(rc.Ping(1849))
	fmt.Println(rc.Ping(4921))
	fmt.Println(rc.Ping(5936))
	fmt.Println(rc.Ping(5957))
}

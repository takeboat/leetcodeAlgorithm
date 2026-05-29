package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	val := uint16(1616) // 06 50
	bytes := (*[2]byte)(unsafe.Pointer(&val))
	fmt.Printf("%02x %02x", bytes[0], bytes[1])
	fmt.Println()
	tmp := make([]byte, 2)
	binary.LittleEndian.PutUint16(tmp, val)
	for _, v := range tmp {
		fmt.Printf("%02x ", v)
	}
	fmt.Println()
	// 50 06
}

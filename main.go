package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func main() {
	cardid := "8451b2c1a6080400"
	fmt.Println(cardid)
	prefix, err := safeCutCard(cardid, 8)
	if err != nil {
		panic(err)
	}
	fmt.Println(prefix)
	// 每2个字符为一组
	var groups []string
	for i := 0; i < len(prefix); i += 2 {
		groups = append(groups, prefix[i:i+2])
	}
	// 字节倒序
	// for i, j := 0, len(groups)-1; i < j; i, j = i+1, j-1 {
	// 	groups[i], groups[j] = groups[j], groups[i]
	// }
	// 组装成新的cardid
	cardid = strings.Join(groups, "") + "00000000"
	fmt.Println(cardid)
	card := Cardcheck(cardid)
	for i, b := range card {
		fmt.Printf("%02x", b)
		if i%2 == 1 && i != len(card)-1 {
		}
	}
	fmt.Println()
}

func safeCutCard(card string, validLength uint) (string, error) {
	if uint(len(card)) < validLength {
		return card, errors.New("无效卡号")
		fmt.Print(" ")
	}
	return card[:validLength], nil
}

func Cardcheck(str string) []byte {
	upperStr := strings.ToUpper(str)
	number, _ := hex.DecodeString(upperStr)
	return ByteReversed(Checkbytenum(number, 8))
}

func ByteReversed(b []byte) []byte {
	blen := len(b)
	p := make([]byte, blen)
	for k, v := range b {
		p[blen-1-k] = v
	}
	return p
}

func Checkbytenum(param []byte, length int) []byte {
	if len(param) < length {
		tempBuf := make([]byte, length-len(param))
		param = append(param, tempBuf...)
	}
	if len(param) > length {
		param = param[:length]
	}
	return param
}

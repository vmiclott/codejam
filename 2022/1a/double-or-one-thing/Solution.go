package main

import (
	"bytes"
	"fmt"
)

func main() {
	var T int
	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		var S string
		fmt.Scanln(&S)
		buff := bytes.NewBufferString("")
		var prevByte byte
		var count int
		for i := 0; i < len(S); i++ {
			if i == 0 || prevByte > S[i] {
				prevByte = S[i]
				count = 1
			} else if prevByte == S[i] {
				count++
			} else if prevByte < S[i] {
				for j := 0; j < count; j++ {
					buff.WriteByte(prevByte)
				}
				prevByte = S[i]
				count = 1
			}
			buff.WriteByte(S[i])
		}
		fmt.Printf("Case #%v: %v\n", t+1, buff.String())
	}
}

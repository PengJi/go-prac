package main

import (
	"fmt"
)

func main() {
	pool := []byte {
		0x34, 0x09, 0x29, 0x34, 0x32, 0x33, 0x36, 0x39,
		0x33, 0x30, 0x35, 0x31, 0x38, 0x31, 0x39, 0x3a,
		0x30, 0x24, 0x31, 0x3b, 0x33, 0x28, 0x35, 0x37,
		0x26, 0x14, 0x37, 0x34, 0x11, 0x39, 0x25, 0x38,
	}

	key := "5506372864345658958606"
	for i := 0; i < len(key)/2; i++ {
		fmt.Print(string(pool[key[i] + key[len(key) - 1 - i] - 96]))
	}
}
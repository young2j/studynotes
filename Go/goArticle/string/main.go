package main

import (
	"fmt"
)

func main() {
	s := "go语言"
	for _, r := range s {
		fmt.Printf("rune: %v  unicode: %#U\n", r, r)
	}
}
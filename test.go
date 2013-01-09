package main

import (
	"fmt"
	"strings"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%s\n", strings.Repeat("A", i))
	}
}

func main() {
	test()
}

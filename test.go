package main

import (
	"bufio"
	"fmt"
	"regexp"
)

func main() {
	c := make(chan []byte)
	for {
		go func() {
			a = append(a, i)
			c <- a
		}()
		a = <-c
		fmt.Println(a)
	}
}

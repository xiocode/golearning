package main

import (
	"fmt"
	// "strconv"
)

func test() {
	// var arr [4]int
	a := [3]int{1, 2, 3}
	// arr[0] = 1
	// arr[1] = 1
	// arr[2] = 1
	// arr[3] = 1
	slice1 := a[:]
	slice2 := slice1[0:2]               //The first 2 elements
	fmt.Printf("slice1: %d \n", slice1) //'xxx: [1,2,3,4]'
	fmt.Printf("slice2: %d \n", slice2)
	// slice2 := arr[:3]
	// println(len(a))
	// println(len(a))
	fmt.Printf("slice1 len: %d \n", len(slice1))
	fmt.Printf("slice1 cap: %d \n", cap(slice1))
	// println(len(arr))
	// println(cap(arr))
	fmt.Printf("slice2 len: %d \n", len(slice2))
	fmt.Printf("slice2 cap: %d \n", cap(slice2))

}

func main() {
	test()
}

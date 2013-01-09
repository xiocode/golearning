package main

import (
	"fmt"
)

func test() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	value, ok := monthdays["Jan"]
	if ok {
		fmt.Printf("value %d, is_ok %t \n", value, ok)
	} else {
		fmt.Printf("value %d, is_ok %t \n", value, ok)
	}
	delete(monthdays, "Jan")
	value1, ok1 := monthdays["Jan"] //:= declare new variable and variable should be new
	if ok1 {
		fmt.Printf("value %d, is_ok %t \n", value1, ok1)
	} else {
		fmt.Printf("value %d, is_ok %t \n", value1, ok1)
	}

}

func main() {
	test()
}

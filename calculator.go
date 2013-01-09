package main

/**
	a simple calcuator
**/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var stack = new(Stack)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) push(v int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = v
	s.i++
}

func (s *Stack) pop() (ret int) {
	s.i--
	if s.i < 0 {
		s.i = 0
		return
	}
	ret = s.data[s.i]
	return
}

func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token += string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				stack.push(r)
				token = ""
			case c == '+':
				p1 := stack.pop()
				q1 := stack.pop()
				fmt.Printf("%d + %d=%d\n", p1, q1, p1+q1)
			case c == '*':
				p2 := stack.pop()
				q2 := stack.pop()
				fmt.Printf("%d * %d=%d\n", p2, q2, p2*q2)
			case c == '-':
				p3 := stack.pop()
				q3 := stack.pop()
				fmt.Printf("%d - %d=%d\n", p3, q3, p3*q3)
			case c == 'q':
				return
			default:
				//error

			}
		}
	}
}

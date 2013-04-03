package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"time"
	//"reflect"
)

var (
	count   = 10000000
	segSize = 100000000
	seg     = math.MaxInt32 / segSize
	retChan = make(chan []int, seg)
)

func Reduce(ch chan int) {
	//fmt.Println("Enter Reduce")
	var s []int
	for v := range ch {
		s = append(s, v)
	}
	if len(s) == 0 {
		//fmt.Println("Leave Empty Reduce")
		return
	}
	sort.Ints(s)
	retChan <- s
	//fmt.Println("Leave Reduce: ", s[0]/segSize)
}

func start_reduce() []chan int {
	chans := make([]chan int, seg)
	for i := 0; i < seg; i++ {
		chans[i] = make(chan int, 10)
		go Reduce(chans[i])
	}
	return chans
}

func Map(s []int, chans []chan int, ch chan bool) {
	//fmt.Println("Enter Map, len(chans)=", len(chans), ", len(s)=", len(s))
	for i := 0; i < len(s); i++ {
		sg := s[i] / segSize
		//fmt.Printf("START: chans[%v] <- s[%v](%v)\n", sg, i, s[i])
		chans[sg] <- s[i]
		//fmt.Printf("END: chans[%v] <- s[%v](%v)\n", sg, i, s[i])
	}
	ch <- true
	//fmt.Println("Leave Map")
}

func start_map(slice []int, chans []chan int, finChan chan bool) {
	var s, e int
	gap := count / seg
	for sg := 0; sg < seg; sg++ {
		s = gap * sg
		e = gap * (sg + 1)
		if s >= count {
			break
		}
		if e > count {
			e = count
		}
		s = 0
		e = 22
		//fmt.Println("sg:", sg, "s:", s, "e:", e, "len(slice):", len(slice), "cap(slice):", cap(slice))
		go Map(slice[s:e], chans, finChan)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	var s = make([]int, count)
	for i := 0; i < count; i++ {
		s[i] = rand.Int()
	}

	st := time.Now()

	reduce_chans := start_reduce()

	finChan := make(chan bool, seg)
	start_map(s, reduce_chans, finChan)
	for i := 0; i < seg; i++ {
		<-finChan
	}

	result := make([][]int, seg)
	for r := range retChan {
		if len(r) == 0 {
			continue
		}
		sg := r[0] / segSize
		result[sg] = r
	}

	dur := time.Since(st)
	fmt.Println(dur)
}

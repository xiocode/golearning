/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          split.go
 * Description:   Split a large log file
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

var pattern *regexp.Regexp
var workers map[string]chan []byte

func init() {
	var err error
	pattern, err = regexp.Compile("http://([^/]+?)/")
	if err != nil {
		log.Fatal(err)
	}
	workers = make(map[string]chan []byte)
}

func getSiteName(line []byte) string {
	result := pattern.FindSubmatch(line)
	return string(result[1])
}

func split(path string) {
	done := make(chan bool)
	stop := make(chan bool)
	out := make(chan []byte, 1024)
	defer close(out)
	defer close(stop)
	defer close(done)
	go readLines(path, out, done)
	go writeLines(out, done, stop)
	<-stop
}

func readLines(path string, out chan<- []byte, done chan<- bool) {
	var (
		line   []byte
		prefix bool
		err    error
		file   *os.File
	)
	file, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("DEBUG")
	reader := bufio.NewReaderSize(file, 1024*10)
	for {
		line, prefix, err = reader.ReadLine()
		if err != nil {
			done <- true
			break
		}
		if !prefix {
			out <- line // lines channel
		}
	}
	if err == io.EOF {
		fmt.Println(err)
	}
}

func writeLines(in <-chan []byte, done <-chan bool, stop chan bool) {
	for {
		select {
		case line := <-in:
			siteName := getSiteName(line)
			fmt.Println(siteName)
			if _, ok := workers[siteName]; ok {
				workers[siteName] <- line
			} else {
				workers[siteName] = make(chan []byte)
				workers[siteName] <- line
				go writer(siteName, stop)
			}
		case <-done:
			for i := len(workers); i >= 0; i-- {
				fmt.Println(i)
				stop <- true
			}
			break
		}
	}
}

func writer(siteName string, stop <-chan bool) {
	file, err := os.OpenFile(siteName+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriterSize(file, 1024*10)
	defer writer.Flush()
	for {
		select {
		case <-stop:
			break
		default:
			line := <-workers[siteName]
			writer.Write(line)
			writer.Write([]byte("\n"))
		}

	}

}

func main() {
	start := time.Now().Second()
	split("test.log")
	fmt.Printf("CostTime: %v", time.Now().Second()-start)
}

//TODO site num groutines

package main

import (
	"bufio"
	"fmt"
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
	result := pattern.FindStringSubmatch(string(line))
	if len(result) < 1 { // <--- check error
		fmt.Println(string(line))
	}
	return result[1]
}

func writeLines(input <-chan []byte, done <-chan bool, finished chan bool, bang chan bool) {
	for {
		select {
		case line := <-input:
			siteName := getSiteName(line)
			if _, ok := workers[siteName]; !ok {
				workers[siteName] = make(chan []byte, 10)
				go writer(siteName, finished, bang)
			}
			workers[siteName] <- line
		case <-done:
			for i := len(workers); i > 0; i-- {
				finished <- true
			}
		}

	}
}

func writer(siteName string, finished chan bool, bang chan bool) {
	file, err := os.OpenFile(siteName+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriterSize(file, 1024*10)
	defer writer.Flush()
	stop := false
	for !stop {
		select {
		case line := <-workers[siteName]:
			writer.Write(line)
			writer.Write([]byte("\n"))
		case <-finished:
			stop = true
			fmt.Println("writer finished!")
			continue
		}
	}
	bang <- true
}

func readLines(path string, output chan<- []byte, done chan<- bool) {

	file, err := os.Open(path)
	if err != nil {
		done <- true
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*10)
	for {
		line, prefix, err := reader.ReadLine()
		if err != nil {
			done <- true
			fmt.Println(err)
			break
		}
		if !prefix {
			output <- line // lines channel
		}
	}
}

func main() {
	start := time.Now().Second()
	input := make(chan []byte, 2)
	done := make(chan bool)
	finished := make(chan bool, 10)
	bang := make(chan bool, 10)
	defer func() { // close
		close(input)
		close(done)
		close(finished)
		close(bang)
	}()

	go readLines("test.log", input, done)
	go writeLines(input, done, finished, bang)

	stop := false
	for !stop {
		_, stop = <-bang
	}
	fmt.Printf("CostTime: %v ", time.Now().Second()-start)
}

/**
for {
                data, ok := <-ch
                if ok {
                        // do something with data
                } else {
                        // no data received, do something else
                }
        }
**/

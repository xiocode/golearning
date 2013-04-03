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

func init() {
	var err error
	pattern, err = regexp.Compile("http://([^/]+?)/")
	if err != nil {
		log.Fatal(err)
	}
}

func getSiteName(line []byte) string {
	result := pattern.FindSubmatch(line)
	return string(result[1])
}

func readLine(path string) (err error) {
	var (
		file   *os.File
		line   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1024*1024*100)
	for {
		if line, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		if !prefix {
			logWriter(line)
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func logWriter(line []byte) (err error) {
	var (
		file *os.File
	)

	siteName := getSiteName(line)

	if file, err = os.OpenFile(siteName+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.Write(line)
	writer.Write([]byte("\n"))
	// return
	return nil
}

func main() {
	start := time.Now().Second()
	readLine("test.log")
	fmt.Printf("CostTime: %v", time.Now().Second()-start)
}

//TODO site num groutines

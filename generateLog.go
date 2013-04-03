package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	FOMART = `%s - - [%s] "%s %s" %s %d "-" "Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:8.0) Gecko/20100101 Firefox/8.0" 0`
	MAXINT = 1 << 26
)

var (
	domainlist []string = []string{"www.asta.com", "www.golang.org", "www.mygolang.com", "www.gocn.im", "www.sina.com"}
	urllist    []string = []string{"1.css", "2.js", "3.html", "4.jpg", "5.php", "6.html", "login"}
	errnums    []string = []string{"200", "302", "404", "403", "501"}
	seedrand   *rand.Rand
)

func init() {
	seedrand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	f, err := os.Create("test.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var i int
	var method string
	for {
		if i%4 == 0 {
			method = "POST"
		} else {
			method = "GET"
		}
		f.WriteString(fmt.Sprintf(FOMART, getIP(), getTime(i), method, getURL(), errnums[seedrand.Intn(len(errnums))]) + "\n")
		i++
		if i > MAXINT {
			break
		}
	}

}

func getIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", seedrand.Intn(255), seedrand.Intn(255), seedrand.Intn(255), seedrand.Intn(255))
}

func getTime(i int) string {
	t := time.Now()
	t.Add(time.Duration(i))
	return Date(t, "d/M/Y:H:i:s -0700")
}

func getURL() string {
	return "http://" + domainlist[seedrand.Intn(len(domainlist))] + "/" + urllist[seedrand.Intn(len(urllist))]
}

// Date takes a PHP like date func to Go's time fomate
func Date(t time.Time, format string) (datestring string) {
	patterns := []string{
		// year
		"Y", "2006", // A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
		"y", "06", //A two digit representation of a year	Examples: 99 or 03

		// month
		"m", "01", // Numeric representation of a month, with leading zeros	01 through 12
		"n", "1", // Numeric representation of a month, without leading zeros	1 through 12
		"M", "Jan", // A short textual representation of a month, three letters	Jan through Dec
		"F", "January", // A full textual representation of a month, such as January or March	January through December

		// day
		"d", "02", // Day of the month, 2 digits with leading zeros	01 to 31
		"j", "2", // Day of the month without leading zeros	1 to 31

		// week
		"D", "Mon", // A textual representation of a day, three letters	Mon through Sun
		"l", "Monday", // A full textual representation of the day of the week	Sunday through Saturday

		// time
		"g", "3", // 12-hour format of an hour without leading zeros	1 through 12
		"G", "15", // 24-hour format of an hour without leading zeros	0 through 23
		"h", "03", // 12-hour format of an hour with leading zeros	01 through 12
		"H", "15", // 24-hour format of an hour with leading zeros	00 through 23

		"a", "pm", // Lowercase Ante meridiem and Post meridiem	am or pm
		"A", "PM", // Uppercase Ante meridiem and Post meridiem	AM or PM

		"i", "04", // Minutes with leading zeros	00 to 59
		"s", "05", // Seconds, with leading zeros	00 through 59
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)
	datestring = t.Format(format)
	return
}

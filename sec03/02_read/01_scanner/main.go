package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I felt so good like anything was possible\nI hit cruise control and rubbed my eyes\n"

	myScanner := bufio.NewScanner(strings.NewReader(s)) // default is ScanLines
	myScanner.Split(bufio.ScanWords)

	for myScanner.Scan() {
		line := myScanner.Text()
		fmt.Println(line)
	}
}

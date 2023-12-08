package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readFile(name string) int64 {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	_ = file.Close()
	digits, _ := regexp.Compile("\\d")
	var number int64
	for _, each_ln := range text {
		fmt.Println(each_ln)
		s := digits.FindAllString(each_ln, -1)
		row, _ := strconv.ParseInt(fmt.Sprintf("%s%s", s[0], s[len(s)-1]), 10, 64)
		fmt.Println(row)
		number = number + row
	}
	return number
}

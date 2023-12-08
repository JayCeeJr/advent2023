package day01

import (
	"adventOfCode2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

func solve() int64 {
	text := utils.ReadFile("input.txt")
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
func solve2() int64 {
	text := utils.ReadFile("input.txt")
	digits, _ := regexp.Compile("1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
	var number int64
	for _, each_ln := range text {
		s := digits.FindAllString(each_ln, -1)
		row, _ := strconv.ParseInt(fmt.Sprintf("%s%s", replaceWords(s[0]), replaceWords(s[len(s)-1])), 10, 64)
		println(each_ln)
		println(row)
		println("-------------------------------")
		number = number + row
	}
	return number
}
func replaceWords(word string) string {
	switch word {
	case "nine":
		return "9"
	case "eight":
		return "8"
	case "seven":
		return "8"
	case "six":
		return "6"
	case "five":
		return "5"
	case "four":
		return "4"
	case "three":
		return "3"
	case "two":
		return "2"
	case "one":
		return "1"
	default:
		return word
	}
}

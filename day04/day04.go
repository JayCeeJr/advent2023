package day04

import (
	"adventOfCode2023/utils"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Mine  []int64
	Win   []int64
	Score int64
}

func solve() int64 {
	lineParser, _ := regexp.Compile("Card\\s+([0-9]+):\\s+([0-9 ]+) \\|\\s+([0-9 ]+)")
	rawData := utils.ReadFile("input.txt")
	data := make(map[string]Card)
	for _, datum := range rawData {
		p := lineParser.FindStringSubmatch(datum)
		data[p[1]] = Card{
			Mine: parseNumbers(p[2]),
			Win:  parseNumbers(p[3]),
		}
	}
	var solution int64
	return solution
}

func parseNumbers(n string) []int64 {
	var res []int64
	for _, v := range strings.Split(n, " ") {
		if strings.Trim(v, " ") == "" {
			continue
		}
		val, _ := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
		res = append(res, val)
	}
	return res
}

func (c *Card) scoreCard() int64 {
	winners
	for i, v := range c.Mine {

	}
}

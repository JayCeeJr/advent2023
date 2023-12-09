package day04

import (
	"adventOfCode2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Copies  int64
	Matched int64
	Mine    []int64
	Win     []int64
	Score   int64
}

func solve() int64 {
	lineParser, _ := regexp.Compile("Card\\s+([0-9]+):\\s+([0-9 ]+) \\|\\s+([0-9 ]+)")
	rawData := utils.ReadFile("input.txt")
	data := make(map[string]Card)
	var solution int64
	for _, datum := range rawData {
		p := lineParser.FindStringSubmatch(datum)
		newCard := Card{
			Mine: parseNumbers(p[2]),
			Win:  parseNumbers(p[3]),
		}
		newCard.scoreCard()
		data[p[1]] = newCard
		solution = solution + newCard.Score
	}
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
	for _, v := range c.Mine {
		if slices.Contains(c.Win, v) {
			c.Matched = c.Matched + 1
			if c.Score == 0 {
				c.Score = 1
			} else {
				c.Score = c.Score * 2
			}
		}
	}
	return c.Score
}

func (c *Card) IncrementCopies(n int64) {
	c.Copies = c.Copies + n
}

func solve2() int64 {
	lineParser, _ := regexp.Compile("Card\\s+([0-9]+):\\s+([0-9 ]+) \\|\\s+([0-9 ]+)")
	rawData := utils.ReadFile("input.txt")
	data := make(map[string]*Card)
	var solution int64
	for _, datum := range rawData {
		p := lineParser.FindStringSubmatch(datum)
		data[p[1]] = &Card{
			Copies: 1,
			Mine:   parseNumbers(p[2]),
			Win:    parseNumbers(p[3]),
		}
		data[p[1]].scoreCard()
	}
	for i := int64(1); i <= int64(len(rawData)); i++ {
		for m := int64(1); m <= data[fmt.Sprintf("%v", i)].Matched; m++ {
			if m+1 <= int64(len(rawData)) {
				data[fmt.Sprintf("%v", m+i)].IncrementCopies(data[fmt.Sprintf("%v", i)].Copies)
			}
		}

	}
	for _, card := range data {
		solution = solution + card.Copies
	}
	return solution
}

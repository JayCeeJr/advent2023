package day03

import (
	"adventOfCode2023/utils"
	"regexp"
	"strings"
)

func solve() int64 {
	symbols := map[int]map[int]bool{}
	symbol, _ := regexp.Compile("[^0-9.]")
	rows := utils.ReadFile("input.txt")
	for i := range rows {
		symbols[i] = map[int]bool{}
	}
	for l, row := range rows {
		for c, val := range strings.Split(row, "") {
			if symbol.MatchString(val) {
				symbols[l][c] = true
			}
		}
	}
	println(symbols)
	return 0
}

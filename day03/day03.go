package day03

import (
	"adventOfCode2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func solve() int64 {
	var found []int64
	symbol, _ := regexp.Compile("[^0-9.]")
	number, _ := regexp.Compile("\\d")
	rows := utils.ReadFile("input.txt")
	var symbols [140][140]bool
	var numbers [140][140]string
	var numbersParsed [140][140]int64
	for l, row := range rows {
		for c, val := range strings.Split(row, "") {
			symbols[l][c] = symbol.MatchString(val)
			if number.MatchString(val) {
				numbers[l][c] = val
			} else {
				numbers[l][c] = ""
			}
		}
	}
	for l, row := range rows {
		for c := range strings.Split(row, "") {
			// join numbers
			if numbers[l][c] != "" {
				num := numbers[l][c]
				// forward
				if c+1 < 140 {
					if numbers[l][c+1] != "" {
						num = num + numbers[l][c+1]
						if c+2 < 140 {
							if numbers[l][c+2] != "" {
								num = num + numbers[l][c+2]
							}
						}
					}

				}
				// backward
				if c-1 >= 0 && numbers[l][c-1] != "" {
					num = numbers[l][c-1] + num
					if c-2 >= 0 && numbers[l][c-2] != "" {
						num = numbers[l][c-2] + num
					}
				}
				numbersParsed[l][c], _ = strconv.ParseInt(num, 10, 64)
			}
		}
	}
	for r, bools := range symbols {
		for c, b := range bools {
			if b {
				var _found []int64
				// this row
				if numbersParsed[r][c-1] != 0 {
					_found = append(_found, numbersParsed[r][c-1])
				}
				if numbersParsed[r][c+1] != 0 {
					_found = append(_found, numbersParsed[r][c+1])
				}
				// last row
				if numbersParsed[r-1][c] != 0 {
					_found = append(_found, numbersParsed[r-1][c])
				}
				if numbersParsed[r-1][c-1] != 0 {
					if !slices.Contains(_found, numbersParsed[r-1][c-1]) {
						_found = append(_found, numbersParsed[r-1][c-1])
					}
				}
				if numbersParsed[r-1][c+1] != 0 {
					if !slices.Contains(_found, numbersParsed[r-1][c+1]) {
						_found = append(_found, numbersParsed[r-1][c+1])
					}
				}
				// next row
				if numbersParsed[r+1][c] != 0 {
					_found = append(_found, numbersParsed[r+1][c])
				}
				if numbersParsed[r+1][c-1] != 0 {
					if !slices.Contains(_found, numbersParsed[r+1][c-1]) {
						_found = append(_found, numbersParsed[r+1][c-1])
					}
				}
				if numbersParsed[r+1][c+1] != 0 {
					if !slices.Contains(_found, numbersParsed[r+1][c+1]) {
						_found = append(_found, numbersParsed[r+1][c+1])
					}
				}
				found = append(found, _found...)
			}
		}
	}
	var res int64
	for _, r := range found {
		res = res + r
	}
	fmt.Println(res)
	return res
}

func solve2() int64 {
	var found []int64
	symbol, _ := regexp.Compile("\\*")
	number, _ := regexp.Compile("\\d")
	rows := utils.ReadFile("input.txt")
	var symbols [140][140]bool
	var numbers [140][140]string
	var numbersParsed [140][140]int64
	for l, row := range rows {
		for c, val := range strings.Split(row, "") {
			symbols[l][c] = symbol.MatchString(val)
			if number.MatchString(val) {
				numbers[l][c] = val
			} else {
				numbers[l][c] = ""
			}
		}
	}
	for l, row := range rows {
		for c := range strings.Split(row, "") {
			// join numbers
			if numbers[l][c] != "" {
				num := numbers[l][c]
				// forward
				if c+1 < 140 {
					if numbers[l][c+1] != "" {
						num = num + numbers[l][c+1]
						if c+2 < 140 {
							if numbers[l][c+2] != "" {
								num = num + numbers[l][c+2]
							}
						}
					}

				}
				// backward
				if c-1 >= 0 && numbers[l][c-1] != "" {
					num = numbers[l][c-1] + num
					if c-2 >= 0 && numbers[l][c-2] != "" {
						num = numbers[l][c-2] + num
					}
				}
				numbersParsed[l][c], _ = strconv.ParseInt(num, 10, 64)
			}
		}
	}
	for r, bools := range symbols {
		for c, b := range bools {
			if b {
				var _found []int64
				// this row
				if numbersParsed[r][c-1] != 0 {
					_found = append(_found, numbersParsed[r][c-1])
				}
				if numbersParsed[r][c+1] != 0 {
					_found = append(_found, numbersParsed[r][c+1])
				}
				// last row
				if numbersParsed[r-1][c] != 0 {
					_found = append(_found, numbersParsed[r-1][c])
				}
				if numbersParsed[r-1][c-1] != 0 {
					if !slices.Contains(_found, numbersParsed[r-1][c-1]) {
						_found = append(_found, numbersParsed[r-1][c-1])
					}
				}
				if numbersParsed[r-1][c+1] != 0 {
					if !slices.Contains(_found, numbersParsed[r-1][c+1]) {
						_found = append(_found, numbersParsed[r-1][c+1])
					}
				}
				// next row
				if numbersParsed[r+1][c] != 0 {
					_found = append(_found, numbersParsed[r+1][c])
				}
				if numbersParsed[r+1][c-1] != 0 {
					if !slices.Contains(_found, numbersParsed[r+1][c-1]) {
						_found = append(_found, numbersParsed[r+1][c-1])
					}
				}
				if numbersParsed[r+1][c+1] != 0 {
					if !slices.Contains(_found, numbersParsed[r+1][c+1]) {
						_found = append(_found, numbersParsed[r+1][c+1])
					}
				}
				if len(_found) == 2 {
					found = append(found, _found[0]*_found[1])
				}
			}
		}
	}
	var res int64
	for _, r := range found {
		res = res + r
	}
	fmt.Println(res)
	return res
}

package day02

import (
	"adventOfCode2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func solve() int64 {
	rows := utils.ReadFile("input.txt")

	limits := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var possibleGames []int64
	for _, r := range rows {
		impossible := false
		t := strings.Split(r, ":")
		gameNumber, _ := strconv.ParseInt(strings.Split(t[0], " ")[1], 10, 64)
		games := strings.Split(t[1], ";")
		for _, game := range games {
			thisSet := map[string]int64{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			g := strings.Split(game, ",")
			for _, s := range g {
				c := strings.Split(strings.Trim(s, " "), " ")
				thisSet[c[1]], _ = strconv.ParseInt(c[0], 10, 64)
			}
			for _, color := range []string{"red", "green", "blue"} {
				if thisSet[color] > limits[color] {
					impossible = true
				}
			}
		}
		if !impossible && !slices.Contains(possibleGames, gameNumber) {
			possibleGames = append(possibleGames, gameNumber)
		}
	}
	var solution int64
	for _, game := range possibleGames {
		solution = solution + game
	}
	fmt.Println(solution)

	return solution
}
func solve2() int64 {
	rows := utils.ReadFile("input.txt")
	var solution int64
	for _, r := range rows {
		t := strings.Split(r, ":")
		games := strings.Split(t[1], ";")
		thisSet := map[string]int64{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range games {
			g := strings.Split(game, ",")
			for _, s := range g {
				c := strings.Split(strings.Trim(s, " "), " ")
				color := c[1]
				num, _ := strconv.ParseInt(c[0], 10, 64)
				if thisSet[color] < num {
					thisSet[color] = num
				}
			}
		}
		solution = solution + (thisSet["red"] * thisSet["green"] * thisSet["blue"])
	}

	fmt.Println(solution)

	return solution
}

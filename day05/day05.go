package day05

import (
	"adventOfCode2023/utils"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds   []int64
	Mapping map[string]map[int64]int64
}

func solve(f string) int64 {
	var solution int64
	data := &Almanac{
		Mapping: map[string]map[int64]int64{
			"seed-to-soil":            make(map[int64]int64),
			"soil-to-fertilizer":      make(map[int64]int64),
			"fertilizer-to-water":     make(map[int64]int64),
			"water-to-light":          make(map[int64]int64),
			"light-to-temperature":    make(map[int64]int64),
			"temperature-to-humidity": make(map[int64]int64),
			"humidity-to-location":    make(map[int64]int64),
		},
	}
	raw := utils.ReadFile(f)
	currentKey := ""
	for i, row := range raw {
		if i == 0 {
			data.AddSeeds(row)
			continue
		}
		if row == "" {
			continue
		}
		if data.ChangeTopic(row) != "" {
			currentKey = data.ChangeTopic(row)
			continue
		}
		data.AddMapping(currentKey, row)
	}
	return solution
}

func (a *Almanac) AddSeeds(s string) []int64 {
	seeds := strings.Split(strings.Split(s, ":")[1], " ")
	for _, seed := range seeds {
		parsed, _ := strconv.ParseInt(strings.Trim(seed, " "), 10, 64)
		if parsed == 0 {
			continue
		}
		a.Seeds = append(a.Seeds, parsed)
	}
	return a.Seeds
}

func (a *Almanac) AddMapping(key, s string) map[int64]int64 {
	vals := strings.Split(s, " ")
	destinationRangeStart := toInt(vals[0])
	sourceRangeStart := toInt(vals[1])
	length := toInt(vals[2])
	for i := int64(0); i < length; i++ {
		a.Mapping[key][sourceRangeStart+i] = destinationRangeStart + i
	}
	return a.Mapping[key]
}

func (a *Almanac) ChangeTopic(s string) string {
	for key := range a.Mapping {
		if strings.Contains(s, key) {
			return key
		}
	}
	return ""
}

func toInt(s string) int64 {
	parsed, _ := strconv.ParseInt(strings.Trim(s, " "), 10, 64)
	return parsed
}

package day2

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	re := regexp.MustCompile(`^Game (\d+): (.*)`)
	ans := 0
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			gameId, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			colorCount := getColorCount(matches[2])
			if validGame(colorCount) {
				ans += gameId
			}
		}
	}
	return ans
}

func Part2(input []string) int {
	re := regexp.MustCompile(`^Game (\d+): (.*)`)
	ans := 0
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			colorCount := getColorCount(matches[2])
			power := 1
			for _, v := range colorCount {
				power *= v
			}
			ans += power
		}
	}
	return ans
}

func getColorCount(str string) map[string]int {
	colorCount := make(map[string]int)
	re := regexp.MustCompile(`(?P<count>\d+) (?P<color>\w+)`)
	for _, part := range strings.Split(str, ";") {
		for _, seg := range strings.Split(part, ",") {
			matches := re.FindStringSubmatch(seg)
			if len(matches) < 3 {
				continue
			}
			count, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			color := matches[2]
			if count > colorCount[color] {
				colorCount[color] = count
			}
		}
	}
	return colorCount
}

func validGame(colors map[string]int) bool {
	return colors["red"] <= 12 && colors["green"] <= 13 && colors["blue"] <= 14
}

package day1

import (
	"regexp"
	"strconv"
)

func Part1(input []string) int {
	digitRegexp := regexp.MustCompile(`\d`)
	osum := 0
	for _, line := range input {
		msb := digitRegexp.FindString(line)
		lsb := digitRegexp.FindString(reverse(line))
		num := getNumber(msb)*10 + getNumber(lsb)
		osum += num
	}
	return osum
}
func Part2(input []string) int {
	digitRegexp := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	digitRegexpReverse := regexp.MustCompile(`\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno`)
	osum := 0
	for _, line := range input {
		msb := digitRegexp.FindString(line)
		lsb := reverse(digitRegexpReverse.FindString(reverse(line)))
		num := getNumber(msb)*10 + getNumber(lsb)
		osum += num
	}
	return osum
}
func getNumber(digit string) int {
	if len(digit) == 0 {
		return 0
	}
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	if v, ok := digits[digit]; ok {
		return v
	}
	v, err := strconv.Atoi(digit)
	if err != nil {
		panic(err)
	}
	return v

}
func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

package main

import (
	"fmt"
	"os"

	"example.com/advent-of-code/day1"
	"example.com/advent-of-code/day2"
	"example.com/advent-of-code/utils"
)

func main() {
	day := os.Args[1]
	file := os.Args[2]
	switch day {
	case "1":
		fmt.Println("Day1 Part1 ", day1.Part1(utils.GetInput(file)))
		fmt.Println("Day1 Part2 ", day1.Part2(utils.GetInput(file)))
	case "2":
		fmt.Println("Day2 Part1 ", day2.Part1(utils.GetInput(file)))
		fmt.Println("Day2 Part2 ", day2.Part2(utils.GetInput(file)))
	}

}

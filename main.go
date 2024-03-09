package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/advent-of-code/day1"
)

func main() {
	file := os.Args[1]
	fmt.Println("Part1 ", day1.Part1(getInput(file)))
	fmt.Println("Part2 ", day1.Part2(getInput(file)))
}

func getInput(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	var input []string
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	return input
}

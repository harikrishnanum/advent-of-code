package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(filename string) []string {
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

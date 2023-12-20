package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func intConcat(line string) int {
	var first, last int
	var reverse string

	for _, char := range line {
		reverse = string(char) + reverse
	}

	for _, char := range line {
		if unicode.IsDigit(char) {
			first = int(char - '0')
			break
		}
	}

	for _, char := range reverse {
		if unicode.IsDigit(char) {
			last = int(char - '0')
			break
		}
	}

	return first*10 + last
}

func main() {
	var sum int
	var lines []string

	file, err := os.Open("../res/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, line := range lines {
		sum = intConcat(line) + sum
	}

	fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var red, green, blue int = 12, 13, 14

	file, err := os.Open("../res/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(lines)
}

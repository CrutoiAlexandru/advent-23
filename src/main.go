package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CrutoiAlexandru/advent-23/calibrations"
)

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
		calibration := calibrations.CalibrationsConcat(line)
		sum = calibration + sum
	}

	fmt.Println(sum)
}

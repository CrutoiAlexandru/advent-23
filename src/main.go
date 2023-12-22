package main

import (
	"fmt"

	"github.com/CrutoiAlexandru/advent-23/cubes"
	"github.com/CrutoiAlexandru/advent-23/reader"
)

func main() {
	var sum int
	var lines []string
	var path_to_res string = "../res/input.txt"

	lines, err := reader.ReadResources(path_to_res)
	if err != nil {
		fmt.Println(err)
		return
	}

	sum, err = cubes.GetGamesMap(lines)
	if err != nil {
		fmt.Errorf("error: %w", err)
		return
	}

	fmt.Println(sum)
}

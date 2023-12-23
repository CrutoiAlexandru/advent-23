package cubes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	re := regexp.MustCompile(`[^\d]`)
	string_number := re.ReplaceAllString(s, "")
	// not handling the error is not safe
	number, _ := strconv.Atoi(string_number)

	return number
}

func splitGame(game string) (int, bool) {
	var powers int = 1
	color_map := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	var is_possible bool = true
	turns := strings.Split(game, ";")
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, turn := range turns {
		cubes := strings.Split(turn, ",")

		for _, cube := range cubes {
			for color, number := range rules {
				cube_number := stringToInt(cube)
				if strings.Contains(cube, color) {
					if cube_number > color_map[color] {
						color_map[color] = cube_number
					}
				}
				if strings.Contains(cube, color) && cube_number > number {
					is_possible = false
				}
			}
		}
	}

	powers = color_map["red"] * color_map["green"] * color_map["blue"]

	return powers, is_possible
}

func GetGamesMap(lines []string) (int, error) {
	var games map[int]string = make(map[int]string)
	var sum int = 0
	re := regexp.MustCompile("[0-9]+")

	for _, line := range lines {
		split := strings.SplitN(line, ":", 2)
		index := re.FindString(split[0])
		game := split[1]

		id, err := strconv.Atoi(index)
		if err != nil {
			return 1, fmt.Errorf("error: %w", err)
		}

		games[id] = game
	}

	for _, game := range games {
		powers, _ := splitGame(game)
		sum = sum + powers
	}

	return sum, nil
}

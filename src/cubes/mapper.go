package cubes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func splitGame(game string) bool {
	turns := strings.Split(game, ",")
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, turn := range turns {
		cubes := strings.Split(turn, ",")

		for _, cube := range cubes {
			for color, number := range rules {
				if strings.Contains(cube, color) && number > rules[color] {
					return false
				}
			}
		}
	}

	return true
}

func GetGamesMap(lines []string) (int, error) {
	var games map[int]string
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

	for id, game := range games {
		game_is_real := splitGame(game)
		if game_is_real {
			sum = sum + id
		}
	}

	return sum, nil
}

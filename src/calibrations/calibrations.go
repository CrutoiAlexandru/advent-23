package calibrations

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func reverseString(s string) string {
	var reversed_string string
	for _, char := range s {
		reversed_string = string(char) + reversed_string
	}
	return reversed_string
}

func spellChecker(line string, valid_digits map[string]int) (int, error) {
	var digit_word string = ""

	for _, char := range line {
		if unicode.IsDigit(char) {
			return int(char - '0'), nil
		}

		digit_word = digit_word + string(char)
		for key, value := range valid_digits {
			if strings.Contains(digit_word, key) {
				return value, nil
			}
		}
	}

	return 0, fmt.Errorf("no numbers found")
}

func spellCheckerReverse(line string, valid_digits map[string]int) (int, error) {
	var digit_word string = ""
	line = reverseString(line)

	for _, char := range line {
		if unicode.IsDigit(char) {
			return int(char - '0'), nil
		}

		digit_word = digit_word + string(char)
		for key, value := range valid_digits {
			reversed_key := reverseString(key)
			if strings.Contains(digit_word, reversed_key) {
				return value, nil
			}
		}
	}

	return 0, fmt.Errorf("no numbers found")
}

func CalibrationsConcat(line string) int {
	valid_digits := map[string]int{
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
	var first, last int
	var err error

	first, err = spellChecker(line, valid_digits)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	last, err = spellCheckerReverse(line, valid_digits)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return first*10 + last
}

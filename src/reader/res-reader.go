package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadResources(path_to_res string) ([]string, error) {
	var lines []string

	file, err := os.Open(path_to_res)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	return lines, nil
}

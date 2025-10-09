package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("Failed to open the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	i := 0

	for scanner.Scan() {
		i++
		line := fmt.Sprintf("%d# %s", i, scanner.Text())
		lines = append(lines, line)
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("Failed to read file")
	}

	return lines, nil
}

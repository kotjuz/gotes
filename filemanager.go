package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadFile(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		return errors.New("Failed to open the file")
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
}

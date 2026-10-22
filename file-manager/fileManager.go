package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return []string{}, errors.New("Error at opening file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return []string{}, errors.New("Error at readind file's lines")
	}

	file.Close()

	return lines, nil
}

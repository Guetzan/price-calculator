package filemanager

import (
	"bufio"
	"encoding/json"
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
		return []string{}, errors.New("Error at reading file's lines")
	}

	file.Close()

	return lines, nil
}

func WriteJson(name string, data any) error {
	file, err := os.Create(name)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode data to json")
	}

	file.Close()
	return nil
}

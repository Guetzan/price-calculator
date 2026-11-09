package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		return []string{}, errors.New("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return []string{}, errors.New("Failed to read lines")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return errors.New("Failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) //line to test concurrency

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to encode data to json")
	}

	return nil
}

func New(inputPath, OutputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: OutputPath,
	}
}

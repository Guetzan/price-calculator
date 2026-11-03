package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
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

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return []string{}, errors.New("Failed to read lines")
	}

	file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return errors.New("Failed to create file")
	}

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to encode data to json")
	}

	file.Close()
	return nil
}

func New(inputPath, OutputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: OutputPath,
	}
}

//teste

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

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputPath)

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

func (fm FileManager) WriteJson(data any) error {
	file, err := os.Create(fm.OutputPath)

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

func New(inputPath, OutputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: OutputPath,
	}
}

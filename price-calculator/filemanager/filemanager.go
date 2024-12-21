package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManger) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Could not open file")
	}

	defer file.Close() // close file either error or finish this function

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Reading file failed")
	}

	// file.Close()
	return lines, nil
}

func (fm FileManger) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}
	defer file.Close()
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert to json.")
	}

	// file.Close()
	return nil
}

func NewFileManager(inputPath, outputPath string) FileManger {
	return FileManger{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

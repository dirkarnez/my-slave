package util

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func ReadFile(path string, onFileRead func(*os.File) error) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return onFileRead(file)
}

func ReadString(path string) (*string, error) {
	buf := new(bytes.Buffer)
	err := ReadFile(path, func(file *os.File) error {
		_, err := buf.ReadFrom(file)
		return err
	})

	if err != nil {
		return nil, err
	}

	content := buf.String()
	return &content, nil
}

func ReadLines(path string, onEachLine func(string) error ) error {
	return ReadFile(path, func(file *os.File) error {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			err := onEachLine(scanner.Text())
			if err != nil {
				return err
			}
		}

		return scanner.Err()
	})
}

func CreateFile(path string, onFileCreate func(*os.File) error ) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return onFileCreate(file)
}

func WriteLines(path string, lines []string) error {
	return CreateFile(path, func(file *os.File) error {
		w := bufio.NewWriter(file)

		for _, line := range lines {
			fmt.Fprintln(w, line)
		}

		return w.Flush()
	})
}
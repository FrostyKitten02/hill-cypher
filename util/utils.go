package util

import (
	"errors"
	"os"
	"strings"
)

func ReadFileString(path string) (*string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	str := string(file)
	return &str, nil
}

func SaveFile(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, writeErr := file.WriteString(data)
	if writeErr != nil {
		return err
	}

	defer file.Close()
	return nil
}

func ParseKey(key *string) ([][]int, error) {
	split := strings.Split(*key, "/")
	keysLen := len(split)

	output := make([][]int, keysLen)
	for i, str := range split {
		strLen := len(str)
		if strLen != keysLen {
			return nil, errors.New("Invalid key, all parts must be same length, key must be NxN matrix.")
		}
		line := make([]int, strLen)
		for j, c := range str {
			val := int(c)
			line[j] = val
		}
		output[i] = line
	}

	return output, nil
}

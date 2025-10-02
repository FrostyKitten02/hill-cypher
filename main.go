package main

import (
	"errors"
	"fmt"
	"hill-cypher/cipher"
	"os"
	"strings"
)

func getKey() (*string, error) {
	file, err := os.ReadFile("key.txt")
	if err != nil {
		return nil, err
	}

	str := string(file)
	return &str, nil
}

func readData() (*string, error) {
	file, err := os.ReadFile("data.txt")
	if err != nil {
		return nil, err
	}

	str := string(file)
	return &str, nil
}

func parseKey(key *string) ([][]int, error) {
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

func main() {
	fmt.Println(string(rune(115)))
	keyStr, keyStrErr := getKey()
	if keyStrErr != nil {
		panic(keyStrErr)
	}

	key, keyErr := parseKey(keyStr)
	if keyErr != nil {
		panic(keyErr)
	}

	data, dataErr := readData()
	if dataErr != nil {
		panic(dataErr)
	}

	if key == nil || data == nil {
		panic("WTF NO DATA!!!")
	}

	fmt.Println("Input: " + *data)
	encrypted := cipher.Encrypt(key, *data)
	fmt.Println("Encrypted: " + encrypted)
	decrypted := cipher.Decrypt(key, encrypted)
	fmt.Println("Decrypted: " + decrypted)
}

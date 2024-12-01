package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	// Read file
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)

		return []string{}, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		// get File read line by line in string
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}

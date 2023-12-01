package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Read file
	filePath := "inputs/day1.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		// get File read line by line in string
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	sum := 0

	number_spelled_out := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range fileLines {
		var f1, f2 string

		// First number in string
	first:
		for i := 0; i < len(line); i++ {
			// Loop over the numbers array
			for j := 0; j < len(number_spelled_out); j++ {
				// last_index is length of the spelled out number or length of the string
				last_index := min((i + len(number_spelled_out[j])), len(line))

				// Get the string number of the length of the number
				string_number := (line[i:last_index])

				if string_number == number_spelled_out[j] {
					f1 = strconv.Itoa(j)
					break first
				}
			}

			// Check if number is valid
			_, err := strconv.Atoi(string(line[i]))

			if err != nil {
				continue
			}

			f1 = string(line[i])
			break
		}

		// Second Number in string
	second:
		for i := len(line) - 1; i >= 0; i-- {
			for j := 0; j < len(number_spelled_out); j++ {
				/*
					reading order remains the same => one * // eno x
					So, last_index remains > start_index(i)
					reading in the same fashion as front but from the end of the string instead of front
					if it had to be read in reverse as well
					(eno is valid then just reverse the line string and number_spelled_out[j] and go with first loop)
				*/
				last_index := min((i + len(number_spelled_out[j])), len(line))
				string_number := (line[i:last_index])
				if string_number == (number_spelled_out[j]) {

					f2 = strconv.Itoa(j)
					break second
				}
			}

			_, err := strconv.Atoi(string(line[i]))

			if err != nil {
				continue
			}

			f2 = string(line[i])
			break
		}

		number := f1 + f2

		i, err := strconv.Atoi(string(number))

		if err != nil {
			fmt.Println(err)
		}

		sum = sum + i
	}

	fmt.Println(sum)
}

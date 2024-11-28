package main

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type indexValue struct {
	number    string
	end_index int
}

type position struct {
	row    int
	column int
}

type symbolPosition struct {
	symbol   string
	position position
}

type numberSymbol struct {
	number   int
	position position
}

func isStringNumber(letter string) bool {
	string_numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	value := false
	for _, number := range string_numbers {
		if letter == number {
			value = true
			break
		}
	}

	return value
}

func getNumber(line string, index int, number string) indexValue {
	if index < len(line) && isStringNumber(string(line[index])) {
		number = number + string(line[index])
		return getNumber(line, index+1, number)
	} else {
		value := indexValue{number, index - 1}

		return value
	}
}

func hasSymbolAdjacent(lines []string, line_index int, start_index int, end_index int) symbolPosition {
	value := ""
	str_position := position{0, 0}

outer:
	for i := utils.Max(line_index-1, 0); i <= utils.Min(line_index+1, len(lines)-1); i++ {
		for j := start_index; j <= end_index; j++ {
			letter := string(lines[i][j])
			if letter == "." || isStringNumber(letter) {
				continue
			}
			str_position = position{i, j}
			value = letter
			break outer

		}
	}

	return symbolPosition{symbol: value, position: str_position}
}

func main() {
	filepath := "inputs/day3.txt"
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var file_lines []string

	for scanner.Scan() {
		file_lines = append(file_lines, scanner.Text())
	}

	sum := 0

	gearRatio := 0

	var star_numbers []numberSymbol

	for line := 0; line < len(file_lines); line++ {
		for index := 0; index < len(file_lines[line]); index++ {
			letter := string(file_lines[line][index])

			if letter == "." {
				continue
			}

			if isStringNumber(string(file_lines[line][index])) {
				start_index := index
				str_number := letter
				numberIndex := getNumber(file_lines[line], index+1, str_number)
				end_index := numberIndex.end_index

				str_number = numberIndex.number
				index = end_index
				number, err := strconv.Atoi(str_number)
				if err != nil {
					panic(err)
				}
				adjacent := hasSymbolAdjacent(file_lines, line, utils.Max(start_index-1, 0), utils.Min(end_index+1, len(file_lines[line])-1))
				if adjacent.symbol != "" {
					sum = sum + number
				}

				if adjacent.symbol == "*" {
					exists := false
					for i := 0; i < len(star_numbers); i++ {
						if star_numbers[i].position == adjacent.position {
							gearRatio = gearRatio + number*star_numbers[i].number
							exists = true
							break
						}
					}
					if exists == false {
						star_numbers = append(star_numbers, numberSymbol{number, adjacent.position})
					}
				}

			}

		}

	}

	fmt.Println(sum, gearRatio)
}

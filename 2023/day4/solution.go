package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "inputs/day4.txt"
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var file_lines []string

	for scanner.Scan() {
		file_lines = append(file_lines, scanner.Text())
	}

	points := 0

	var winning_games []int

	for i := 0; i < len(file_lines); i++ {
		winning_games = append(winning_games, 1)
	}

	for line_number := 0; line_number < len(file_lines); line_number++ {
		line := file_lines[line_number]
		str := strings.Split(line, ":")
		numbers := strings.Split(str[1], "|")
		winnings := 0

		var winning_numbers []int
		var ticket_numbers []int

		str_winning_numbers_array := strings.Split(numbers[0], " ")
		str_ticket_numbers_array := strings.Split(numbers[1], " ")

		for i := 0; i < len(str_winning_numbers_array); i++ {
			value, err := strconv.Atoi(str_winning_numbers_array[i])

			if err != nil {
				continue
			}

			winning_numbers = append(winning_numbers, value)
		}

		for i := 0; i < len(str_ticket_numbers_array); i++ {
			value, err := strconv.Atoi(str_ticket_numbers_array[i])

			if err != nil {
				continue
			}

			ticket_numbers = append(ticket_numbers, value)

			for j := 0; j < len(winning_numbers); j++ {
				if winning_numbers[j] == value {
					winnings = winnings + 1
				}
			}
		}
		for i := 1; i <= winnings; i++ {
			for j := 0; j < winning_games[line_number]; j++ {
				winning_games[line_number+i] = winning_games[line_number+i] + 1
			}
		}

		if winnings > 0 {
			moneysss := math.Pow(2, float64(winnings-1))
			points = points + int(moneysss)
		}

	}

	new_points := 0

	for i := 0; i < len(winning_games); i++ {
		new_points = new_points + winning_games[i]
	}

	fmt.Println(points, new_points)
}

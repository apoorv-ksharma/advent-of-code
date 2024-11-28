package main

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	red        int
	blue       int
	green      int
	gameNumber int
}

func main() {
	filepath := "inputs/day2.txt"
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var file_lines []string

	for scanner.Scan() {
		file_lines = append(file_lines, scanner.Text())
	}

	var games []game

	for _, line := range file_lines {
		str := strings.Split(line, ":")

		gameNumber, err := strconv.Atoi(strings.Split(str[0], " ")[1])

		if err != nil {
			panic(err)
		}

		game := game{red: 0, green: 0, blue: 0, gameNumber: gameNumber}

		subsets := strings.Split(str[1], ";")

		for _, subset := range subsets {
			results := strings.Split(subset, ",")

			for _, result := range results {
				number, err := strconv.Atoi(strings.Split(result, " ")[1])

				if err != nil {
					panic(err)
				}

				color := strings.Split(result, " ")[2]

				if color == "red" {
					game.red = utils.Max(number, game.red)
				}
				if color == "blue" {
					game.blue = utils.Max(number, game.blue)
				}
				if color == "green" {
					game.green = utils.Max(number, game.green)
				}
			}

		}

		games = append(games, game)

	}

	sum := 0
	power := 0

	for _, game := range games {
		red, blue, green, gameNumber := game.red, game.blue, game.green, game.gameNumber

		if red <= 12 && blue <= 14 && green <= 13 {
			sum = sum + gameNumber
		}

		multiple := red * blue * green

		power = power + multiple
	}

	fmt.Println("Sum", sum, "Power", power)
}

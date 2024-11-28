package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type crop struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type almanac_map struct {
	source      int
	destination int
	ext_range   int
}

type maps struct {
}

func main() {
	file_lines, err := utils.ReadLines("day5/example.txt")
	if err != nil {
		panic(err)
	}
	str_seeds := strings.Split(strings.Split(file_lines[0], ":")[1], " ")

	seeds := make([]int, 0)

	for i := 0; i < len(str_seeds); i++ {
		value, err := strconv.Atoi(str_seeds[i])
		if err == nil {
			seeds = append(seeds, value)
		}
	}

	var seed_to_soil []almanac_map
	var soil_to_fertilizer []almanac_map
	var fertilizer_to_water []almanac_map
	var water_to_light []almanac_map
	var light_to_temperature []almanac_map
	var temperature_to_humidity []almanac_map
	var humidity_to_location []almanac_map

	for i := 2; i < len(file_lines); i++ {
	}

	fmt.Println(seeds)
}

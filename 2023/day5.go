package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type crop struct {
	seed_start int
	seed_range int
}

type cropMap struct {
	source      string
	destination string
	s_number    int
	d_number    int
	crop_range  int
}

func findCropProp(source string, destination string, s_number int, array []cropMap) int {
	result := s_number
	for i := 0; i < len(array); i++ {
		crop := array[i]

		if crop.source == source && crop.destination == destination {

			min := crop.s_number
			max := crop.s_number + crop.crop_range
			if min <= s_number && max >= s_number {
				range_to_add := s_number - crop.s_number

				if range_to_add < 0 {
					fmt.Println(crop.crop_range, crop.s_number, s_number)
				}

				result = crop.d_number + range_to_add
				break
			}
		}

	}
	return result

}

func main() {
	file_lines, err := utils.ReadLines("inputs/day5.txt")

	if err != nil {
		fmt.Println("Error reading files")
	}

	var crops []crop

	var crop_map []cropMap

	var source string
	var destination string

	for line_number := 0; line_number < len(file_lines); line_number++ {

		str := strings.Split(file_lines[line_number], ":")
		if len(str) == 1 && str[0] == "" {
			continue
		}

		if str[0] == "seeds" {
			seeds := strings.Split(str[1], " ")

			for i := 0; i < len(seeds); i++ {
				if seeds[i] == "" {
					continue
				}

				start, err1 := strconv.Atoi(string(seeds[i]))
				seed_range, err2 := strconv.Atoi(string(seeds[i+1]))

				i = i + 1

				if err1 != nil || err2 != nil {
					continue
				}

				crops = append(crops, crop{seed_start: start, seed_range: seed_range})
			}

			continue
		}

		str_map := strings.Split(str[0], " ")

		if len(str_map) > 1 && str_map[1] == "map" {
			cat := strings.Split(str_map[0], "-")
			source = cat[0]
			destination = cat[2]
			continue
		}

		values := strings.Split(str[0], " ")

		var s_number int
		var d_number int
		var crop_range int

		for i := 0; i < len(values); i++ {
			number, err := strconv.Atoi(string(values[i]))

			if err != nil {
				continue
			}

			if i == 0 {
				d_number = number
			}
			if i == 1 {
				s_number = number
			}
			if i == 2 {
				crop_range = number
			}
		}
		crop_to_add := cropMap{source: source, destination: destination, crop_range: crop_range, s_number: s_number, d_number: d_number}
		crop_map = append(crop_map, crop_to_add)
	}
	min_location := 0

	fmt.Println(len(crops))

	for _, crop := range crops {
		var wg sync.WaitGroup
		wg.Add(crop.seed_range)
		fmt.Println(crop.seed_range)
		for i := 0; i < crop.seed_range; i++ {
			go func(i int) {
				defer wg.Done()
				seed := crop.seed_start + i
				soil := (findCropProp("seed", "soil", seed, crop_map))
				fertilizer := (findCropProp("soil", "fertilizer", soil, crop_map))
				water := (findCropProp("fertilizer", "water", fertilizer, crop_map))
				light := (findCropProp("water", "light", water, crop_map))
				temperature := (findCropProp("light", "temperature", light, crop_map))
				humidity := (findCropProp("temperature", "humidity", temperature, crop_map))
				location := (findCropProp("humidity", "location", humidity, crop_map))

				if i == 0 {
					min_location = location
				}

				min_location = utils.Min(min_location, location)
			}(i)
		}

		wg.Wait()
	}

	fmt.Println(min_location)
}

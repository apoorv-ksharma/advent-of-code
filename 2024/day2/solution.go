package main

import (
	utils "advent_of_code"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
	increasing, decreasing := false, false

	for i := 0; i < len(report)-1; i++ {
		difference := report[i+1] - report[i]

		increasing, decreasing = increasing || difference > 0, decreasing || difference < 0

		switch true {
		case difference == 0:
			fallthrough
		case increasing == decreasing:
			fallthrough
		case difference > 3:
			fallthrough
		case difference < -3:
			return false
		}
	}
	return true
}

func main() {
	lines, err := utils.ReadLines("day2/problem.txt")

	if err != nil {
		fmt.Println(err)

		return
	}

	var reports [][]int

	for _, line := range lines {
		str_numbers := strings.Split(line, " ")
		var array []int
		for _, str_number := range str_numbers {
			number, err := strconv.Atoi(str_number)
			if err == nil {
				array = append(array, number)
			}
		}
		reports = append(reports, array)
	}

	var safe_reports [][]int
	var safe_reports_with_bad_level [][]int

	for _, report := range reports {
		if isReportSafe(report) {
			safe_reports = append(safe_reports, report)
		}

		for i := range report {
			if isReportSafe(slices.Delete(slices.Clone(report), i, i+1)) {
				safe_reports_with_bad_level = append(safe_reports_with_bad_level, report)
				break
			}
		}
	}

	fmt.Println(len(safe_reports))
	fmt.Println(len(safe_reports_with_bad_level))
}

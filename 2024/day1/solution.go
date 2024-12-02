package main

import (
	utils "advent_of_code"
	"fmt"
	"strconv"
	"strings"
)

func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}

	return array
}

func main() {
	lines, err := utils.ReadLines("day1/problem.txt")

	if err != nil {
		fmt.Println(err)

		return
	}

	var leftArr, rightArr []int

	for _, line := range lines {
		numbers := strings.Split(line, "   ")

		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])

		leftArr, rightArr = append(leftArr, left), append(rightArr, right)
	}
	leftArr, rightArr = bubbleSort(leftArr), bubbleSort(rightArr)

	var sum int
	var similarity int

	for i := 0; i < len(leftArr); i++ {
		var distance int
		if rightArr[i] > leftArr[i] {
			distance = rightArr[i] - leftArr[i]
		} else {
			distance = leftArr[i] - rightArr[i]

		}
		sum += distance
	}

	for i := 0; i < len(leftArr); i++ {
		score := 0
		for j := 0; j < len(leftArr); j++ {
			if leftArr[i] == rightArr[j] {
				score += 1
			}
		}
		similarity += leftArr[i] * score
	}

	fmt.Println(sum, similarity)
}

package main

import (
	"2024/shared"
	"fmt"
	"strconv"

	// "io"
	"log"
)

func checkCondition(arrayInput []int, condition func(int, int) bool) (bool, int) {
	isCondition := false
	increment := 0
	prevIncrement := 0

	array := append([]int(nil), arrayInput...)

	for i, curr := range array[1:] {
		//fmt.Println("index: ", i, "Current: ", curr)

		prev := array[i]
		if condition(prev, curr) {
			currIncrement := shared.AbSub(prev, curr)
			if prevIncrement < currIncrement {
				increment = currIncrement
				prevIncrement = increment
			}
			//fmt.Println("Increment: ", increment)
		} else {
			return false, increment
		}
	}
	isCondition = true

	return isCondition, increment
}

func decreaseCheck(array []int) (bool, int) {
	return checkCondition(array, func(prev, curr int) bool { return prev > curr })
}
func increaseCheck(array []int) (bool, int) {
	return checkCondition(array, func(prev, curr int) bool { return prev < curr })
}

// part 1
func checkSafe(linesInput []string) int {
	lines := append([]string(nil), linesInput...)
	safeList := []int{}

	safeCount := 0

	for i, line := range lines {
		fmt.Println(i, line)
		lineArray := []int{}
		var numberString string
		var number int
		for i, _ := range line {

			//fmt.Println(i, "->", line[i:i+1])
			if line[i:i+1] == " " {
				lineArray = append(lineArray, number)
				numberString = ""
			} else {
				numberString += line[i : i+1]
			}
			number, _ = strconv.Atoi(numberString)
		}
		lineArray = append(lineArray, number)
		numberString = ""
		fmt.Println("lineArray", lineArray)

		decrease, incrementDecrease := decreaseCheck(lineArray)
		fmt.Println("Decrease: ", decrease, "Increment: ", incrementDecrease)

		increase, incrementIncrease := increaseCheck(lineArray)
		fmt.Println("Increase: ", increase, "Increment: ", incrementIncrease)

		if (decrease && (incrementDecrease <= 3)) || increase && (incrementIncrease <= 3) {
			safeList = append(safeList, i)
			goto safe
		}

		continue
	safe:
		safeCount++

	}
	fmt.Println("SafeCount: ", safeCount)
	fmt.Println("safeList: ", safeList)

	return safeCount
}

// part 2
func checkSafe2(linesInput []string) int {
	lines := append([]string(nil), linesInput...)
	safeList := []int{}
	safeListOG := []int{}
	safeListNew := []int{}

	safeCount := 0

	for i, line := range lines {
		curr_index_to_remove := 0
		fmt.Println(i, line)
		lineArray := []int{}
		var numberString string
		var number int
		for i, _ := range line {

			//fmt.Println(i, "->", line[i:i+1])
			if line[i:i+1] == " " {
				lineArray = append(lineArray, number)
				numberString = ""
			} else {
				numberString += line[i : i+1]
			}
			number, _ = strconv.Atoi(numberString)
		}
		lineArray = append(lineArray, number)
		numberString = ""
		//fmt.Println("lineArray", lineArray)
		lineArrayCopy := append([]int(nil), lineArray...)

	try_again:
		decrease, incrementDecrease := decreaseCheck(lineArrayCopy)
		fmt.Println("Decrease: ", decrease, "Increment: ", incrementDecrease)
		increase, incrementIncrease := increaseCheck(lineArrayCopy)
		fmt.Println("Increase: ", increase, "Increment: ", incrementIncrease)

		if (decrease && (incrementDecrease <= 3)) || increase && (incrementIncrease <= 3) {
			if curr_index_to_remove > 0 {
				fmt.Println("SafeArray: ", lineArray, "DampedArray", lineArrayCopy, "curr_index_to_remove: ", curr_index_to_remove)
				safeListNew = append(safeListNew, i)
			}
			if curr_index_to_remove == 0 {
				safeListOG = append(safeListOG, i)
			}
			safeList = append(safeList, i)

			goto safe
		}
		fmt.Println("array: ", lineArrayCopy)
		fmt.Println("curr_index_to_remove: ", curr_index_to_remove)

		lineArrayCopy = append([]int(nil), lineArray...)

		if curr_index_to_remove == len(lineArray) {
			continue
		}

		lineArrayCopy = shared.Remove(lineArrayCopy, curr_index_to_remove)

		curr_index_to_remove++
		goto try_again

		continue

	safe:
		safeCount++

	}

	fmt.Println("Part 2 - SafeCount: ", safeCount)
	fmt.Println("safeList: ", safeList)
	fmt.Println("safeListOG: ", safeListOG)
	fmt.Println("safeListNew: ", safeListNew)
	fmt.Println("safeListLenOG Part 1: ", len(safeListOG))
	fmt.Println("safeListLen Part 2: ", len(safeList))
	fmt.Println("safeListLenNew: ", len(safeListNew))

	return safeCount
}

func main() {

	lines, err := shared.ReadLines("day-2/input-day-2.txt")
	if err != nil {
		log.Fatalf("readFail: %s", err)
	}

	//checkSafe(lines)

	checkSafe2(lines)

}

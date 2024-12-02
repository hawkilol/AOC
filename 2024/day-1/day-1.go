package main

import (
	"2024/shared"
	"fmt"
	"strconv"

	// "io"
	"log"
)

// Part 1
func diffSum(firstListInput []int, secondListInput []int) {
	firstList := append([]int(nil), firstListInput...)
	secondList := append([]int(nil), secondListInput...)
	// name := "kalil"
	// fmt.Print(name + " Hello :)\n")
	// input := openFile("input-day-1.txt")
	// print(input[:5])
	sum := 0
	for i := range firstList {
		fmt.Println(i)
		smallestFirst, indexSmallestFirst := shared.GetSmallest(firstList)
		smallestSecond, indexSmallestSecond := shared.GetSmallest(secondList)
		fmt.Println("1st", smallestFirst, "2nd", smallestSecond)

		diffLocal := shared.AbSub(smallestFirst, smallestSecond)
		fmt.Println("diffLocal", diffLocal)
		sum += diffLocal

		firstList = shared.Remove(firstList, indexSmallestFirst)
		secondList = shared.Remove(secondList, indexSmallestSecond)

	}

	fmt.Println("Sum: ", sum)
}

// Part 2
func similarySum(firstListInput []int, secondListInput []int) {
	firstList := append([]int(nil), firstListInput...)
	secondList := append([]int(nil), secondListInput...)
	sum := 0

	for _, firstN := range firstList {
		similaryCount := 0

		fmt.Println("1st", firstN)

		similaryCount = shared.ContainsCount(secondList, firstN)
		similaryCount *= firstN
		fmt.Println("similaryCount", similaryCount)
		sum += similaryCount

	}
	fmt.Println("Sum: ", sum)
}

func main() {
	inputSizes := []int{5, 8}
	// inputSizes := []int{1, 4}
	// inputSizes := []int{5, 8}
	print(inputSizes[1])
	lines, err := shared.ReadLines("day-1/input-day-1.txt")
	if err != nil {
		log.Fatalf("readFail: %s", err)
	}

	var firstListInput []int
	var secondListInput []int

	for i, line := range lines {
		fmt.Println(i, line)

		firstN, _ := strconv.Atoi(line[:inputSizes[0]])
		secondN, _ := strconv.Atoi(line[inputSizes[1]:])

		firstListInput = append(firstListInput, firstN)
		secondListInput = append(secondListInput, secondN)
	}

	fmt.Println("firstListInput", firstListInput)
	fmt.Println("secondListInput", secondListInput)

	diffSum(firstListInput, secondListInput)
	fmt.Println("firstListInput", firstListInput)
	fmt.Println("secondListInput", secondListInput)
	similarySum(firstListInput, secondListInput)

}

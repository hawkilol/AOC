package main

import (
	"bufio"
	"fmt"
	"strconv"

	// "io"
	"log"
	"os"
)

func abSub(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func containsCount(array []int, value int) int {
	count := 0
	for _, e := range array {
		if value == e {
			count++
		}
	}
	return count
}

//	func openFile(fileName string) string {
//		file, error := os.Open(fileName)
//		if error != nil {
//			fmt.Print("Error :(")
//			log.Fatal(error)
//		}
//		defer func() {
//			if error = file.Close(); error != nil {
//				log.Fatal(error)
//			}
//		}()
//		bytes, error := io.ReadAll(file)
//		return string(bytes)
//	}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func getSmallest(array []int) (int, int) {
	smallestE := array[0]
	indexSmallest := 0
	for i, e := range array {
		if e < smallestE {
			smallestE = e
			indexSmallest = i
		}
	}
	return smallestE, indexSmallest
}

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
		smallestFirst, indexSmallestFirst := getSmallest(firstList)
		smallestSecond, indexSmallestSecond := getSmallest(secondList)
		fmt.Println("1st", smallestFirst, "2nd", smallestSecond)

		diffLocal := abSub(smallestFirst, smallestSecond)
		fmt.Println("diffLocal", diffLocal)
		sum += diffLocal

		firstList = remove(firstList, indexSmallestFirst)
		secondList = remove(secondList, indexSmallestSecond)

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

		similaryCount = containsCount(secondList, firstN)
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
	lines, err := readLines("input-day-1.txt")
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

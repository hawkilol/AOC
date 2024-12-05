package shared

import (
	"bufio"
	"strconv"

	// "io"
	"os"
)

func AbSub(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
func ReadLines(path string) ([]string, error) {
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

func ContainsCount(array []int, value int) int {
	count := 0
	for _, e := range array {
		if value == e {
			count++
		}
	}
	return count
}

func Contains(array []int, value int) bool {
	for _, e := range array {
		if value == e {
			return true
		}
	}
	return false
}

func ContainString(array []string, value string) bool {
	for _, e := range array {
		if value == e {
			return true
		}
	}
	return false
}
func ContainStringIndex(array []string, value string) (bool, int) {
	for i, e := range array {
		if value == e {
			return true, i
		}
	}
	return false, 0
}
func ContainLastStringIndex(array []string, value string) (bool, int) {
	match := false
	indexMatch := 0
	for i, e := range array {
		if value == e {
			match = true
			indexMatch = i
		}
	}
	return match, indexMatch
}
func ContainIntIndex(array []int, value int) (bool, int) {
	for i, e := range array {
		if value == e {
			return true, i
		}
	}
	return false, 0
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
func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func GetSmallest(array []int) (int, int) {
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

func IsInt(stringInput string) bool {
	if _, err := strconv.Atoi(stringInput); err == nil {
		return true
	}
	return false
}

func CheckCondition(arrayInput []int, condition func(int, int) bool) (bool, int) {
	isCondition := false
	increment := 0
	prevIncrement := 0

	array := append([]int(nil), arrayInput...)

	for i, curr := range array[1:] {
		//fmt.Println("index: ", i, "Current: ", curr)

		prev := array[i]
		if condition(prev, curr) {
			currIncrement := AbSub(prev, curr)
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

func DecreaseCheck(array []int) (bool, int) {
	return CheckCondition(array, func(prev, curr int) bool { return prev > curr })
}
func IncreaseCheck(array []int) (bool, int) {
	return CheckCondition(array, func(prev, curr int) bool { return prev < curr })
}
func IncreaseOrEqualCheck(array []int) (bool, int) {
	return CheckCondition(array, func(prev, curr int) bool { return prev <= curr })
}

func RangeSlice(start, end int) []int {
	slice := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		slice = append(slice, i)
	}
	return slice
}

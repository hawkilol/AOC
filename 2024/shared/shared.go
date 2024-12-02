package shared

import (
	"bufio"
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

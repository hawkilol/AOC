package main

import (
	"2024/shared"
	"fmt"
	"strconv"
	"strings"

	// "io"
	"log"
)

// Rotate 90 degrees
func rotateInput(linesInput []string) []string {
	rows := len(linesInput)
	cols := len(linesInput[0])

	rotatedMatrix := make([][]byte, cols)

	rotatedString := []string{}

	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]byte, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//rotatedMatrix[i][j] = linesInput[j][i]
			rotatedMatrix[j][rows-1-i] = linesInput[i][j]
		}
	}

	for i := 0; i < cols; i++ {
		fmt.Println("Rotated: ", string(rotatedMatrix[i]))
		rotatedString = append(rotatedString, string(rotatedMatrix[i]))
	}
	return rotatedString

}

// Rotate 90 degrees counterclockwise (negative 90)
func rotateInputNegative90(linesInput []string) []string {
	rows := len(linesInput)    // Get number of rows
	cols := len(linesInput[0]) // Get number of columns

	// Create a new matrix for the rotated result (cols x rows)
	rotatedMatrix := make([][]byte, cols)

	// Initialize each row of the rotated matrix
	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]byte, rows)
	}

	// Fill in the rotated matrix by shifting and transposing indices
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotatedMatrix[cols-j-1][i] = linesInput[i][j]
		}
	}

	// Convert the rotated matrix to a slice of strings
	rotatedString := []string{}
	for i := 0; i < cols; i++ {
		rotatedString = append(rotatedString, string(rotatedMatrix[i]))
	}

	return rotatedString
}
func revert45Rotation(rotatedInput []string, rows, cols int) []string {
	// Create original matrix filled with spaces
	originalMatrix := make([][]byte, rows)
	for i := range originalMatrix {
		originalMatrix[i] = make([]byte, cols)
		for j := range originalMatrix[i] {
			originalMatrix[i][j] = ' ' // Fill with spaces
		}
	}
	reversedStrings := make([]string, rows)
	switchSides := false
	// Reverse the diagonal rotation
	for diagonal, diagonalString := range rotatedInput {
		for index, ch := range diagonalString {
			// Calculate row and column based on diagonal and index
			row := index
			col := diagonal - index

			// rowTest := index
			// colTest := index
			// fmt.Println("rows: ", rows)
			// fmt.Println("row: ", row, "col: ", col)
			// fmt.Println("diagonalString: ", diagonalString)
			// fmt.Println("diagonalString Len: ", len(diagonalString))
			// fmt.Println("index: ", index)
			//fmt.Println("test row: ", row, " test col: ", col)

			// fmt.Println("Char: ", string(ch))
			if col >= cols {

				diff := ((rows - 1) - len(diagonalString)) + 1
				row = row + diff
				col = col - diff

				// fmt.Println("diff: ", diff)
				// fmt.Println("Match row: ", row, "Match col: ", col)
				originalMatrix[row][col] = byte(ch)
				switchSides = true
				// fmt.Println("switchSides", switchSides)
				// fmt.Println("originalMatrix: ", originalMatrix)

				continue
			}
			if switchSides {

				diff := ((rows - 1) - len(diagonalString)) + 1
				row = row + diff
				col = col - diff
				// fmt.Println("diff: ", diff)

				// fmt.Println("switchSides row: ", row, "switchSides col: ", col)
				originalMatrix[row][col] = byte(ch)
				// fmt.Println("originalMatrix: ", originalMatrix)

				continue
			}

			if row >= 0 && row < rows && col >= 0 && col < cols {
				originalMatrix[row][col] = byte(ch)

			}

			// fmt.Println("originalMatrix: ", originalMatrix)
		}

	}
	for i := 0; i < rows; i++ {
		reversedStrings[i] = string(originalMatrix[i])
	}

	return reversedStrings
}

// Rotate 45 degrees diagonals
func rotateInput45(linesInput []string) []string {
	rows := len(linesInput)
	cols := len(linesInput[0])
	diagonals := rows + cols - 1
	rotatedMatrix := make([][]byte, diagonals)

	for i := 0; i < diagonals; i++ {
		rotatedMatrix[i] = []byte{}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotatedMatrix[i+j] = append(rotatedMatrix[i+j], linesInput[i][j])
		}
	}

	rotatedString := make([]string, 0, diagonals)
	for i := 0; i < diagonals; i++ {
		rotatedString = append(rotatedString, string(rotatedMatrix[i]))
		fmt.Println("rotatedString: ", string(rotatedMatrix[i])) // Debugging the rotated string
	}

	return rotatedString
}

// Rotate 45 degrees diagonals
func rotateInputNegative45(linesInput []string) []string {
	rows := len(linesInput)
	cols := len(linesInput[0])
	diagonals := rows + cols - 1
	rotatedMatrix := make([][]byte, diagonals)

	for i := range rotatedMatrix {
		rotatedMatrix[i] = []byte{}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			index := j - i + (cols - 1)
			if index >= 0 && index < diagonals {
				rotatedMatrix[index] = append(rotatedMatrix[index], linesInput[i][j])
			}
		}
	}

	rotatedString := []string{}
	for i := 0; i < diagonals; i++ {
		rotatedString = append(rotatedString, string(rotatedMatrix[i]))
	}

	return rotatedString
}

// part 1
// Check if cleanXmas correct
func cleanStringMatch(linesInput []string, matchString []string, matchMatrixIndexes [][][2]int) ([]string, int) {
	lines := append([]string(nil), linesInput...)
	cleanMulArray := []string{}
	cleanMulArrayOrder := []int{}
	// matchString := []string{"m", "u", "l", "(", ",", ")"}
	mulIndexes := shared.RangeSlice(0, len(matchString)-1)
	fmt.Println("mulIndexes", mulIndexes)
	sum := 0
	for j, line := range lines {
		fmt.Println(j, line)
		var numberString string
		//var number int
		indexMatchOrder := []int{}
		indexMatchCorrect := false
		isFirstNumber := true
		numberX := 0
		numberY := 0
		X := []int{0, 0}
		Y := []int{0, 0}

		increment := 0
		for i := range line {

			currChar := line[i : i+1]
			currCharCheck := currChar
			//fmt.Println(i, "->", currChar)

			// if its a number make a exception then pass the contains and keep the list
			if shared.IsInt(currChar) {
				if isFirstNumber {
					currCharCheck = "IntX"
				} else {
					currCharCheck = "IntY"
				}

				// save the number 1 from 1-3(if out of range kill even matchorder), stop saving then save number 2
				numberString += currChar

			}

			contains, indexMatch := shared.ContainStringIndex(matchString, currCharCheck)

			if contains {
				indexMatchOrder = append(indexMatchOrder, indexMatch)
				cleanMulArrayOrder = append(cleanMulArrayOrder, indexMatch)

				if isFirstNumber {
					numberX, _ = strconv.Atoi(numberString)
					isFirstNumber = false
					numberString = ""
					X[0] = j
					Y[0] = i

				} else {
					// if should switch
					X[1] = j
					Y[1] = i
					numberString = ""
					numberX, numberY = 0, 0
					isFirstNumber = true

					numberY, _ = strconv.Atoi(numberString)

				}

				// if its a number make a exception then pass the contains and keep the list
			} else {
				indexMatchOrder = []int{}
				numberString = ""
				numberX, numberY = 0, 0
				isFirstNumber = true
				X = []int{0, 0}
				Y = []int{0, 0}
			}

			fmt.Println("Order: ", indexMatchOrder)

			// fmt.Println("Numberstring: ", numberString)
			//fmt.Println("indexMatchOrder: ", indexMatchOrder)
			// fmt.Println("numberX: ", numberX)
			// fmt.Println("numberY: ", numberY)
			//, if the indexMatch is correct then put these numbers on hash map or just to lists to multiply later
			if len(indexMatchOrder) > 0 {
				indexMatchCorrect, increment = shared.CheckCondition(indexMatchOrder, func(prev, curr int) bool { return (prev < curr) })
				if indexMatchCorrect && increment <= 1 {
					//check if all are
					containsAll := true
					for _, e := range mulIndexes {
						containsFinal, _ := shared.ContainIntIndex(indexMatchOrder, e)
						if !containsFinal {
							containsAll = false
						}
					}

					if containsAll {
						//fmt.Println("Save X: ", numberX, "Save Y: ", numberY)

						stringClean := "mul(" + strconv.Itoa(numberX) + "," + strconv.Itoa(numberY) + ")"
						cleanMulArray = append(cleanMulArray, stringClean)

						fmt.Println("X[0]: ", X[0], "Y[0]}: ", Y[0], "{X[1]: ", X[1], "Y[1]: ", Y[1])

						matchMatrixIndexes = append(matchMatrixIndexes, [][2]int{{X[0], Y[0]}, {X[1], Y[1]}})

						sum++
					}

				} else {

					if indexMatchOrder[len(indexMatchOrder)-1] == mulIndexes[0] {
						indexMatchOrder = []int{mulIndexes[0]}
						fmt.Println("NewindexMatchOrder: ", indexMatchOrder)
						isFirstNumber = true
					} else {
						indexMatchOrder = []int{}
						numberString = ""
						numberX, numberY = 0, 0
						isFirstNumber = true
					}

				}
			}
		}

	}
	fmt.Println("cleanMulArrayOrder: ", cleanMulArrayOrder)

	return cleanMulArray, sum
}

// part 2
// Check if cleanXmas correct 2
func cleanStringMatch2(linesInput []string, matchString []string, matchMatrixIndexes [][][2]int) ([]string, []string, int) {
	lines := append([]string(nil), linesInput...)
	cleanMulArray := []string{}
	cleanMulArrayOrder := []int{}
	// matchString := []string{"m", "u", "l", "(", ",", ")"}
	mulIndexes := shared.RangeSlice(0, len(matchString)-1)
	fmt.Println("mulIndexes", mulIndexes)
	sum := 0
	for j, line := range lines {
		fmt.Println(j, line)
		var numberString string
		//var number int
		indexMatchOrder := []int{}
		indexMatchCorrect := false
		isFirstNumber := true
		numberX := 0
		numberY := 0
		X := []int{0, 0}
		Y := []int{0, 0}

		increment := 0
		for i := range line {

			currChar := line[i : i+1]
			currCharCheck := currChar
			//fmt.Println(i, "->", currChar)

			// if its a number make a exception then pass the contains and keep the list
			if shared.IsInt(currChar) {
				if isFirstNumber {
					currCharCheck = "IntX"
				} else {
					currCharCheck = "IntY"
				}

				// save the number 1 from 1-3(if out of range kill even matchorder), stop saving then save number 2
				numberString += currChar

			}

			contains, indexMatch := shared.ContainStringIndex(matchString, currCharCheck)

			if contains {
				indexMatchOrder = append(indexMatchOrder, indexMatch)
				cleanMulArrayOrder = append(cleanMulArrayOrder, indexMatch)

				if isFirstNumber {
					numberX, _ = strconv.Atoi(numberString)
					isFirstNumber = false
					numberString = ""
					X[0] = j
					Y[0] = i

				} else {
					// if should switch
					X[1] = j
					Y[1] = i
					numberString = ""
					numberX, numberY = 0, 0
					isFirstNumber = true

					numberY, _ = strconv.Atoi(numberString)

				}

				// if its a number make a exception then pass the contains and keep the list
			} else {
				indexMatchOrder = []int{}
				numberString = ""
				numberX, numberY = 0, 0
				isFirstNumber = true
				X = []int{0, 0}
				Y = []int{0, 0}
			}

			fmt.Println("Order: ", indexMatchOrder)

			// fmt.Println("Numberstring: ", numberString)
			//fmt.Println("indexMatchOrder: ", indexMatchOrder)
			// fmt.Println("numberX: ", numberX)
			// fmt.Println("numberY: ", numberY)
			//, if the indexMatch is correct then put these numbers on hash map or just to lists to multiply later
			if len(indexMatchOrder) > 0 {
				indexMatchCorrect, increment = shared.CheckCondition(indexMatchOrder, func(prev, curr int) bool { return (prev < curr) })
				if indexMatchCorrect && increment <= 1 {
					//check if all are
					containsAll := true
					for _, e := range mulIndexes {
						containsFinal, _ := shared.ContainIntIndex(indexMatchOrder, e)
						if !containsFinal {
							containsAll = false
						}
					}

					if containsAll {
						//fmt.Println("Save X: ", numberX, "Save Y: ", numberY)

						stringClean := "mul(" + strconv.Itoa(numberX) + "," + strconv.Itoa(numberY) + ")"
						cleanMulArray = append(cleanMulArray, stringClean)

						fmt.Println("X[0]: ", X[0], "Y[0]}: ", Y[0], "{X[1]: ", X[1], "Y[1]: ", Y[1])

						matchMatrixIndexes = append(matchMatrixIndexes, [][2]int{{X[0], Y[0]}, {X[1], Y[1]}})

						prevChar := line[i-1 : i]
						fmt.Println("prevChar: ", prevChar)

						if prevChar == "<" {
							fmt.Println("Match >")
							sum++
						} else {
							fmt.Println("Match no >")
							lineRunes := []rune(lines[j])

							lineRunes[i-1] = '<'

							lines[j] = string(lineRunes)

							fmt.Println("prevChar New: ", line[i-1:i])
							fmt.Println("lines ", lines)

						}
						// else {

						// }

					}

				} else {

					if indexMatchOrder[len(indexMatchOrder)-1] == mulIndexes[0] {
						indexMatchOrder = []int{mulIndexes[0]}
						fmt.Println("NewindexMatchOrder: ", indexMatchOrder)
						isFirstNumber = true
					} else {
						indexMatchOrder = []int{}
						numberString = ""
						numberX, numberY = 0, 0
						isFirstNumber = true
					}

				}
			}
		}

	}
	fmt.Println("cleanMulArrayOrder: ", cleanMulArrayOrder)

	return lines, cleanMulArray, sum
}
func part1Match(lines []string) {
	sumSides := 0
	//XMAS
	matchString := []string{"X", "M", "A", "S"}
	matchStringReverse := []string{"S", "A", "M", "X"}

	var matchMatrixIndexes [][][2]int

	//matchMatrixIndexes = append(matchMatrixIndexes, [][2]int{{x,y},{x1,y1}})

	//Find horizontal
	_, sum := cleanStringMatch(lines, matchString, matchMatrixIndexes)
	fmt.Println("Sum Normal: ", sum)
	sumSides += sum
	//Reverse and Test
	_, sum = cleanStringMatch(lines, matchStringReverse, matchMatrixIndexes)
	fmt.Println("Sum Normal Reverse: ", sum)
	sumSides += sum

	//Vertical
	rotatedInput := rotateInput(lines)
	fmt.Println("Rotated ", 90, ": ", rotatedInput)
	_, sum = cleanStringMatch(rotatedInput, matchString, matchMatrixIndexes)
	fmt.Println("Sum ", 90, ": ", sum)
	sumSides += sum
	//Reverse and Test
	_, sum = cleanStringMatch(rotatedInput, matchStringReverse, matchMatrixIndexes)
	fmt.Println("Sum Reverse", 90, ": ", sum)
	sumSides += sum

	//Diagonal left - right
	rotatedInput = rotateInputNegative45(lines)
	fmt.Println("Rotated ", 45, ": ", rotatedInput)
	_, sum = cleanStringMatch(rotatedInput, matchString, matchMatrixIndexes)
	fmt.Println("Sum ", 45, ": ", sum)
	sumSides += sum
	//Reverse and Test
	_, sum = cleanStringMatch(rotatedInput, matchStringReverse, matchMatrixIndexes)
	fmt.Println("Sum Reverse", 45, ": ", sum)
	sumSides += sum

	//Diagonal right - left
	rotatedInput = rotateInput45(lines)
	fmt.Println("Rotated ", 45, ": ", rotatedInput)
	_, sum = cleanStringMatch(rotatedInput, matchString, matchMatrixIndexes)
	fmt.Println("Sum ", 45, ": ", sum)
	sumSides += sum
	//Reverse and Test
	_, sum = cleanStringMatch(rotatedInput, matchStringReverse, matchMatrixIndexes)
	fmt.Println("Sum Reverse", 45, ": ", sum)
	sumSides += sum

	fmt.Println("Part 1 Total Sum: ", sumSides)
}

// part 2
func part2Match(lines []string) {
	sumSides := 0
	sum := 0
	// //XMAS
	matchString := []string{"M", "A", "S"}
	matchStringReverse := []string{"S", "A", "M"}
	matchStringAlt := []string{"M", "<", "S"}
	matchStringReverseAlt := []string{"S", "<", "M"}

	var matchMatrixIndexes [][][2]int
	rows := len(lines)
	cols := len(lines[0])
	rotatedInputSplited := []string{}

	// //Diagonal right - left
	rotatedInput := rotateInput45(lines)
	fmt.Println("Rotated ", 45, ": ", rotatedInput)
	rotatedInput, _, sum = cleanStringMatch2(rotatedInput, matchString, matchMatrixIndexes)
	fmt.Println("Sum ", 45, ": ", sum)
	sumSides += sum
	//Reverse and Test
	rotatedInput, _, sum = cleanStringMatch2(rotatedInput, matchStringReverse, matchMatrixIndexes)
	fmt.Println("Sum Reverse", 45, ": ", sum)
	sumSides += sum

	fmt.Println("Rotated  45: ", rotatedInput)
	fmt.Println("Rotated: ", rotatedInput[3])
	fmt.Println("rows: ", rows)
	fmt.Println("cols: ", cols)

	//Iterate through each line in rotatedInput
	for _, line := range rotatedInput {
		// fmt.Println("Processing line:", line)
		var tempString string
		for j := 0; j < len(line); j++ {
			char := line[j : j+1]
			// fmt.Println("Current char:", char)
			if char != " " {
				tempString += char
			} else {
				if tempString != "" {
					rotatedInputSplited = append(rotatedInputSplited, tempString)
					tempString = ""
				}
			}
		}
		if tempString != "" {
			rotatedInputSplited = append(rotatedInputSplited, tempString)
		}
	}
	for i := range rotatedInputSplited {
		rotatedInputSplited[i] = strings.TrimSpace(rotatedInputSplited[i])
	}

	fmt.Println("rotatedInputSplited 45: ", rotatedInputSplited)

	rotatedInput = revert45Rotation(rotatedInputSplited, rows, cols)

	for i := range rotatedInput {
		rotatedInput[i] = strings.TrimSpace(rotatedInput[i])
	}

	fmt.Println("lines: ", lines)
	fmt.Println("rotatedInput: ", rotatedInput)

	for i, elem := range lines {
		fmt.Printf("lines[%d]: '%s' (length: %d)\n", i, elem, len(elem))
	}
	for i, elem := range rotatedInput {
		fmt.Printf("rotatedInput[%d]: '%s' (length: %d)\n", i, elem, len(elem))
	}

	rotatedInput = rotateInputNegative45(rotatedInput)
	fmt.Println("Rotated ", 45, ": ", rotatedInput)
	rotatedInput, _, sum = cleanStringMatch2(rotatedInput, matchStringAlt, matchMatrixIndexes)
	fmt.Println("Sum ", 45, ": ", sum)
	sumSides += sum
	//Reverse and Test
	rotatedInput, _, sum = cleanStringMatch2(rotatedInput, matchStringReverseAlt, matchMatrixIndexes)
	fmt.Println("Sum Reverse", 45, ": ", sum)
	sumSides += sum

	fmt.Println("Part 2 Total Sum: ", sumSides)
}

func main() {
	lines, err := shared.ReadLines("day-4/input-day-4.txt")
	if err != nil {
		log.Fatalf("readFail: %s", err)
	}
	//part1Match(lines)
	part2Match(lines)
	//change the cleanXmeans to also check reverse values
}

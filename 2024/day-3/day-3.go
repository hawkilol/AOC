package main

import (
	"2024/shared"
	"fmt"
	"strconv"

	// "io"
	"log"
)

// part 1
// Check if mul correct
func cleanMul(linesInput []string) ([]string, int) {
	lines := append([]string(nil), linesInput...)
	cleanMulArray := []string{}
	cleanMulArrayOrder := []int{}
	// mulString := []string{"m", "u", "l", "(", ",", ")"}
	mulString := []string{"m", "u", "l", "(", "IntX", ",", "IntY", ")"}
	mulIndexes := shared.RangeSlice(0, len(mulString)-1)
	fmt.Println("mulIndexes", mulIndexes)
	maxNumberLen := 3
	sum := 0
	for i, line := range lines {
		fmt.Println(i, line)
		var numberString string
		//var number int
		indexMatchOrder := []int{}
		indexMatchCorrect := false
		isFirstNumber := true
		numberX := 0
		numberY := 0

		increment := 0
		for i, _ := range line {

			currChar := line[i : i+1]
			currCharCheck := currChar
			fmt.Println(i, "->", currChar)

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

			contains, indexMatch := shared.ContainStringIndex(mulString, currCharCheck)

			if contains {
				indexMatchOrder = append(indexMatchOrder, indexMatch)
				cleanMulArrayOrder = append(cleanMulArrayOrder, indexMatch)

				if isFirstNumber {
					if currCharCheck == "," {
						if len(numberString) > maxNumberLen {
							indexMatchOrder = []int{}
							numberString = ""
							numberX, numberY = 0, 0
							isFirstNumber = true
							continue
						}
						numberX, _ = strconv.Atoi(numberString)
						isFirstNumber = false
						numberString = ""
					}
				} else if currCharCheck == ")" {
					// if should switch
					if len(numberString) > maxNumberLen {
						indexMatchOrder = []int{}
						numberString = ""
						numberX, numberY = 0, 0
						isFirstNumber = true
						continue
					}
					numberY, _ = strconv.Atoi(numberString)

					isFirstNumber = true
					numberString = ""
				}

				// if its a number make a exception then pass the contains and keep the list
			} else {
				indexMatchOrder = []int{}
				numberString = ""
				numberX, numberY = 0, 0
				isFirstNumber = true
			}

			fmt.Println("Order: ", indexMatchOrder)

			fmt.Println("Numberstring: ", numberString)
			fmt.Println("indexMatchOrder: ", indexMatchOrder)
			fmt.Println("numberX: ", numberX)
			fmt.Println("numberY: ", numberY)
			//, if the indexMatch is correct then put these numbers on hash map or just to lists to multiply later
			if len(indexMatchOrder) > 0 {
				indexMatchCorrect, increment = shared.CheckCondition(indexMatchOrder, func(prev, curr int) bool { return (prev < curr) || ((prev == 4 || prev == 6) && (prev <= curr)) })
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
						fmt.Println("Save X: ", numberX, "Save Y: ", numberY)

						stringClean := "mul(" + strconv.Itoa(numberX) + "," + strconv.Itoa(numberY) + ")"
						cleanMulArray = append(cleanMulArray, stringClean)

						sum += numberX * numberY
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
// Check if mul correct + conditions
func cleanMul2(linesInput []string) ([]string, int) {
	lines := append([]string(nil), linesInput...)
	cleanMulArray := []string{}
	cleanMulArrayOrder := []int{}
	// mulString := []string{"m", "u", "l", "(", ",", ")"}
	mulString := []string{"m", "u", "l", "(", "IntX", ",", "IntY", ")"}
	mulIndexes := shared.RangeSlice(0, len(mulString)-1)
	doString := []string{"d", "o", "(", ")"}
	doIndexes := shared.RangeSlice(0, len(doString)-1)

	dontString := []string{"d", "o", "n", "'", "t", "(", ")"}
	dontIndexes := shared.RangeSlice(0, len(dontString)-1)

	fmt.Println("mulIndexes", mulIndexes)
	maxNumberLen := 3
	sum := 0
	isMulAllowed := true

	for i, line := range lines {
		fmt.Println(i, line)
		var numberString string
		//var number int
		indexMatchOrder := []int{}
		indexMatchOrderDo := []int{}
		indexMatchOrderDont := []int{}
		indexMatchCorrect := false
		isFirstNumber := true
		numberX := 0
		numberY := 0

		increment := 0
		for i, _ := range line {

			currChar := line[i : i+1]
			currCharCheck := currChar
			fmt.Println(i, "->", currChar)

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

			contains, indexMatch := shared.ContainStringIndex(mulString, currCharCheck)

			containsDont, indexMatchDont := shared.ContainStringIndex(dontString, currCharCheck)

			containsDo, indexMatchDo := shared.ContainStringIndex(doString, currCharCheck)

			if containsDont {
				indexMatchOrderDont = append(indexMatchOrderDont, indexMatchDont)
			} else {
				indexMatchOrderDont = []int{}
			}

			if containsDo {
				indexMatchOrderDo = append(indexMatchOrderDo, indexMatchDo)
			} else {
				indexMatchOrderDo = []int{}
			}

			//Do -> enables isMulAllowed = true
			if len(indexMatchOrderDo) > 0 {
				indexMatchCorrectDo, incrementDo := shared.IncreaseCheck(indexMatchOrderDo)
				if indexMatchCorrectDo && incrementDo == 1 {
					//check if all are
					containsAllDo := true
					for _, e := range doIndexes {
						containsFinal, _ := shared.ContainIntIndex(indexMatchOrderDo, e)
						if !containsFinal {
							containsAllDo = false

						}
					}

					if containsAllDo {
						isMulAllowed = true
						indexMatchOrderDo = []int{}
					}
				} else {

					if indexMatchOrderDo[len(indexMatchOrderDo)-1] == doIndexes[0] {
						indexMatchOrderDo = []int{doIndexes[0]}
						fmt.Println("NewindexMatchOrderDo: ", indexMatchOrderDo)
					} else {
						indexMatchOrderDo = []int{}
					}
				}
			}
			//Dont -> disables isMulAllowed = false
			if len(indexMatchOrderDont) > 0 {
				indexMatchCorrectDont, incrementDont := shared.IncreaseCheck(indexMatchOrderDont)
				if indexMatchCorrectDont && incrementDont == 1 {
					//check if all are
					containsAllDont := true
					for _, e := range dontIndexes {
						containsFinal, _ := shared.ContainIntIndex(indexMatchOrderDont, e)
						if !containsFinal {

							containsAllDont = false
						}
					}

					if containsAllDont {

						isMulAllowed = false
						indexMatchOrderDont = []int{}
					}

				} else {

					if indexMatchOrderDont[len(indexMatchOrderDont)-1] == dontIndexes[0] {
						indexMatchOrderDont = []int{dontIndexes[0]}
						fmt.Println("NewindexMatchOrderDont: ", indexMatchOrderDont)
					} else {
						indexMatchOrderDont = []int{}
					}
				}
			}

			if contains {
				indexMatchOrder = append(indexMatchOrder, indexMatch)
				cleanMulArrayOrder = append(cleanMulArrayOrder, indexMatch)

				if isFirstNumber {
					if currCharCheck == "," {
						if len(numberString) > maxNumberLen {
							indexMatchOrder = []int{}
							numberString = ""
							numberX, numberY = 0, 0
							isFirstNumber = true
							continue
						}
						numberX, _ = strconv.Atoi(numberString)
						isFirstNumber = false
						numberString = ""
					}
				} else if currCharCheck == ")" {
					// if should switch
					if len(numberString) > maxNumberLen {
						indexMatchOrder = []int{}
						numberString = ""
						numberX, numberY = 0, 0
						isFirstNumber = true
						continue
					}
					numberY, _ = strconv.Atoi(numberString)

					isFirstNumber = true
					numberString = ""
				}

				// if its a number make a exception then pass the contains and keep the list
			} else {
				indexMatchOrder = []int{}
				numberString = ""
				numberX, numberY = 0, 0
				isFirstNumber = true
			}

			fmt.Println("isMulAllowed: ", isMulAllowed)
			fmt.Println("indexMatchOrderDont: ", indexMatchOrderDont)
			fmt.Println("indexMatchOrderDo: ", indexMatchOrderDo)

			fmt.Println("Order1: ", indexMatchOrder)

			fmt.Println("Numberstring: ", numberString)
			fmt.Println("indexMatchOrder: ", indexMatchOrder)
			fmt.Println("numberX: ", numberX)
			fmt.Println("numberY: ", numberY)
			//, if the indexMatch is correct then put these numbers on hash map or just to lists to multiply later
			if len(indexMatchOrder) > 0 {
				indexMatchCorrect, increment = shared.CheckCondition(indexMatchOrder, func(prev, curr int) bool { return (prev < curr) || ((prev == 4 || prev == 6) && (prev <= curr)) })
				if indexMatchCorrect && increment <= 1 {
					//check if all are
					containsAll := true
					for _, e := range mulIndexes {
						containsFinal, _ := shared.ContainIntIndex(indexMatchOrder, e)
						if !containsFinal {
							containsAll = false
						}
					}

					if containsAll && isMulAllowed {
						fmt.Println("Save X: ", numberX, "Save Y: ", numberY)

						stringClean := "mul(" + strconv.Itoa(numberX) + "," + strconv.Itoa(numberY) + ")"
						cleanMulArray = append(cleanMulArray, stringClean)

						sum += numberX * numberY
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

func main() {
	lines, err := shared.ReadLines("day-3/input-day-3.txt")
	if err != nil {
		log.Fatalf("readFail: %s", err)
	}

	// cleanMulArray, sum := cleanMul(lines)
	// fmt.Println("CleanMulInput: ", cleanMulArray)
	// fmt.Println("Part 1 - Sum: ", sum)

	cleanMulArray2, sum2 := cleanMul2(lines)
	fmt.Println("CleanMulInput: ", cleanMulArray2)
	fmt.Println("Part 2 - Sum2: ", sum2)

}

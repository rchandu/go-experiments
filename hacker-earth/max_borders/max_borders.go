package main

import (
	"fmt"
)

func main() {
	testCount := 0
	fmt.Scanln(&testCount)

	testOutputArr := make([]int, testCount)

	for testIdx := 0; testIdx < testCount; testIdx++ {
		rowCount := 0
		colCount := 0
		fmt.Scanln(&rowCount, &colCount)

		currTestMax := 0
		for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
			currRowMax := 0
			currCount := 0
			var prevChar byte

			rowStr := ""
			fmt.Scanln(&rowStr)
			for colIdx := 0; colIdx < colCount; colIdx++ {
				currCell := rowStr[colIdx]
				if currCell == '#' {
					switch {
					case currCell == prevChar:
						currCount++
					default:
						currCount = 1
					}
				} else {
					currCount = 0
				}
				if currCount > currRowMax {
					currRowMax = currCount
				}
				prevChar = currCell
			}

			if currRowMax > currTestMax {
				currTestMax = currRowMax
			}
			testOutputArr[testIdx] = currTestMax
		}
	}

	for i := 0; i < testCount; i++ {
		fmt.Println(testOutputArr[i])
	}
}

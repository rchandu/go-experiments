package main

import (
	"fmt"
)

/*
Problem
You are given a table with  rows and  columns. Each cell is colored with white or black. Considering the shapes created by black cells, what is the maximum border of these shapes? Border of a shape means the maximum number of consecutive black cells in any row or column without any white cell in between.

A shape is a set of connected cells. Two cells are connected if they share an edge. Note that no shape has a hole in it.

Input format

The first line contains  denoting the number of test cases.
The first line of each test case contains integers  denoting the number of rows and columns of the matrix. Here, '#' represents a black cell and '.' represents a white cell.
Each of the next  lines contains  integers.
Output format

Print the maximum border of the shapes.
*/

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

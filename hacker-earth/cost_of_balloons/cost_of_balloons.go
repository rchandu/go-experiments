package main

import "fmt"

func main() {
	testCount := 0
	fmt.Scanln(&testCount)

	testOutputArr := make([]int, testCount)

	for j := 0; j < testCount; j++ {
		greenPrice := 0
		purplePrice := 0
		fmt.Scanln(&greenPrice, &purplePrice)

		usersCount := 0
		fmt.Scanln(&usersCount)

		firstSolvedCount := 0
		secondSolvedCount := 0
		for i := 0; i < usersCount; i++ {
			solvedFirstProblem := 0
			solvedSecondProblem := 0
			fmt.Scanln(&solvedFirstProblem, &solvedSecondProblem)
			firstSolvedCount += solvedFirstProblem
			secondSolvedCount += solvedSecondProblem
		}

		variant1 := (firstSolvedCount * greenPrice) + (secondSolvedCount * purplePrice)
		variant2 := (firstSolvedCount * purplePrice) + (secondSolvedCount * greenPrice)

		if variant1 < variant2 {
			testOutputArr[j] = variant1
		} else {
			testOutputArr[j] = variant2
		}
	}

	for i := 0; i < testCount; i++ {
		fmt.Println(testOutputArr[i])
	}
}

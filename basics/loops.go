package basics

import (
	"fmt"
	"rchandu/experiments/utils"
)

func LoopExperiment() {
	utils.PrintLine("Loops")
	i := 9
	for i >= 0 {
		fmt.Println("i", i)
		i--
	}
	fmt.Println("i after", i)

	// For loop with initialization inside loop. Here j is not available outside of for context
	for j := 7; j <= 9; j++ {
		fmt.Println("j", j)
	}

	// while(true) kind of loop
	for {
		fmt.Println("loop will be printing I break it")
		i = i - 1
		if i <= -5 {
			break
		}
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println("Skipping even numbers")
			fmt.Println("Skipping even numbers")
			continue
		}
		fmt.Println(n)
	}
}

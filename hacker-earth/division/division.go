package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0
	fmt.Scanln(&count)

	var lastValue byte
	for i := 0; i < count; i++ {
		currValue := ""
		fmt.Scanf("%s", &currValue)
		lastPos := len(currValue) - 1
		lastValue = currValue[lastPos]
	}
	allLastValue, _ := strconv.ParseInt(string(lastValue), 10, 64)
	switch allLastValue == 0 {
	case true:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
	}
}

package utils

import (
	"fmt"
	"strings"
)

func PrintLine(titleStr string) {
	lineLen := len(titleStr) + 4
	fmt.Println()
	lineStr := strings.Repeat("-", lineLen)
	fmt.Println(lineStr)
	fmt.Println("| " + titleStr + " |")
	fmt.Println(lineStr)
}

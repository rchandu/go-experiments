package main

import (
	"fmt"
)

func main() {
	var myname = ""
	fmt.Scanf("%s", &myname)
	zCount := 0
	oCount := 0
	for j := 0; j < len(myname); j++ {
		currChar := myname[j]
		switch currChar {
		case 'z':
			zCount++
		case 'o':
			oCount++
		}
	}
	switch {
	case zCount*2 == oCount:
		{
			fmt.Println("yes")
		}
	default:
		fmt.Println("no")
	}
}

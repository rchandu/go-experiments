package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"
)

func printLine(titleStr string) {
	lineLen := len(titleStr) + 4
	fmt.Println()
	lineStr := strings.Repeat("-", lineLen)
	fmt.Println(lineStr)
	fmt.Println("| " + titleStr + " |")
	fmt.Println(lineStr)
}

func printSizesOfTypes() {
	printLine("Sizes of different types")

	var c_1 complex64
	fmt.Printf("The size of complex64 is %d\n", unsafe.Sizeof(c_1))

	var c_2 complex128
	fmt.Printf("The size of complex128 is %d\n", unsafe.Sizeof(c_2))

	var f_1 float32
	fmt.Printf("The size of float32 is %d\n", unsafe.Sizeof(f_1))

	var f_2 float64
	fmt.Printf("The size of float64 is %d\n", unsafe.Sizeof(f_2))

	var ui_8 uint8
	fmt.Printf("The size of uint8 is %d\n", unsafe.Sizeof(ui_8))

	var ui_16 uint16
	fmt.Printf("The size of uint16 is %d\n", unsafe.Sizeof(ui_16))

	var ui_32 uint32
	fmt.Printf("The size of uint32 is %d\n", unsafe.Sizeof(ui_32))

	var ui_64 uint64
	fmt.Printf("The size of uint64 is %d\n", unsafe.Sizeof(ui_64))

	var i_8 int8
	fmt.Printf("The size of int8 is %d\n", unsafe.Sizeof(i_8))

	var i_16 int16
	fmt.Printf("The size of int16 is %d\n", unsafe.Sizeof(i_16))

	var i_32 int32 // alias to rune
	fmt.Printf("The size of int32 is %d\n", unsafe.Sizeof(i_32))

	var i_64 int64
	fmt.Printf("The size of int64 is %d\n", unsafe.Sizeof(i_64))

	var u_p uintptr
	fmt.Printf("The size of uintptr is %d\n", unsafe.Sizeof(u_p))

	var err error
	fmt.Printf("The size of error is %d\n", unsafe.Sizeof(err))

	var b bool
	fmt.Printf("The size of bool is %d\n", unsafe.Sizeof(b))
}

func primitiveOperations() {
	// Go lang has 3 primitve types: string, numeric, bool and error type
	// It also contains numerous different int and float types with differing storages

	// String concatenation works too
	printLine("Operations between different types")
	fmt.Println("strings" + " are " + "concatenated")

	// Regular numeric math operations
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Basic boolean operations
	fmt.Println("true && false", true && false)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)
	fmt.Println("!!true", !!true)
}

func initializationAndDefaultValues() {
	// Initialization and unintialized values of different types
	printLine("Initialization and Uninitalized variables:")
	var a = "initial"
	fmt.Println("variable with initial value: ", a)
	var uninitializedInt int
	fmt.Println("uninitializedInt -", uninitializedInt)
	var uninitializedStr string
	fmt.Println("uninitializedStr -", uninitializedStr)
	var uninitializedBool bool
	fmt.Println("uninitializedBool -", uninitializedBool)
	shorthandVal := 200
	fmt.Println("shorthandVal -", shorthandVal)
}

func loops() {
	printLine("Loops")
	i := 9
	for i >= 0 {
		fmt.Println("i", i)
		i--
	}
	fmt.Println("i after", i)

	for j := 7; j <= 9; j++ {
		fmt.Println("j", j)
	}

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

func tryingOutSwitch() {
	printLine("Switch statement")

	i := 2
	switch i {
	case 1:
		fmt.Println("Hello from 1")
	case 2:
		fmt.Println("Hello from 2")
	case 3:
		fmt.Println("Hello from 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 4:
		fmt.Println("Night owl I see")
	case t.Hour() < 6:
		fmt.Println("Way too early in the morning...")
	case t.Hour() < 10:
		fmt.Println("Freshly morning")
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() < 8 || t.Hour() > 3:
		fmt.Println("It's evening time!!!")
	default:
		fmt.Println("It's night")
	}

	type task struct {
		name string
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case task:
			fmt.Println("A task struct type")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(task{name: "Do something"})
}

func main() {
	// Printing on console is done through fmt package.
	fmt.Println("hello there!")
	primitiveOperations()
	initializationAndDefaultValues()
	printSizesOfTypes()
	loops()
	tryingOutSwitch()
}

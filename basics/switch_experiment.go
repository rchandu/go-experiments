package basics

import (
	"fmt"
	"rchandu/experiments/utils"
	"time"
)

func SwitchExperiment() {
	utils.PrintLine("Switch statement")

	i := 2
	// Regular switch based on value
	switch i {
	case 1:
		fmt.Println("Hello from 1")
	case 2:
		fmt.Println("Hello from 2")
	case 3:
		fmt.Println("Hello from 3")
	}

	switch time.Now().Weekday() {
	// You can have multiple cases be pointing to a single case
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	// Switch doesn't really need an expression at top - weird but really cool
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

	// We are creating a basic struct
	type task struct {
		name string
	}

	whatAmI := func(i interface{}) {
		// This is a type switch. This compares type of value rather than value
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

package basics

import (
	"fmt"
	"os"
	"strconv"
)

func ControlFlow() {
	if len(os.Args) != 2 {
		fmt.Print("Usage: go run basics/2_control_flow.go <number>\n")
		return
	}
	argument := os.Args[1]

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to integer")
		return
	}

	switch {
	case value == 0:
		fmt.Println("Zero")
	case value > 0:
		fmt.Println("Positive")
	case value < 0:
		fmt.Println("Negative")
	default:
	}

}

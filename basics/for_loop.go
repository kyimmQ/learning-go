package basics

import "fmt"

func ForLoop() {
	// Code
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		fmt.Println(i, number)
	}

	for i := range 10 {
		fmt.Println(i)
	}
}

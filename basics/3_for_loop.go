package basics

import "fmt"

func ForLoop() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	// or
	// for ok := true; ok; ok = j < 5 {
	// 	fmt.Println(j)
	// 	j++
	// }

	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		fmt.Println(i, number)
	}

	for i := range 10 {
		fmt.Println(i)
	}
}

package main

import "fmt"

func main() {
	myInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, thisInt := range myInts {
		// fmt.Println(i, thisInt)
		// fmt.Printf(thisInt)

		if thisInt%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")

		}
	}
}

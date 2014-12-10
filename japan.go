package main

import "fmt"

func main() {
	//places := [5]int{30, 20, 20, 10}
	//places := [3]int{5, 7, 3}
	places := [9]int{6, 8, 5, 4, 7, 4, 2, 3, 1}

	total := 0

	for i := 1; i < len(places); i++ {
		if places[i] > places[i-1] {
			fmt.Println("decreace from ", places[i], " to ", places[i-1])
			total += places[i] - places[i-1]
		}
	}
	fmt.Println("Total: ", total)
}
